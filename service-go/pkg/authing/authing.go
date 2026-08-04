package authing

import (
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"sync"
)

var authingOnce sync.Once
var au *Authing

func GetAuthing() *Authing {
	authingOnce.Do(func() {
		au = NewAuthing()
	})

	return au
}

func GetUserToken() *UserToken {
	return GetAuthing().UserToken
}

type Authing struct {
	UserToken *UserToken
	Config    *Config
}

func NewAuthing() *Authing {
	conf := &Config{}
	err := config.ScanFrom(conf, "authing")
	logs.Warnw("failed to load authing config", "error", err)

	return &Authing{
		UserToken: NewUserToken(),
		Config:    conf,
	}
}
