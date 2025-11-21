package unitable

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"time"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("time.Time", &TimeCodec{})
}

type TimeCodec struct {
}

func (codec *TimeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	t := (*time.Time)(ptr)
	if a.ValueType() == jsoniter.StringValue {
		str := a.ToString()
		if len(str) == 0 {
			t = nil
		} else {
			if err := t.UnmarshalText([]byte(a.ToString())); err != nil {
				iter.ReportError("decode time failed", err.Error())
			}
		}
	} else if a.ValueType() == jsoniter.NilValue {
		t = nil
	}
}
