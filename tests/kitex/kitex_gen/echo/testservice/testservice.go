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
		"EchoInt":   kitex.NewMethodInfo(echoIntHandler, newTestServiceEchoIntArgs, newTestServiceEchoIntResult, false),
		"EchoByte":  kitex.NewMethodInfo(echoByteHandler, newTestServiceEchoByteArgs, newTestServiceEchoByteResult, false),
		"EchoBytes": kitex.NewMethodInfo(echoBytesHandler, newTestServiceEchoBytesArgs, newTestServiceEchoBytesResult, false),
		"EchoInt8":  kitex.NewMethodInfo(echoInt8Handler, newTestServiceEchoInt8Args, newTestServiceEchoInt8Result, false),
		"EchoInt8s": kitex.NewMethodInfo(echoInt8sHandler, newTestServiceEchoInt8sArgs, newTestServiceEchoInt8sResult, false),
		"Echo":      kitex.NewMethodInfo(echoHandler, newTestServiceEchoArgs, newTestServiceEchoResult, false),
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

func echoByteHandler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoByteArgs)
	realResult := result.(*echo.TestServiceEchoByteResult)
	success, err := handler.(echo.TestService).EchoByte(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}

func newTestServiceEchoByteArgs() interface{} {
	return echo.NewTestServiceEchoByteArgs()
}

func newTestServiceEchoByteResult() interface{} {
	return echo.NewTestServiceEchoByteResult()
}

func echoBytesHandler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBytesArgs)
	realResult := result.(*echo.TestServiceEchoBytesResult)
	success, err := handler.(echo.TestService).EchoBytes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newTestServiceEchoBytesArgs() interface{} {
	return echo.NewTestServiceEchoBytesArgs()
}

func newTestServiceEchoBytesResult() interface{} {
	return echo.NewTestServiceEchoBytesResult()
}

func echoInt8Handler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt8Args)
	realResult := result.(*echo.TestServiceEchoInt8Result)
	success, err := handler.(echo.TestService).EchoInt8(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}

func newTestServiceEchoInt8Args() interface{} {
	return echo.NewTestServiceEchoInt8Args()
}

func newTestServiceEchoInt8Result() interface{} {
	return echo.NewTestServiceEchoInt8Result()
}

func echoInt8sHandler(ctx context.Context, handler, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt8sArgs)
	realResult := result.(*echo.TestServiceEchoInt8sResult)
	success, err := handler.(echo.TestService).EchoInt8s(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newTestServiceEchoInt8sArgs() interface{} {
	return echo.NewTestServiceEchoInt8sArgs()
}

func newTestServiceEchoInt8sResult() interface{} {
	return echo.NewTestServiceEchoInt8sResult()
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

func (p *kClient) EchoByte(ctx context.Context, req byte) (r byte, err error) {
	var _args echo.TestServiceEchoByteArgs
	_args.Req = req
	var _result echo.TestServiceEchoByteResult
	if err = p.c.Call(ctx, "EchoByte", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBytes(ctx context.Context, req []byte) (r []byte, err error) {
	var _args echo.TestServiceEchoBytesArgs
	_args.Req = req
	var _result echo.TestServiceEchoBytesResult
	if err = p.c.Call(ctx, "EchoBytes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt8(ctx context.Context, req int8) (r int8, err error) {
	var _args echo.TestServiceEchoInt8Args
	_args.Req = req
	var _result echo.TestServiceEchoInt8Result
	if err = p.c.Call(ctx, "EchoInt8", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt8s(ctx context.Context, req []int8) (r []int8, err error) {
	var _args echo.TestServiceEchoInt8sArgs
	_args.Req = req
	var _result echo.TestServiceEchoInt8sResult
	if err = p.c.Call(ctx, "EchoInt8s", &_args, &_result); err != nil {
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
