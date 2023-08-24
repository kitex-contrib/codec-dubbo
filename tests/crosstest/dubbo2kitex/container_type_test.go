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
