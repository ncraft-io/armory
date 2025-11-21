package unitable

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func init() {
	core.RegisterJSONTypeDecoder("unitable.DbQuery_Parameter", &DbQueryParameterCodec{})
	core.RegisterJSONTypeEncoder("unitable.DbQuery_Parameter", &DbQueryParameterCodec{})
}

type DbQueryParameterCodec struct {
}

func (codec *DbQueryParameterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	parameter := (*DbQuery_Parameter)(ptr)
	if a.ValueType() == jsoniter.StringValue {
		if err := parameter.Parse(a.ToString()); err != nil {
			iter.ReportError("decode db query parameter", err.Error())
		}
	}
}

func (codec *DbQueryParameterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	parameter := (*DbQuery_Parameter)(ptr)
	return parameter == nil || len(parameter.Name) == 0
}

func (codec *DbQueryParameterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	parameter := (*DbQuery_Parameter)(ptr)
	stream.WriteString(parameter.Format())
}
