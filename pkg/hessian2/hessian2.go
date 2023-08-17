package hessian2

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

func NewEncoder() iface.Encoder {
	return hessian.NewEncoder()
}

func NewDecoder(b []byte) iface.Decoder {
	return hessian.NewDecoder(b)
}
