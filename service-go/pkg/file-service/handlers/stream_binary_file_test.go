package handlers

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/file"
	pb "github.com/ncraft-io/armory/go/pkg/armory/file/v1"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/storage"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/svc"
	nhttp "github.com/ncraft-io/ncraft/go/pkg/gokit/transport/http"
)

type trackedStore struct {
	content              string
	opens, closes, reads int
	openErr, readErr     error
	lastCtx              context.Context
}

func (s *trackedStore) Open(ctx context.Context, _ string) (io.ReadSeekCloser, storage.Info, error) {
	s.opens++
	s.lastCtx = ctx
	if s.openErr != nil {
		return nil, storage.Info{}, s.openErr
	}
	return &trackedReader{Reader: bytes.NewReader([]byte(s.content)), store: s}, storage.Info{
		Size: int64(len(s.content)), ContentType: "text/plain", ETag: "revision", Modified: time.Unix(1700000000, 0),
	}, nil
}
func (*trackedStore) Write(context.Context, string, []byte, string) error { return nil }

type trackedReader struct {
	*bytes.Reader
	store *trackedStore
}

func (r *trackedReader) Close() error { r.store.closes++; return nil }
func (r *trackedReader) Read(b []byte) (int, error) {
	r.store.reads++
	if r.store.readErr != nil {
		return 0, r.store.readErr
	}
	return r.Reader.Read(b)
}

func downloadRouter(s *fileServer, middleware endpoint.Middleware) http.Handler {
	e := svc.MakeGetFileEndpoint(s)
	if middleware != nil {
		e = middleware(e)
	}
	router := mux.NewRouter()
	svc.RegisterHttpHandler(router, svc.Endpoints{GetFileEndpoint: e}, nil, log.NewNopLogger())
	return router
}

func TestGeneratedDownloadProtocol(t *testing.T) {
	for _, tc := range []struct {
		name, method string
		headers      map[string]string
		status       int
		body         string
		noRead       bool
	}{
		{"get", "GET", nil, 200, "0123456789", false},
		{"head", "HEAD", nil, 200, "", true},
		{"range", "GET", map[string]string{"Range": "bytes=2-5"}, 206, "2345", false},
		{"suffix", "GET", map[string]string{"Range": "bytes=-2"}, 206, "89", false},
		{"unsatisfiable", "GET", map[string]string{"Range": "bytes=20-"}, 416, "", true},
		{"etag", "GET", map[string]string{"If-None-Match": `"revision"`}, 304, "", true},
		{"modified", "GET", map[string]string{"If-Modified-Since": time.Unix(1700000000, 0).UTC().Format(http.TimeFormat)}, 304, "", true},
		{"if-range", "GET", map[string]string{"Range": "bytes=2-5", "If-Range": `"revision"`}, 206, "2345", false},
		{"if-range-stale", "GET", map[string]string{"Range": "bytes=2-5", "If-Range": `"stale"`}, 200, "0123456789", false},
		{"envelope", "GET", nil, 200, "0123456789", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			store := &trackedStore{content: "0123456789"}
			calls := 0
			handler := downloadRouter(&fileServer{store: store}, func(next endpoint.Endpoint) endpoint.Endpoint {
				return func(ctx context.Context, request interface{}) (interface{}, error) {
					calls++
					response, err := next(ctx, request)
					if store.opens != 0 {
						t.Fatal("reader opened before middleware finished")
					}
					if err == nil && len(response.(*file.BinaryFile).Content) != 0 {
						t.Fatal("buffered HTTP body")
					}
					return response, err
				}
			})
			r := httptest.NewRequest(tc.method, "/armory/file/v1/files/file.txt?envelope=true", nil)
			for key, value := range tc.headers {
				r.Header.Set(key, value)
			}
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)
			if w.Code != tc.status {
				t.Fatalf("status=%d body=%q", w.Code, w.Body.String())
			}
			if tc.status != 416 && w.Body.String() != tc.body {
				t.Fatalf("body=%q", w.Body.String())
			}
			if tc.method == "HEAD" && w.Header().Get("Content-Length") != "10" {
				t.Fatal("missing HEAD length")
			}
			if tc.status == 206 && w.Header().Get("Content-Range") != "bytes 2-5/10" && tc.name == "range" {
				t.Fatal("wrong range")
			}
			if calls != 1 || store.opens != 1 || store.closes != 1 {
				t.Fatalf("calls=%d store=%+v", calls, store)
			}
			if tc.noRead && store.reads != 0 {
				t.Fatalf("unnecessary read: %+v", store)
			}
		})
	}
}

