package dubbo2kitex

import (
	"context"
	"fmt"
	"testing"
)

func TestEchoMethodA(t *testing.T) {
	var req bool = true
	resp, err := cli.EchoMethodA(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("A:%v", req), resp)
}

func TestEchoMethodB(t *testing.T) {
	var req int32 = 1
	resp, err := cli.EchoMethodB(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("B:%v", req), resp)
}

func TestEchoMethodC(t *testing.T) {
	var req int32 = 1
	resp, err := cli.EchoMethodC(context.Background(), req)
	assertEcho(t, err, fmt.Sprintf("C:%v", req), resp)
}

func TestEchoMethodD(t *testing.T) {
	var req1 bool = true
	var req2 int32 = 1
	resp, err := cli.EchoMethodD(context.Background(), req1, req2)
	assertEcho(t, err, fmt.Sprintf("D:%v,%v", req1, req2), resp)
}
