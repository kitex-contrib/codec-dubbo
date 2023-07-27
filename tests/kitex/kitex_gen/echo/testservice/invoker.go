package testservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler echo.TestService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
