// Code generated by Kitex v0.7.1. DO NOT EDIT.

package testservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

func serviceInfo() *kitex.ServiceInfo {
	return testServiceServiceInfo
}

var testServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TestService"
	handlerType := (*echo.TestService)(nil)
	methods := map[string]kitex.MethodInfo{
		"EchoInt":                kitex.NewMethodInfo(echoIntHandler, newTestServiceEchoIntArgs, newTestServiceEchoIntResult, false),
		"EchoBool":               kitex.NewMethodInfo(echoBoolHandler, newTestServiceEchoBoolArgs, newTestServiceEchoBoolResult, false),
		"EchoByte":               kitex.NewMethodInfo(echoByteHandler, newTestServiceEchoByteArgs, newTestServiceEchoByteResult, false),
		"EchoInt16":              kitex.NewMethodInfo(echoInt16Handler, newTestServiceEchoInt16Args, newTestServiceEchoInt16Result, false),
		"EchoInt32":              kitex.NewMethodInfo(echoInt32Handler, newTestServiceEchoInt32Args, newTestServiceEchoInt32Result, false),
		"EchoInt64":              kitex.NewMethodInfo(echoInt64Handler, newTestServiceEchoInt64Args, newTestServiceEchoInt64Result, false),
		"EchoDouble":             kitex.NewMethodInfo(echoDoubleHandler, newTestServiceEchoDoubleArgs, newTestServiceEchoDoubleResult, false),
		"EchoString":             kitex.NewMethodInfo(echoStringHandler, newTestServiceEchoStringArgs, newTestServiceEchoStringResult, false),
		"EchoBinary":             kitex.NewMethodInfo(echoBinaryHandler, newTestServiceEchoBinaryArgs, newTestServiceEchoBinaryResult, false),
		"Echo":                   kitex.NewMethodInfo(echoHandler, newTestServiceEchoArgs, newTestServiceEchoResult, false),
		"EchoBoolList":           kitex.NewMethodInfo(echoBoolListHandler, newTestServiceEchoBoolListArgs, newTestServiceEchoBoolListResult, false),
		"EchoByteList":           kitex.NewMethodInfo(echoByteListHandler, newTestServiceEchoByteListArgs, newTestServiceEchoByteListResult, false),
		"EchoInt16List":          kitex.NewMethodInfo(echoInt16ListHandler, newTestServiceEchoInt16ListArgs, newTestServiceEchoInt16ListResult, false),
		"EchoInt32List":          kitex.NewMethodInfo(echoInt32ListHandler, newTestServiceEchoInt32ListArgs, newTestServiceEchoInt32ListResult, false),
		"EchoInt64List":          kitex.NewMethodInfo(echoInt64ListHandler, newTestServiceEchoInt64ListArgs, newTestServiceEchoInt64ListResult, false),
		"EchoDoubleList":         kitex.NewMethodInfo(echoDoubleListHandler, newTestServiceEchoDoubleListArgs, newTestServiceEchoDoubleListResult, false),
		"EchoStringList":         kitex.NewMethodInfo(echoStringListHandler, newTestServiceEchoStringListArgs, newTestServiceEchoStringListResult, false),
		"EchoBinaryList":         kitex.NewMethodInfo(echoBinaryListHandler, newTestServiceEchoBinaryListArgs, newTestServiceEchoBinaryListResult, false),
		"EchoBool2BoolMap":       kitex.NewMethodInfo(echoBool2BoolMapHandler, newTestServiceEchoBool2BoolMapArgs, newTestServiceEchoBool2BoolMapResult, false),
		"EchoBool2ByteMap":       kitex.NewMethodInfo(echoBool2ByteMapHandler, newTestServiceEchoBool2ByteMapArgs, newTestServiceEchoBool2ByteMapResult, false),
		"EchoBool2Int16Map":      kitex.NewMethodInfo(echoBool2Int16MapHandler, newTestServiceEchoBool2Int16MapArgs, newTestServiceEchoBool2Int16MapResult, false),
		"EchoBool2Int32Map":      kitex.NewMethodInfo(echoBool2Int32MapHandler, newTestServiceEchoBool2Int32MapArgs, newTestServiceEchoBool2Int32MapResult, false),
		"EchoBool2Int64Map":      kitex.NewMethodInfo(echoBool2Int64MapHandler, newTestServiceEchoBool2Int64MapArgs, newTestServiceEchoBool2Int64MapResult, false),
		"EchoBool2DoubleMap":     kitex.NewMethodInfo(echoBool2DoubleMapHandler, newTestServiceEchoBool2DoubleMapArgs, newTestServiceEchoBool2DoubleMapResult, false),
		"EchoBool2StringMap":     kitex.NewMethodInfo(echoBool2StringMapHandler, newTestServiceEchoBool2StringMapArgs, newTestServiceEchoBool2StringMapResult, false),
		"EchoBool2BinaryMap":     kitex.NewMethodInfo(echoBool2BinaryMapHandler, newTestServiceEchoBool2BinaryMapArgs, newTestServiceEchoBool2BinaryMapResult, false),
		"EchoBool2BoolListMap":   kitex.NewMethodInfo(echoBool2BoolListMapHandler, newTestServiceEchoBool2BoolListMapArgs, newTestServiceEchoBool2BoolListMapResult, false),
		"EchoBool2ByteListMap":   kitex.NewMethodInfo(echoBool2ByteListMapHandler, newTestServiceEchoBool2ByteListMapArgs, newTestServiceEchoBool2ByteListMapResult, false),
		"EchoBool2Int16ListMap":  kitex.NewMethodInfo(echoBool2Int16ListMapHandler, newTestServiceEchoBool2Int16ListMapArgs, newTestServiceEchoBool2Int16ListMapResult, false),
		"EchoBool2Int32ListMap":  kitex.NewMethodInfo(echoBool2Int32ListMapHandler, newTestServiceEchoBool2Int32ListMapArgs, newTestServiceEchoBool2Int32ListMapResult, false),
		"EchoBool2Int64ListMap":  kitex.NewMethodInfo(echoBool2Int64ListMapHandler, newTestServiceEchoBool2Int64ListMapArgs, newTestServiceEchoBool2Int64ListMapResult, false),
		"EchoBool2DoubleListMap": kitex.NewMethodInfo(echoBool2DoubleListMapHandler, newTestServiceEchoBool2DoubleListMapArgs, newTestServiceEchoBool2DoubleListMapResult, false),
		"EchoBool2StringListMap": kitex.NewMethodInfo(echoBool2StringListMapHandler, newTestServiceEchoBool2StringListMapArgs, newTestServiceEchoBool2StringListMapResult, false),
		"EchoBool2BinaryListMap": kitex.NewMethodInfo(echoBool2BinaryListMapHandler, newTestServiceEchoBool2BinaryListMapArgs, newTestServiceEchoBool2BinaryListMapResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "echo",
		"ServiceFilePath": `api.thrift`,
		"IDLAnnotations": map[string][]string{
			"JavaClassName": []string{"org.apache.dubbo.tests.api.UserProvider"},
		},
	}

	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoIntHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
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

func echoBoolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBoolArgs)
	realResult := result.(*echo.TestServiceEchoBoolResult)
	success, err := handler.(echo.TestService).EchoBool(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoBoolArgs() interface{} {
	return echo.NewTestServiceEchoBoolArgs()
}

func newTestServiceEchoBoolResult() interface{} {
	return echo.NewTestServiceEchoBoolResult()
}

func echoByteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
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

func echoInt16Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt16Args)
	realResult := result.(*echo.TestServiceEchoInt16Result)
	success, err := handler.(echo.TestService).EchoInt16(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoInt16Args() interface{} {
	return echo.NewTestServiceEchoInt16Args()
}

func newTestServiceEchoInt16Result() interface{} {
	return echo.NewTestServiceEchoInt16Result()
}

func echoInt32Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt32Args)
	realResult := result.(*echo.TestServiceEchoInt32Result)
	success, err := handler.(echo.TestService).EchoInt32(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoInt32Args() interface{} {
	return echo.NewTestServiceEchoInt32Args()
}

func newTestServiceEchoInt32Result() interface{} {
	return echo.NewTestServiceEchoInt32Result()
}

func echoInt64Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt64Args)
	realResult := result.(*echo.TestServiceEchoInt64Result)
	success, err := handler.(echo.TestService).EchoInt64(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoInt64Args() interface{} {
	return echo.NewTestServiceEchoInt64Args()
}

func newTestServiceEchoInt64Result() interface{} {
	return echo.NewTestServiceEchoInt64Result()
}

func echoDoubleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoDoubleArgs)
	realResult := result.(*echo.TestServiceEchoDoubleResult)
	success, err := handler.(echo.TestService).EchoDouble(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoDoubleArgs() interface{} {
	return echo.NewTestServiceEchoDoubleArgs()
}

func newTestServiceEchoDoubleResult() interface{} {
	return echo.NewTestServiceEchoDoubleResult()
}

func echoStringHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoStringArgs)
	realResult := result.(*echo.TestServiceEchoStringResult)
	success, err := handler.(echo.TestService).EchoString(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newTestServiceEchoStringArgs() interface{} {
	return echo.NewTestServiceEchoStringArgs()
}

func newTestServiceEchoStringResult() interface{} {
	return echo.NewTestServiceEchoStringResult()
}

func echoBinaryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBinaryArgs)
	realResult := result.(*echo.TestServiceEchoBinaryResult)
	success, err := handler.(echo.TestService).EchoBinary(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBinaryArgs() interface{} {
	return echo.NewTestServiceEchoBinaryArgs()
}

func newTestServiceEchoBinaryResult() interface{} {
	return echo.NewTestServiceEchoBinaryResult()
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
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

func echoBoolListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBoolListArgs)
	realResult := result.(*echo.TestServiceEchoBoolListResult)
	success, err := handler.(echo.TestService).EchoBoolList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBoolListArgs() interface{} {
	return echo.NewTestServiceEchoBoolListArgs()
}

func newTestServiceEchoBoolListResult() interface{} {
	return echo.NewTestServiceEchoBoolListResult()
}

func echoByteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoByteListArgs)
	realResult := result.(*echo.TestServiceEchoByteListResult)
	success, err := handler.(echo.TestService).EchoByteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoByteListArgs() interface{} {
	return echo.NewTestServiceEchoByteListArgs()
}