func TestDownloadMiddlewareMayDiscardResponse(t *testing.T) {
	for _, mode := range []string{"reject-before", "reject-after", "replace"} {
		t.Run(mode, func(t *testing.T) {
			store := &trackedStore{content: "secret"}
			h := downloadRouter(&fileServer{store: store}, func(next endpoint.Endpoint) endpoint.Endpoint {
				return func(ctx context.Context, request interface{}) (interface{}, error) {
					if mode == "reject-before" {
						return nil, core.NewInvalidArgumentError("denied")
					}
					result, err := next(ctx, request)
					if err != nil {
						return nil, err
					}
					if mode == "reject-after" {
						return result, core.NewInvalidArgumentError("denied")
					}
					return &file.BinaryFile{Name: "replacement"}, nil
				}
			})
			w := httptest.NewRecorder()
			h.ServeHTTP(w, httptest.NewRequest("GET", "/armory/file/v1/files/file.txt", nil))
			if store.opens != 0 || strings.Contains(w.Body.String(), "secret") {
				t.Fatalf("discarded response streamed: %+v %q", store, w.Body.String())
			}
			if mode == "replace" && !strings.Contains(w.Body.String(), "replacement") {
				t.Fatal("replacement was lost")
			}
			if mode != "replace" && w.Code != http.StatusBadRequest {
				t.Fatalf("status=%d", w.Code)
			}
		})
	}
}

func TestDownloadStorageErrorsAndEmptyFiles(t *testing.T) {
	for _, tc := range []struct {
		name   string
		store  trackedStore
		status int
	}{
		{"missing", trackedStore{openErr: storage.ErrNotFound}, 404},
		{"unavailable", trackedStore{openErr: errors.New("secret backend error")}, 500},
		{"empty", trackedStore{}, 200},
		{"read-failure", trackedStore{content: "data", readErr: errors.New("secret read error")}, 200},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			downloadRouter(&fileServer{store: &tc.store}, nil).ServeHTTP(w, httptest.NewRequest("GET", "/armory/file/v1/files/test", nil))
			if w.Code != tc.status {
				t.Fatalf("status=%d body=%q", w.Code, w.Body.String())
			}
			if strings.Contains(w.Body.String(), "secret") {
				t.Fatal("leaked backend error")
			}
			if tc.store.openErr == nil && tc.store.closes != 1 {
				t.Fatal("reader not closed")
			}
			if tc.name == "empty" && (w.Body.Len() != 0 || w.Header().Get("Content-Length") != "0") {
				t.Fatal("empty file became JSON metadata")
			}
			if tc.name == "read-failure" && w.Body.Len() != 0 {
				t.Fatal("appended an error after file headers")
			}
		})
	}
}

func TestStreamingCancellationAndRPCFallback(t *testing.T) {
	store := &trackedStore{content: "body"}
	s := &fileServer{store: store}
	ctx, cancel := context.WithCancel(context.Background())
	ctx = nhttp.RequestToContext(ctx, httptest.NewRequest("GET", "/", nil))
	result, err := s.GetFile(ctx, &pb.GetFileRequest{Name: "file"})
	if err != nil {
		t.Fatal(err)
	}
	cancel()
	w := httptest.NewRecorder()
	if err := svc.EncodeHTTPGenericResponse(ctx, w, result); err != nil {
		t.Fatal(err)
	}
	if store.lastCtx.Err() != context.Canceled || store.closes != 1 || w.Body.Len() != 0 {
		t.Fatal("cancellation did not reach stream")
	}
	rpc, err := s.GetFile(context.Background(), &pb.GetFileRequest{Name: "file"})
	if err != nil || string(rpc.GetContent()) != "body" {
		t.Fatalf("RPC fallback: %v %v", rpc, err)
	}
	encoded, err := svc.EncodeGRPCGetFileResponse(context.Background(), rpc)
	if err != nil || encoded != rpc {
		t.Fatal("gRPC response type changed")
	}
}

func TestNestedNamesStillWorkThroughRPC(t *testing.T) {
	s, _ := localService(t)
	name := "nested/文件 #?.txt"
	if _, err := s.CreateFile(context.Background(), &pb.CreateFileRequest{File: &file.BinaryFile{Name: name, Content: []byte("nested")}}); err != nil {
		t.Fatal(err)
	}
	result, err := s.GetFile(context.Background(), &pb.GetFileRequest{Name: name})
	if err != nil || string(result.GetContent()) != "nested" {
		t.Fatalf("nested RPC: %v %v", result, err)
	}
}
