package hessian2

import "github.com/apache/dubbo-go-hessian2/java_exception"

type Throwabler interface {
	java_exception.Throwabler
}

func NewException(detailMessage string) Throwabler {
	return java_exception.NewException(detailMessage)
}