func newTestServiceEchoByteListResult() interface{} {
	return echo.NewTestServiceEchoByteListResult()
}

func echoInt16ListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt16ListArgs)
	realResult := result.(*echo.TestServiceEchoInt16ListResult)
	success, err := handler.(echo.TestService).EchoInt16List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoInt16ListArgs() interface{} {
	return echo.NewTestServiceEchoInt16ListArgs()
}

func newTestServiceEchoInt16ListResult() interface{} {
	return echo.NewTestServiceEchoInt16ListResult()
}

func echoInt32ListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt32ListArgs)
	realResult := result.(*echo.TestServiceEchoInt32ListResult)
	success, err := handler.(echo.TestService).EchoInt32List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoInt32ListArgs() interface{} {
	return echo.NewTestServiceEchoInt32ListArgs()
}

func newTestServiceEchoInt32ListResult() interface{} {
	return echo.NewTestServiceEchoInt32ListResult()
}

func echoInt64ListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoInt64ListArgs)
	realResult := result.(*echo.TestServiceEchoInt64ListResult)
	success, err := handler.(echo.TestService).EchoInt64List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoInt64ListArgs() interface{} {
	return echo.NewTestServiceEchoInt64ListArgs()
}

func newTestServiceEchoInt64ListResult() interface{} {
	return echo.NewTestServiceEchoInt64ListResult()
}

func echoDoubleListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoDoubleListArgs)
	realResult := result.(*echo.TestServiceEchoDoubleListResult)
	success, err := handler.(echo.TestService).EchoDoubleList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoDoubleListArgs() interface{} {
	return echo.NewTestServiceEchoDoubleListArgs()
}

func newTestServiceEchoDoubleListResult() interface{} {
	return echo.NewTestServiceEchoDoubleListResult()
}

func echoStringListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoStringListArgs)
	realResult := result.(*echo.TestServiceEchoStringListResult)
	success, err := handler.(echo.TestService).EchoStringList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoStringListArgs() interface{} {
	return echo.NewTestServiceEchoStringListArgs()
}

func newTestServiceEchoStringListResult() interface{} {
	return echo.NewTestServiceEchoStringListResult()
}

func echoBinaryListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBinaryListArgs)
	realResult := result.(*echo.TestServiceEchoBinaryListResult)
	success, err := handler.(echo.TestService).EchoBinaryList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBinaryListArgs() interface{} {
	return echo.NewTestServiceEchoBinaryListArgs()
}

func newTestServiceEchoBinaryListResult() interface{} {
	return echo.NewTestServiceEchoBinaryListResult()
}

func echoBool2BoolMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2BoolMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2BoolMapResult)
	success, err := handler.(echo.TestService).EchoBool2BoolMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2BoolMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2BoolMapArgs()
}

