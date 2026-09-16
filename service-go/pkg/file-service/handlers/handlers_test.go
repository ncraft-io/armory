package handlers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/file"
	pb "github.com/ncraft-io/armory/go/pkg/armory/file/v1"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/svc"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/util"
	"github.com/rs/cors"
)

func localService(t *testing.T) (*fileServer, string) {
	t.Helper()
	root := filepath.Join(t.TempDir(), "files")
	s, err := NewServiceWithConfig(&util.Config{RootUrl: "https://files.example.com/", StaticProviderConfig: &util.StaticProviderConfig{Root: root}})
	if err != nil {
		t.Fatal(err)
	}
	return s, root
}

func TestLocalFileRoundTripAndDownloads(t *testing.T) {
	s, _ := localService(t)
	testRoundTrip(t, s)
}

func testRoundTrip(t *testing.T, s *fileServer) {
	t.Helper()
	for _, name := range []string{"文件 #?%.txt", "nested/deep/文件 #?%.txt", "nested/literal%2F.txt"} {
		t.Run(name, func(t *testing.T) { testNamedRoundTrip(t, s, name) })
	}
}

func testNamedRoundTrip(t *testing.T, s *fileServer, name string) {
	t.Helper()
	ctx := context.Background()
	router := mux.NewRouter()
	svc.RegisterHttpHandler(router, svc.Endpoints{
		GetFileEndpoint:         svc.MakeGetFileEndpoint(s),
		CreateFileEndpoint:      svc.MakeCreateFileEndpoint(s),
		BatchCreateFileEndpoint: svc.MakeBatchCreateFileEndpoint(s),
	}, nil, log.NewNopLogger())
	handler := cors.AllowAll().Handler(router)
	created, err := s.CreateFile(ctx, &pb.CreateFileRequest{File: &file.BinaryFile{Name: name, Content: []byte("0123456789"), Size: 999}})
	if err != nil {
		t.Fatal(err)
	}
	if created.Name != name || created.Size != 10 || created.CreateTime == nil || len(created.Content) != 0 {
		t.Fatalf("invalid upload metadata: %v", created)
	}
	u, err := url.Parse(created.Url.Format())
	if err != nil || u.Path != "/armory/file/v1/files/"+name || u.RawQuery != "" || u.Fragment != "" {
		t.Fatalf("invalid file URL: %v %v", created.Url, err)
	}
	got, err := s.GetFile(ctx, &pb.GetFileRequest{Name: name})
	if err != nil || string(got.GetContent()) != "0123456789" || got.Size != 10 {
		t.Fatalf("read: %v %v", got, err)
	}
	for _, tc := range []struct {
		method, rangeHeader, body string
		status                    int
	}{
		{"GET", "", "0123456789", 200}, {"GET", "bytes=2-5", "2345", 206}, {"GET", "bytes=-3", "789", 206},
		{"GET", "bytes=20-", "", 416}, {"HEAD", "", "", 200}, {"OPTIONS", "", "", 204}, {"POST", "", "", 404},
	} {
		t.Run(tc.method+tc.rangeHeader, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, u.String(), nil)
			r.Header.Set("Range", tc.rangeHeader)
			if tc.method == "OPTIONS" {
				r.Header.Set("Origin", "https://client.example.com")
				r.Header.Set("Access-Control-Request-Method", "GET")
				r.Header.Set("Access-Control-Request-Headers", "Range")
			}
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)
			if w.Code != tc.status {
				t.Fatalf("status=%d body=%q", w.Code, w.Body.String())
			}
			if tc.body != "" && w.Body.String() != tc.body {
				t.Fatalf("body=%q", w.Body.String())
			}
			if tc.method == "HEAD" && (w.Body.Len() != 0 || w.Header().Get("Content-Length") != "10") {
				t.Fatal("invalid HEAD response")
			}
			if tc.status == 206 && w.Header().Get("Content-Range") == "" {
				t.Fatal("missing content range")
			}
		})
	}
	if _, err := s.CreateFile(ctx, &pb.CreateFileRequest{File: &file.BinaryFile{Name: name, Content: []byte("updated")}}); err != nil {
		t.Fatal(err)
	}
	got, err = s.GetFile(ctx, &pb.GetFileRequest{Name: name})
	if err != nil || string(got.Content) != "updated" {
		t.Fatal("overwrite failed")
	}
	if _, err := s.GetFile(ctx, &pb.GetFileRequest{Name: "missing"}); !core.IsNotFoundError(err) {
		t.Fatalf("not found: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", "/armory/file/v1/files/missing", nil))
	if w.Code != 404 {
		t.Fatalf("missing download status %d", w.Code)
	}
}

