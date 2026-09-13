package handlers

import (
	"sync"
)

var unitablePool sync.Pool

func init() {
	unitablePool = sync.Pool{
		New: func() interface{} {
			return NewUnitable()
		},
	}
}

func GetUnitable() *Unitable {
	return unitablePool.Get().(*Unitable)
}

func PutUnitable(ts *Unitable) {
	if ts != nil {
		unitablePool.Put(ts)
	}
}

type Unitable struct {
}

func NewUnitable() *Unitable {
	return &Unitable{}
}
