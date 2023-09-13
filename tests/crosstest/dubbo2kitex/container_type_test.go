package dubbo2kitex

import (
	"context"
	"testing"
)

func TestEchoBoolList(t *testing.T) {
	req := []bool{true, false}
	resp, err := cli.EchoBoolList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// todo(DMwangnima): enhance hessian2.ReflectResponse to support reflecting []int8
//func TestEchoByteList(t *testing.T) {
//	var req = []int8{1, 2}
//	resp, err := cli.EchoByteList(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// todo(DMwangnima): enhance hessian2.ReflectResponse to support reflecting []int16
//func TestEchoInt16List(t *testing.T) {
//	var req = []int16{1, 2}
//	resp, err := cli.EchoInt16List(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

func TestEchoInt32List(t *testing.T) {
	req := []int32{1, 2}
	resp, err := cli.EchoInt32List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64List(t *testing.T) {
	req := []int64{1, 2}
	resp, err := cli.EchoInt64List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDoubleList(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli.EchoDoubleList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoStringList(t *testing.T) {
	req := []string{"1", "2"}
	resp, err := cli.EchoStringList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-go hessian2 does not support [][]byte, please refer to github.com/apache/dubbo-go-hessian2/list.go
//func TestEchoBinaryList(t *testing.T) {
//	var req = [][]byte{{'1'}, {'2'}}
//	resp, err := cli.EchoBinaryList(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// We have supported other map types (refer to /pkg/hessian2/response_test),
// but ReflectResponse in dubbo-go side could not support. As a result dubbo-go could not parse
// map types response correctly.
// We would finish this part of tests in dubbo -> kitex.

//func TestEchoBool2BoolMap(t *testing.T) {
//	req := map[bool]bool{
//		true: true,
//	}
//	resp, err := cli.EchoBool2BoolMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2ByteMap(t *testing.T) {
//	req := map[bool]byte{
//		true: 1,
//	}
//	resp, err := cli.EchoBool2ByteMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int16Map(t *testing.T) {
//	req := map[bool]int16{
//		true: 1,
//	}
//	resp, err := cli.EchoBool2Int16Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int32Map(t *testing.T) {
//	req := map[bool]int32{
//		true: 1,
//	}
//	resp, err := cli.EchoBool2Int32Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int64Map(t *testing.T) {
//	req := map[bool]int64{
//		true: 1,
//	}
//	resp, err := cli.EchoBool2Int64Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2DoubleMap(t *testing.T) {
//	req := map[bool]float64{
//		true: 12.34,
//	}
//	resp, err := cli.EchoBool2DoubleMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2StringMap(t *testing.T) {
//	req := map[bool]string{
//		true: "1",
//	}
//	resp, err := cli.EchoBool2StringMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2BinaryMap(t *testing.T) {
//	req := map[bool][]byte{
//		true: {'1', '2'},
//	}
//	resp, err := cli.EchoBool2BinaryMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}
