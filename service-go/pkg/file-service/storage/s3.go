package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/service-go/pkg/file-service/util"
	nstorage "github.com/ncraft-io/ncraft/go/pkg/ncraft/storage"
	_ "github.com/ncraft-io/ncraft/go/pkg/ncraft/storage/minio"
)

// The S3 implementation and configuration live in ncraft; this adapter only
// translates between object metadata and the file service's backend interface.
type s3Store struct{ backend nstorage.StreamingStorage }

func newS3(ctx context.Context, cfg *util.S3ProviderConfig) (*s3Store, error) {
	if cfg == nil {
		return nil, fmt.Errorf("file.s3Provider is required")
	}
	conf := *cfg
	if conf.Vendor == "" {
		conf.Vendor = "s3"
	}
	if conf.Vendor != "s3" && conf.Vendor != "minio" {
		return nil, fmt.Errorf("file.s3Provider.vendor must be s3 or minio")
	}
	backend, err := nstorage.NewStorageWithContext(ctx, &conf)
	if err != nil {
		return nil, err
	}
	streaming, ok := backend.(nstorage.StreamingStorage)
	if !ok {
		return nil, fmt.Errorf("storage vendor %q does not support streaming", conf.Vendor)
	}
	return &s3Store{backend: streaming}, nil
}

func (s *s3Store) Open(ctx context.Context, name string) (io.ReadSeekCloser, Info, error) {
	if err := ValidateName(name); err != nil {
		return nil, Info{}, err
	}
	reader, obj, err := s.backend.Open(ctx, name)
	if core.IsNotFoundError(err) {
		return nil, Info{}, ErrNotFound
	}
	if err != nil {
		return nil, Info{}, err
	}
	info := Info{Size: obj.Size, ETag: obj.Etag}
	if obj.ContentType != nil {
		info.ContentType = obj.ContentType.Format()
	}
	if obj.LastModified != nil {
		info.Modified = obj.LastModified.ToTime()
	}
	return reader, info, nil
}

func (s *s3Store) Write(ctx context.Context, name string, content []byte, contentType string) error {
	if err := ValidateName(name); err != nil {
		return err
	}
	obj := &nstorage.Object{Key: name, Content: content, Size: int64(len(content))}
	if contentType != "" {
		var err error
		obj.ContentType, err = core.ParseMediaType(contentType)
		if err != nil {
			return err
		}
	}
	return s.backend.WriteContext(ctx, obj, nil)
}
