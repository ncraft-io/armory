package unitable

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"net/http"

	nhttp "github.com/ncraft-io/ncraft/go/pkg/gokit/transport/http"
)

var cfg *nhttp.Config

func init() {
	cfg = nhttp.NewConfig()
}

func (x *ExportRowResponse) WriteHttpResponse(ctx context.Context, writer http.ResponseWriter) error {
	if x == nil {
		return nil
	}

	_ = ctx

	isExportFile := func(obj *core.Object) bool {
		if obj != nil && len(obj.Vals) == 2 {
			return obj.GetValue("@fileName") != nil && obj.GetValue("@fileContent") != nil
		}
		return false
	}

	if len(x.Objects) == 1 && isExportFile(x.Objects[0]) {
		writer.Header().Set("Content-Type", "application/octet-stream")

		content := x.Objects[0].GetValue("@fileContent").GetBytes()
		if _, err := writer.Write(content); err != nil {
			return err
		}
		return nil
	} else {
		stream := jsoniter.NewStream(jsoniter.ConfigFastest, writer, 512)

		enveloped := nhttp.IsEnvelopeStyle(ctx, cfg.GetStyle())
		if enveloped {
			stream.WriteObjectStart()
			stream.WriteObjectField("code")
			stream.WriteString("200")
			stream.WriteRaw(",")
			stream.WriteObjectField("message")
			stream.WriteString("OK")
			stream.WriteRaw(",")
			stream.WriteObjectField("data")
			stream.WriteVal(x.Objects)
			stream.WriteObjectEnd()
		} else {
			stream.WriteVal(x.Objects)
		}

		if err := stream.Flush(); err != nil {
			return err
		}
	}
	return nil
}
