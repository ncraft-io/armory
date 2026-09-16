package handlers

import (
	"context"
	"errors"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/file"
	pb "github.com/ncraft-io/armory/go/pkg/armory/file/v1"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/storage"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/util"
	nhttp "github.com/ncraft-io/ncraft/go/pkg/gokit/transport/http"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

type fileServer struct {
	pb.UnimplementedFileServer
	rootURL string
	store   storage.Store
	initErr error
}

// NewService preserves the generated service factory signature.
func NewService() pb.FileServer {
	cfg, err := util.GetConfig()
	if err != nil {
		return &fileServer{initErr: err}
	}
	service, err := NewServiceWithConfig(cfg)
	if err != nil {
		logs.Errorw("initialize file storage", "error", err)
		return &fileServer{initErr: err}
	}
	return service
}

// NewServiceWithConfig constructs the backend once and reports startup errors.
func NewServiceWithConfig(cfg *util.Config) (*fileServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	backend, err := storage.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &fileServer{rootURL: strings.TrimRight(cfg.RootUrl, "/"), store: backend}, nil
}

func (s *fileServer) fileURL(name string) *core.Url {
	parts := strings.Split(name, "/")
	for i := range parts {
		parts[i] = url.PathEscape(parts[i])
	}
	result, _ := core.NewUrl(s.rootURL + "/armory/file/v1/files/" + strings.Join(parts, "/"))
	return result
}

func (s *fileServer) GetFile(ctx context.Context, in *pb.GetFileRequest) (*file.BinaryFile, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("file request is required")
	}
	if err := storage.ValidateName(in.Name); err != nil {
		return nil, core.NewInvalidArgumentError("%s", err)
	}
	if s.initErr != nil {
		return nil, core.NewInternalError("file storage is unavailable")
	}
	if request, ok := nhttp.RequestFromContext(ctx); ok &&
		(request.Method == http.MethodGet || request.Method == http.MethodHead) {
		// Keep the RPC return type. The HTTP encoder opens the stream only after
		// all endpoint middleware has accepted this exact response object.
		name := in.Name
		response := &file.BinaryFile{Name: name, Url: s.fileURL(name)}
		writer := &StreamBinaryFile{Name: name, Open: func(ctx context.Context) (io.ReadSeekCloser, storage.Info, error) {
			return s.store.Open(ctx, name)
		}}
		if err := nhttp.BindResponseWriter(ctx, response, writer); err != nil {
			return nil, err
		}
		return response, nil
	}
	reader, info, err := s.store.Open(ctx, in.Name)
	if err != nil {
		return nil, fileError(ctx, in.Name, err)
	}
	defer reader.Close()
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fileError(ctx, in.Name, err)
	}
	return &file.BinaryFile{Name: in.Name, Content: content, Size: int64(len(content)), MineType: info.ContentType, Url: s.fileURL(in.Name), CreateTime: core.FromTime(info.Modified)}, nil
}

func validateFile(f *file.BinaryFile) error {
	if f == nil {
		return core.NewInvalidArgumentError("file is required")
	}
	if err := storage.ValidateName(f.Name); err != nil {
		return core.NewInvalidArgumentError("%s", err)
	}
	if len(f.Content) == 0 {
		return core.NewInvalidArgumentError("file content is required")
	}
	if f.MineType != "" {
		if mediaType, _, err := mime.ParseMediaType(f.MineType); err != nil || !strings.Contains(mediaType, "/") {
			return core.NewInvalidArgumentError("invalid file MIME type")
		}
	}
	return nil
}

func (s *fileServer) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*file.BinaryFile, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("file request is required")
	}
	if err := validateFile(in.File); err != nil {
		return nil, err
	}
	if s.initErr != nil {
		return nil, core.NewInternalError("file storage is unavailable")
	}
	contentType := in.File.MineType
	if contentType == "" {
		contentType = mime.TypeByExtension(path.Ext(in.File.Name))
	}
	if contentType == "" {
		contentType = http.DetectContentType(in.File.Content)
	}
	if err := s.store.Write(ctx, in.File.Name, in.File.Content, contentType); err != nil {
		return nil, fileError(ctx, in.File.Name, err)
	}
	return &file.BinaryFile{Name: in.File.Name, MineType: contentType, Size: int64(len(in.File.Content)), Url: s.fileURL(in.File.Name), CreateTime: core.Now()}, nil
}

func (s *fileServer) BatchCreateFile(ctx context.Context, in *pb.BatchCreateFileRequest) (*pb.BatchCreateFileResponse, error) {
	if in == nil || len(in.Files) == 0 {
		return nil, core.NewInvalidArgumentError("files are required")
	}
	// Validate the entire batch before writing. Storage failures are not transactional.
	for _, f := range in.Files {
		if err := validateFile(f); err != nil {
			return nil, err
		}
	}
	resp := &pb.BatchCreateFileResponse{}
	for _, f := range in.Files {
		created, err := s.CreateFile(ctx, &pb.CreateFileRequest{File: f})
		if err != nil {
			return nil, err
		}
		resp.BinaryFiles = append(resp.BinaryFiles, created)
	}
	return resp, nil
}

func fileError(ctx context.Context, name string, err error) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if errors.Is(err, storage.ErrNotFound) {
		return core.NewNotFoundError("file %s not found", name)
	}
	logs.Errorw("file storage operation failed", "file", name, "error", err)
	return core.NewInternalError("file storage operation failed for %s", name)
}
