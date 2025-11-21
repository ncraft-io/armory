package file

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (x *BinaryFile) WriteHttpResponse(ctx context.Context, writer http.ResponseWriter) error {
	if x != nil {
		if len(x.Content) > 0 {
			//switch x.Format {
			//case "mvt":
			//	writer.Header().Set("Content-Type", "application/vnd.mapbox-vector-tile")
			//case "png":
			//	writer.Header().Set("Content-Type", "image/png")
			//case "jpeg", "jpg":
			//	writer.Header().Set("Content-Type", "image/jpeg")
			//}
			writer.Header().Set("Content-Description", "File Transfer")
			writer.Header().Set("Content-Type", "application/octet-stream")
			writer.Header().Set("Content-Disposition", "attachment; filename=\""+x.Name+"\"")

			//if len(x.Encoding) > 0 {
			//	writer.Header().Set("Content-Encoding", x.Encoding)
			//}
			size, err := writer.Write(x.Content)
			if err != nil {
				return err
			}
			if size != len(x.Content) {
				return errors.New(fmt.Sprintf("failed to write all the bytes to http, excepted: %d, actually: %d", len(x.Content), size))
			}
		} else {
			content, err := jsoniter.Marshal(x)
			if err != nil {
				return err
			}
			writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			writer.Header().Set("Content-Length", strconv.Itoa(len(content)))
			size, err := writer.Write(content)
			if err != nil {
				return err
			}
			if size != len(content) {
				return errors.New(fmt.Sprintf("failed to write all the bytes to http, excepted: %d, actually: %d", len(x.Content), size))
			}
		}
	}
	return nil
}
