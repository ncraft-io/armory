package handlers

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/util"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"os"
	"path"
	"sync"

	"github.com/ncraft-io/armory/go/pkg/armory/file"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/file/v1"
)

var (
	_ = file.BinaryFile{}
)

var fileOnce sync.Once
var fileCfg *util.Config

func getConfig() *util.Config {
	fileOnce.Do(func() {
		fileCfg, _ = util.GetConfig()
	})

	return fileCfg
}

type fileServer struct {
	pb.UnimplementedFileServer

	Cfg *util.Config
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.FileServer {
	return fileServer{
		Cfg: getConfig(),
	}
}

// GetFile implements Interface.
func (s fileServer) GetFile(ctx context.Context, in *pb.GetFileRequest) (*file.BinaryFile, error) {
	if len(in.Name) == 0 {
		return nil, core.NewInvalidArgumentError("not set the file name")
	}

	p := path.Join(s.Cfg.StaticProviderConfig.Root, in.Name)
	_, err := os.Stat(p)
	if err == nil {
		if content, err := os.ReadFile(p); err != nil {
			return nil, core.NewInternalError("failed to read the file from local. file: %s, error: %s", in.Name, err.Error())
		} else {
			resp := &file.BinaryFile{
				Name:    in.Name,
				Content: content,
				// Name:
				// MineType:
				// Url:
				// Size:
				// Content:
				// CreateTime:
			}
			return resp, nil
		}
	} else if os.IsNotExist(err) {
		return nil, core.NewNotFoundError("failed to found the file from local. file: %s", in.Name)
	} else {
		return nil, core.NewInternalError("failed to found the file from local. file: %s, error: %s", in.Name, err.Error())
	}
}

// CreateFile implements Interface.
func (s fileServer) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*file.BinaryFile, error) {
	if in.File == nil {
		return nil, core.NewInvalidArgumentError("not set the file")
	}
	if len(in.File.Name) == 0 {
		return nil, core.NewInvalidArgumentError("not set the file name")
	}
	if len(in.File.Content) == 0 {
		return nil, core.NewInvalidArgumentError("not set the file content")
	}

	p := path.Join(s.Cfg.StaticProviderConfig.Root, in.File.Name)
	if err := os.WriteFile(p, in.File.Content, os.ModePerm); err != nil {
		return nil, core.NewInternalError("failed to write the file to local. error: %s", err.Error())
	}

	url, _ := core.NewUrl(s.Cfg.RootUrl + "/armory/file/v1/files/" + in.File.Name)
	resp := &file.BinaryFile{
		Url:        url,
		CreateTime: core.Now(),
	}
	logs.Infow("create the file", "file", url.Format())
	return resp, nil
}

// BatchCreateFile implements Interface.
func (s fileServer) BatchCreateFile(ctx context.Context, in *pb.BatchCreateFileRequest) (*pb.BatchCreateFileResponse, error) {
	resp := &pb.BatchCreateFileResponse{
		// BinaryFiles:
	}
	return resp, nil
}