func TestValidationAndBatch(t *testing.T) {
	s, root := localService(t)
	ctx := context.Background()
	for _, name := range []string{"", "../escape", "/absolute", "a/../escape", "a//b", "a\\b", ".", "a\x00b"} {
		if _, err := s.CreateFile(ctx, &pb.CreateFileRequest{File: &file.BinaryFile{Name: name, Content: []byte("test")}}); !core.IsInvalidArgumentError(err) {
			t.Errorf("name %q: %v", name, err)
		}
	}
	for _, req := range []*pb.CreateFileRequest{nil, {}, {File: &file.BinaryFile{Name: "empty"}}, {File: &file.BinaryFile{Name: "bad", Content: []byte("test"), MineType: "invalid"}}} {
		if _, err := s.CreateFile(ctx, req); !core.IsInvalidArgumentError(err) {
			t.Errorf("invalid request: %v", err)
		}
	}
	if _, err := s.GetFile(ctx, nil); !core.IsInvalidArgumentError(err) {
		t.Fatal(err)
	}
	if _, err := s.BatchCreateFile(ctx, &pb.BatchCreateFileRequest{Files: []*file.BinaryFile{{Name: "valid", Content: []byte("a")}, nil}}); err == nil {
		t.Fatal("accepted invalid batch")
	}
	if _, err := os.Stat(filepath.Join(root, "valid")); !os.IsNotExist(err) {
		t.Fatal("invalid batch wrote a file")
	}
	resp, err := s.BatchCreateFile(ctx, &pb.BatchCreateFileRequest{Files: []*file.BinaryFile{{Name: "a.txt", Content: []byte("a")}, {Name: "b.txt", Content: []byte("b")}}})
	if err != nil || len(resp.GetBinaryFiles()) != 2 {
		t.Fatalf("batch: %v %v", resp, err)
	}
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := s.GetFile(ctx, &pb.GetFileRequest{Name: "a.txt"}); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation: %v", err)
	}
}

func TestLocalStorageCannotEscapeThroughSymlink(t *testing.T) {
	s, root := localService(t)
	outside := t.TempDir()
	if err := os.WriteFile(filepath.Join(outside, "secret"), []byte("secret"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(root, "link")); err != nil {
		t.Fatal(err)
	}
	if _, err := s.GetFile(context.Background(), &pb.GetFileRequest{Name: "link/secret"}); err == nil {
		t.Fatal("read escaped root")
	}
	if _, err := s.CreateFile(context.Background(), &pb.CreateFileRequest{File: &file.BinaryFile{Name: "link/secret", Content: []byte("overwrite")}}); err == nil {
		t.Fatal("write escaped root")
	}
	content, _ := os.ReadFile(filepath.Join(outside, "secret"))
	if string(content) != "secret" {
		t.Fatal("outside file changed")
	}
}

func TestInvalidConfiguration(t *testing.T) {
	for _, cfg := range []*util.Config{nil, {}, {Provider: "other"}, {Provider: "s3"}} {
		if _, err := NewServiceWithConfig(cfg); err == nil {
			t.Fatal("invalid configuration accepted")
		}
	}
	s := &fileServer{initErr: errors.New("secret configuration error")}
	_, err := s.GetFile(context.Background(), &pb.GetFileRequest{Name: "file"})
	if err == nil || strings.Contains(err.Error(), "secret") {
		t.Fatalf("unavailable storage: %v", err)
	}
}

func TestS3FileRoundTripAndDownloads(t *testing.T) {
	var mu sync.Mutex
	objects := map[string][]byte{}
	types := map[string]string{}
	endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if !strings.HasPrefix(r.URL.Path, "/armory-files/service/") {
			t.Errorf("unexpected object path: %s", r.URL.Path)
		}
		if !strings.HasPrefix(r.Header.Get("Authorization"), "AWS4-HMAC-SHA256 ") {
			t.Error("request was not signed")
		}
		if r.Method == "PUT" {
			var body io.Reader = r.Body
			if strings.Contains(r.Header.Get("Content-Encoding"), "aws-chunked") {
				body = httputil.NewChunkedReader(body)
			}
			data, err := io.ReadAll(body)
			if err != nil {
				t.Error(err)
				w.WriteHeader(400)
				return
			}
			objects[r.URL.Path], types[r.URL.Path] = data, r.Header.Get("Content-Type")
			w.Header().Set("ETag", `"etag"`)
			return
		}
		data, ok := objects[r.URL.Path]
		if !ok {
			w.Header().Set("X-Amz-Error-Code", "NoSuchKey")
			w.WriteHeader(404)
			fmt.Fprint(w, "<Error><Code>NoSuchKey</Code></Error>")
			return
		}
		w.Header().Set("ETag", `"etag"`)
		w.Header().Set("Content-Type", types[r.URL.Path])
		http.ServeContent(w, r, "file", time.Unix(1700000000, 0), bytes.NewReader(data))
	}))
	defer endpoint.Close()
	create := false
	s, err := NewServiceWithConfig(&util.Config{Provider: "s3", RootUrl: "https://files.example.com/", S3ProviderConfig: &util.S3ProviderConfig{
		Endpoint: endpoint.URL, BucketName: "armory-files", Region: "us-east-1", AccessKey: "test-key", SecretKey: "test-secret", Prefix: "service", CreateBucket: &create,
	}})
	if err != nil {
		t.Fatal(err)
	}
	testRoundTrip(t, s)
}
