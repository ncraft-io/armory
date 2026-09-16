// Package storage provides the shared backend for file RPCs and HTTP downloads.
package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/ncraft-io/armory/service-go/pkg/file-service/util"
)

var ErrNotFound = errors.New("file not found")

type Info struct {
	Size        int64
	ContentType string
	ETag        string
	Modified    time.Time
}

type Store interface {
	Open(context.Context, string) (io.ReadSeekCloser, Info, error)
	Write(context.Context, string, []byte, string) error
}

// ValidateName keeps names relative and identical across filesystem and S3 backends.
func ValidateName(name string) error {
	if !fs.ValidPath(name) || name == "." || strings.ContainsAny(name, "\\\x00\r\n") {
		return fmt.Errorf("invalid file name %q: expected a relative path without dot segments", name)
	}
	return nil
}

func New(ctx context.Context, cfg *util.Config) (Store, error) {
	if cfg == nil {
		return nil, errors.New("file configuration is required")
	}
	switch strings.ToLower(cfg.Provider) {
	case "", "local", "static":
		if cfg.StaticProviderConfig == nil || strings.TrimSpace(cfg.StaticProviderConfig.Root) == "" {
			return nil, errors.New("file.staticProvider.root is required")
		}
		root, err := filepath.Abs(cfg.StaticProviderConfig.Root)
		if err != nil {
			return nil, err
		}
		if err = os.MkdirAll(root, 0755); err != nil {
			return nil, err
		}
		return &localStore{root: root}, nil
	case "s3", "minio":
		return newS3(ctx, cfg.S3ProviderConfig)
	default:
		return nil, fmt.Errorf("unsupported file provider %q", cfg.Provider)
	}
}

type localStore struct{ root string }

func (s *localStore) Open(ctx context.Context, name string) (io.ReadSeekCloser, Info, error) {
	if err := ctx.Err(); err != nil {
		return nil, Info{}, err
	}
	if err := ValidateName(name); err != nil {
		return nil, Info{}, err
	}
	root, err := os.OpenRoot(s.root)
	if err != nil {
		return nil, Info{}, err
	}
	defer root.Close()
	f, err := root.Open(filepath.FromSlash(name))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, Info{}, ErrNotFound
	}
	if err != nil {
		return nil, Info{}, err
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, Info{}, err
	}
	if !info.Mode().IsRegular() {
		f.Close()
		return nil, Info{}, ErrNotFound
	}
	return f, Info{Size: info.Size(), ContentType: mime.TypeByExtension(path.Ext(name)), Modified: info.ModTime()}, nil
}

func (s *localStore) Write(ctx context.Context, name string, content []byte, contentType string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := ValidateName(name); err != nil {
		return err
	}
	root, err := os.OpenRoot(s.root)
	if err != nil {
		return err
	}
	defer root.Close()
	if err = root.MkdirAll(filepath.FromSlash(path.Dir(name)), 0755); err != nil {
		return err
	}
	return root.WriteFile(filepath.FromSlash(name), content, 0644)
}
