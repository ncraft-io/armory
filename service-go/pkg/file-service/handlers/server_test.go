package handlers

import (
	"errors"
	"testing"

	nserver "github.com/ncraft-io/ncraft/go/pkg/gokit/server"
)

func TestFileStartupReportsInitializationResult(t *testing.T) {
	service, _ := localService(t)
	// No address is needed: the service does not open a separate Range listener.
	if err := service.Start(nserver.Config{}); err != nil {
		t.Fatal(err)
	}
	failure := errors.New("storage initialization failed")
	if err := (&fileServer{initErr: failure}).Start(nserver.Config{}); !errors.Is(err, failure) {
		t.Fatalf("startup error = %v", err)
	}
}
