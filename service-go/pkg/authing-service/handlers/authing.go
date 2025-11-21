package handlers

import (
	"github.com/ncraft-io/armory/service-go/pkg/authing"
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

func GetUserToken() *authing.UserToken {
	return GetAuthing().UserToken
}

type Authing struct {
	UserToken *authing.UserToken
}

func NewAuthing() *Authing {
	return &Authing{UserToken: authing.NewUserToken()}
}
