package file

import (
	"github.com/ncraft-io/armory/go/pkg/armory/file"
	"io"
	"net/http"
)

func (x *CreateFileRequest) DecodeHttpRequest(request *http.Request) error {
	f, header, err := request.FormFile("file")
	if err != nil {
		return err
	}

	x.File = &file.BinaryFile{
		Name: header.Filename,
		Size: header.Size,
		//MineType: header.Header,
	}

	defer func() {
		_ = f.Close()
	}()

	x.File.Content, err = io.ReadAll(f)
	if err != nil {
		return err
	}

	return nil
}
