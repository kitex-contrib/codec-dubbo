package echo

import hessian "github.com/apache/dubbo-go-hessian2"

func init() {
	hessian.RegisterPOJO(&EchoRequest{})
	hessian.RegisterPOJO(&EchoResponse{})
}
