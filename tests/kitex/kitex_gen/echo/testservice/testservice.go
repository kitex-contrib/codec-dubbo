package testservice

import (
	"context"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

func serviceInfo() *kitex.ServiceInfo {
	return testServiceServiceInfo
}

var testServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TestService"
	handlerType := (*echo.TestService)(nil)
	methods := map[string]kitex.MethodInfo{
		"EchoInt": kitex.NewMethodInfo(echoIntHandler, newTestServiceEchoIntArgs, newTestServiceEchoIntResult, false),
		"Echo":    kitex.NewMethodInfo(echoHandler, newTestServiceEchoArgs, newTestServiceEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "echo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoIntHandler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoIntArgs)
	realResult := result.(*echo.TestServiceEchoIntResult)
	success, err := handler.(echo.TestService).EchoInt(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}

func newTestServiceEchoIntArgs() interface{} {
	return echo.NewTestServiceEchoIntArgs()
}

func newTestServiceEchoIntResult() interface{} {
	return echo.NewTestServiceEchoIntResult()
}

func echoHandler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoArgs)
	realResult := result.(*echo.TestServiceEchoResult)
	success, err := handler.(echo.TestService).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newTestServiceEchoArgs() interface{} {
	return echo.NewTestServiceEchoArgs()
}

func newTestServiceEchoResult() interface{} {
	return echo.NewTestServiceEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoInt(ctx context.Context, req int32) (r int32, err error) {
	var _args echo.TestServiceEchoIntArgs
	_args.Req = req
	var _result echo.TestServiceEchoIntResult
	if err = p.c.Call(ctx, "EchoInt", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Echo(ctx context.Context, req *echo.EchoRequest) (r *echo.EchoResponse, err error) {
	var _args echo.TestServiceEchoArgs
	_args.Req = req
	var _result echo.TestServiceEchoResult
	if err = p.c.Call(ctx, "Echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
