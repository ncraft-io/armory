package handlers

import (
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"regexp"
	"sync"
)

var nameRegex = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

var unitableOnce sync.Once
var ut *Instance

type Instance struct {
	Synchro *synchro.Synchro
	Queries map[string]*unitable.DbQuery
}

func (s unitableServer) Synchro() *synchro.Synchro {
	return s.instance.Synchro
}

func (s unitableServer) Queries() map[string]*unitable.DbQuery {
	return s.instance.Queries
}
