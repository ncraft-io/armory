package handlers

import (
	"context"
	"errors"
	"io"
	"mime"
	"net/http"
	"path"
	"strconv"

	"github.com/ncraft-io/armory/service-go/pkg/file-service/storage"
	nhttp "github.com/ncraft-io/ncraft/go/pkg/gokit/transport/http"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

// StreamBinaryFile is the HTTP representation attached to a BinaryFile result.
// Open must return a fresh reader and its metadata; no resources are held until
// the encoder runs. RPC callers continue to receive BinaryFile.Content.
type StreamBinaryFile struct {
	Name string
	Open func(context.Context) (io.ReadSeekCloser, storage.Info, error)
}

var _ nhttp.ResponseWriter = (*StreamBinaryFile)(nil)

func (f *StreamBinaryFile) WriteHttpResponse(ctx context.Context, w http.ResponseWriter) error {
	request, ok := nhttp.RequestFromContext(ctx)
	if !ok {
		return errors.New("streaming file response requires an HTTP request")
	}
	reader, info, err := f.Open(ctx)
	if err != nil {
		return fileError(ctx, f.Name, err)
	}
	defer reader.Close()

	// ServerAfter sets a JSON content type. Remove it so ServeContent can sniff
	// the content when the backend does not supply a MIME type.
	w.Header().Del("Content-Type")
	if info.ContentType != "" {
		w.Header().Set("Content-Type", info.ContentType)
	}
	w.Header().Set("Content-Disposition", mime.FormatMediaType("attachment", map[string]string{"filename": path.Base(f.Name)}))
	w.Header().Set("Access-Control-Expose-Headers", "Accept-Ranges, Content-Range, Content-Length, Last-Modified, ETag, Content-Disposition")
	if info.ETag != "" {
		w.Header().Set("ETag", strconv.Quote(info.ETag))
	}

	input := &downloadReader{ReadSeekCloser: reader, ctx: ctx}
	output := &downloadWriter{ResponseWriter: w}
	http.ServeContent(output, request.WithContext(ctx), path.Base(f.Name), info.Modified, input)
	// ServeContent owns the response status, including 304/416. Once it starts
	// writing, a transport failure must not append a JSON error to the file.
	if err := errors.Join(input.err, output.err); err != nil {
		logs.Warnw("file download interrupted", "file", f.Name, "error", err)
	}
	return nil
}

type downloadReader struct {
	io.ReadSeekCloser
	ctx context.Context
	err error
}

func (r *downloadReader) Read(p []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		r.err = err
		return 0, err
	}
	n, err := r.ReadSeekCloser.Read(p)
	if err != nil && err != io.EOF {
		r.err = err
	}
	return n, err
}

type downloadWriter struct {
	http.ResponseWriter
	err error
}

func (w *downloadWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

func (w *downloadWriter) Write(p []byte) (int, error) {
	n, err := w.ResponseWriter.Write(p)
	if err != nil {
		w.err = err
	}
	return n, err
}
