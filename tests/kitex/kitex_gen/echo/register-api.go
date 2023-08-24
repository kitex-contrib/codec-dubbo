package echo

import hessian2 "github.com/kitex-contrib/codec-dubbo/pkg/hessian2"

var objects = []interface{}{
	&EchoRequest{},
	&EchoResponse{},
}

func init() {
	hessian2.Register(objects)
}