func newTestServiceEchoBool2BoolMapResult() interface{} {
	return echo.NewTestServiceEchoBool2BoolMapResult()
}

func echoBool2ByteMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2ByteMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2ByteMapResult)
	success, err := handler.(echo.TestService).EchoBool2ByteMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2ByteMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2ByteMapArgs()
}

func newTestServiceEchoBool2ByteMapResult() interface{} {
	return echo.NewTestServiceEchoBool2ByteMapResult()
}

func echoBool2Int16MapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int16MapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int16MapResult)
	success, err := handler.(echo.TestService).EchoBool2Int16Map(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int16MapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int16MapArgs()
}

func newTestServiceEchoBool2Int16MapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int16MapResult()
}

func echoBool2Int32MapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int32MapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int32MapResult)
	success, err := handler.(echo.TestService).EchoBool2Int32Map(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int32MapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int32MapArgs()
}

func newTestServiceEchoBool2Int32MapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int32MapResult()
}

func echoBool2Int64MapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int64MapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int64MapResult)
	success, err := handler.(echo.TestService).EchoBool2Int64Map(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int64MapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int64MapArgs()
}

func newTestServiceEchoBool2Int64MapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int64MapResult()
}

func echoBool2DoubleMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2DoubleMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2DoubleMapResult)
	success, err := handler.(echo.TestService).EchoBool2DoubleMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2DoubleMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2DoubleMapArgs()
}

func newTestServiceEchoBool2DoubleMapResult() interface{} {
	return echo.NewTestServiceEchoBool2DoubleMapResult()
}

func echoBool2StringMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2StringMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2StringMapResult)
	success, err := handler.(echo.TestService).EchoBool2StringMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2StringMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2StringMapArgs()
}

func newTestServiceEchoBool2StringMapResult() interface{} {
	return echo.NewTestServiceEchoBool2StringMapResult()
}

func echoBool2BinaryMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2BinaryMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2BinaryMapResult)
	success, err := handler.(echo.TestService).EchoBool2BinaryMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2BinaryMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2BinaryMapArgs()
}

func newTestServiceEchoBool2BinaryMapResult() interface{} {
	return echo.NewTestServiceEchoBool2BinaryMapResult()
}

func echoBool2BoolListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2BoolListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2BoolListMapResult)
	success, err := handler.(echo.TestService).EchoBool2BoolListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2BoolListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2BoolListMapArgs()
}

func newTestServiceEchoBool2BoolListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2BoolListMapResult()
}

func echoBool2ByteListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2ByteListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2ByteListMapResult)
	success, err := handler.(echo.TestService).EchoBool2ByteListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2ByteListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2ByteListMapArgs()
}

func newTestServiceEchoBool2ByteListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2ByteListMapResult()
}

func echoBool2Int16ListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int16ListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int16ListMapResult)
	success, err := handler.(echo.TestService).EchoBool2Int16ListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int16ListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int16ListMapArgs()
}

func newTestServiceEchoBool2Int16ListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int16ListMapResult()
}

func echoBool2Int32ListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int32ListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int32ListMapResult)
	success, err := handler.(echo.TestService).EchoBool2Int32ListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int32ListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int32ListMapArgs()
}

func newTestServiceEchoBool2Int32ListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int32ListMapResult()
}

func echoBool2Int64ListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2Int64ListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2Int64ListMapResult)
	success, err := handler.(echo.TestService).EchoBool2Int64ListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2Int64ListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2Int64ListMapArgs()
}

func newTestServiceEchoBool2Int64ListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2Int64ListMapResult()
}

func echoBool2DoubleListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2DoubleListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2DoubleListMapResult)
	success, err := handler.(echo.TestService).EchoBool2DoubleListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2DoubleListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2DoubleListMapArgs()
}

func newTestServiceEchoBool2DoubleListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2DoubleListMapResult()
}

func echoBool2StringListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2StringListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2StringListMapResult)
	success, err := handler.(echo.TestService).EchoBool2StringListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2StringListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2StringListMapArgs()
}

