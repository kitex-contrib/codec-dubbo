package testservice

import (
	"github.com/cloudwego/kitex/server"
	hessian2 "github.com/kitex-contrib/codec-hessian2/pkg"
	"github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo.TestService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCodec(hessian2.NewHessian2Codec()))

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
