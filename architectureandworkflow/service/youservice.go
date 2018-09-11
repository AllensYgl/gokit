package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type YouServiceName interface {
	YouServiceFuncName( /* you need params */ ) /* (you need renturn value) */
	//....
}

type YouServiceStruct struct {
	youServiceAttributeName string
	//...
}

// you need to follow the interface (you define service)
func (ys YouServiceStruct) YouServiceFuncName() {
	// func content
	//...
}
