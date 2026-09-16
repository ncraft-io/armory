package handlers

import nserver "github.com/ncraft-io/ncraft/go/pkg/gokit/server"

var _ nserver.Starter = (*fileServer)(nil)

// Start reports storage initialization failures before transports accept requests.
// The file service owns no listeners or background tasks; the generated server
// handles HTTP, including Range requests, and gRPC lifecycle management.
func (s *fileServer) Start(nserver.Config) error { return s.initErr }