func newTestServiceEchoBool2StringListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2StringListMapResult()
}

func echoBool2BinaryListMapHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.TestServiceEchoBool2BinaryListMapArgs)
	realResult := result.(*echo.TestServiceEchoBool2BinaryListMapResult)
	success, err := handler.(echo.TestService).EchoBool2BinaryListMap(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTestServiceEchoBool2BinaryListMapArgs() interface{} {
	return echo.NewTestServiceEchoBool2BinaryListMapArgs()
}

func newTestServiceEchoBool2BinaryListMapResult() interface{} {
	return echo.NewTestServiceEchoBool2BinaryListMapResult()
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

func (p *kClient) EchoBool(ctx context.Context, req bool) (r bool, err error) {
	var _args echo.TestServiceEchoBoolArgs
	_args.Req = req
	var _result echo.TestServiceEchoBoolResult
	if err = p.c.Call(ctx, "EchoBool", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoByte(ctx context.Context, req int8) (r int8, err error) {
	var _args echo.TestServiceEchoByteArgs
	_args.Req = req
	var _result echo.TestServiceEchoByteResult
	if err = p.c.Call(ctx, "EchoByte", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt16(ctx context.Context, req int16) (r int16, err error) {
	var _args echo.TestServiceEchoInt16Args
	_args.Req = req
	var _result echo.TestServiceEchoInt16Result
	if err = p.c.Call(ctx, "EchoInt16", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt32(ctx context.Context, req int32) (r int32, err error) {
	var _args echo.TestServiceEchoInt32Args
	_args.Req = req
	var _result echo.TestServiceEchoInt32Result
	if err = p.c.Call(ctx, "EchoInt32", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt64(ctx context.Context, req int64) (r int64, err error) {
	var _args echo.TestServiceEchoInt64Args
	_args.Req = req
	var _result echo.TestServiceEchoInt64Result
	if err = p.c.Call(ctx, "EchoInt64", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoDouble(ctx context.Context, req float64) (r float64, err error) {
	var _args echo.TestServiceEchoDoubleArgs
	_args.Req = req
	var _result echo.TestServiceEchoDoubleResult
	if err = p.c.Call(ctx, "EchoDouble", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoString(ctx context.Context, req string) (r string, err error) {
	var _args echo.TestServiceEchoStringArgs
	_args.Req = req
	var _result echo.TestServiceEchoStringResult
	if err = p.c.Call(ctx, "EchoString", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBinary(ctx context.Context, req []byte) (r []byte, err error) {
	var _args echo.TestServiceEchoBinaryArgs
	_args.Req = req
	var _result echo.TestServiceEchoBinaryResult
	if err = p.c.Call(ctx, "EchoBinary", &_args, &_result); err != nil {
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

func (p *kClient) EchoBoolList(ctx context.Context, req []bool) (r []bool, err error) {
	var _args echo.TestServiceEchoBoolListArgs
	_args.Req = req
	var _result echo.TestServiceEchoBoolListResult
	if err = p.c.Call(ctx, "EchoBoolList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoByteList(ctx context.Context, req []int8) (r []int8, err error) {
	var _args echo.TestServiceEchoByteListArgs
	_args.Req = req
	var _result echo.TestServiceEchoByteListResult
	if err = p.c.Call(ctx, "EchoByteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt16List(ctx context.Context, req []int16) (r []int16, err error) {
	var _args echo.TestServiceEchoInt16ListArgs
	_args.Req = req
	var _result echo.TestServiceEchoInt16ListResult
	if err = p.c.Call(ctx, "EchoInt16List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt32List(ctx context.Context, req []int32) (r []int32, err error) {
	var _args echo.TestServiceEchoInt32ListArgs
	_args.Req = req
	var _result echo.TestServiceEchoInt32ListResult
	if err = p.c.Call(ctx, "EchoInt32List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoInt64List(ctx context.Context, req []int64) (r []int64, err error) {
	var _args echo.TestServiceEchoInt64ListArgs
	_args.Req = req
	var _result echo.TestServiceEchoInt64ListResult
	if err = p.c.Call(ctx, "EchoInt64List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoDoubleList(ctx context.Context, req []float64) (r []float64, err error) {
	var _args echo.TestServiceEchoDoubleListArgs
	_args.Req = req
	var _result echo.TestServiceEchoDoubleListResult
	if err = p.c.Call(ctx, "EchoDoubleList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoStringList(ctx context.Context, req []string) (r []string, err error) {
	var _args echo.TestServiceEchoStringListArgs
	_args.Req = req
	var _result echo.TestServiceEchoStringListResult
	if err = p.c.Call(ctx, "EchoStringList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBinaryList(ctx context.Context, req [][]byte) (r [][]byte, err error) {
	var _args echo.TestServiceEchoBinaryListArgs
	_args.Req = req
	var _result echo.TestServiceEchoBinaryListResult
	if err = p.c.Call(ctx, "EchoBinaryList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2BoolMap(ctx context.Context, req map[bool]bool) (r map[bool]bool, err error) {
	var _args echo.TestServiceEchoBool2BoolMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2BoolMapResult
	if err = p.c.Call(ctx, "EchoBool2BoolMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2ByteMap(ctx context.Context, req map[bool]int8) (r map[bool]int8, err error) {
	var _args echo.TestServiceEchoBool2ByteMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2ByteMapResult
	if err = p.c.Call(ctx, "EchoBool2ByteMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int16Map(ctx context.Context, req map[bool]int16) (r map[bool]int16, err error) {
	var _args echo.TestServiceEchoBool2Int16MapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int16MapResult
	if err = p.c.Call(ctx, "EchoBool2Int16Map", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int32Map(ctx context.Context, req map[bool]int32) (r map[bool]int32, err error) {
	var _args echo.TestServiceEchoBool2Int32MapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int32MapResult
	if err = p.c.Call(ctx, "EchoBool2Int32Map", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int64Map(ctx context.Context, req map[bool]int64) (r map[bool]int64, err error) {
	var _args echo.TestServiceEchoBool2Int64MapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int64MapResult
	if err = p.c.Call(ctx, "EchoBool2Int64Map", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2DoubleMap(ctx context.Context, req map[bool]float64) (r map[bool]float64, err error) {
	var _args echo.TestServiceEchoBool2DoubleMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2DoubleMapResult
	if err = p.c.Call(ctx, "EchoBool2DoubleMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2StringMap(ctx context.Context, req map[bool]string) (r map[bool]string, err error) {
	var _args echo.TestServiceEchoBool2StringMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2StringMapResult
	if err = p.c.Call(ctx, "EchoBool2StringMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte) (r map[bool][]byte, err error) {
	var _args echo.TestServiceEchoBool2BinaryMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2BinaryMapResult
	if err = p.c.Call(ctx, "EchoBool2BinaryMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool) (r map[bool][]bool, err error) {
	var _args echo.TestServiceEchoBool2BoolListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2BoolListMapResult
	if err = p.c.Call(ctx, "EchoBool2BoolListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8) (r map[bool][]int8, err error) {
	var _args echo.TestServiceEchoBool2ByteListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2ByteListMapResult
	if err = p.c.Call(ctx, "EchoBool2ByteListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16) (r map[bool][]int16, err error) {
	var _args echo.TestServiceEchoBool2Int16ListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int16ListMapResult
	if err = p.c.Call(ctx, "EchoBool2Int16ListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32) (r map[bool][]int32, err error) {
	var _args echo.TestServiceEchoBool2Int32ListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int32ListMapResult
	if err = p.c.Call(ctx, "EchoBool2Int32ListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64) (r map[bool][]int64, err error) {
	var _args echo.TestServiceEchoBool2Int64ListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2Int64ListMapResult
	if err = p.c.Call(ctx, "EchoBool2Int64ListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64) (r map[bool][]float64, err error) {
	var _args echo.TestServiceEchoBool2DoubleListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2DoubleListMapResult
	if err = p.c.Call(ctx, "EchoBool2DoubleListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2StringListMap(ctx context.Context, req map[bool][]string) (r map[bool][]string, err error) {
	var _args echo.TestServiceEchoBool2StringListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2StringListMapResult
	if err = p.c.Call(ctx, "EchoBool2StringListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte) (r map[bool][][]byte, err error) {
	var _args echo.TestServiceEchoBool2BinaryListMapArgs
	_args.Req = req
	var _result echo.TestServiceEchoBool2BinaryListMapResult
	if err = p.c.Call(ctx, "EchoBool2BinaryListMap", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
