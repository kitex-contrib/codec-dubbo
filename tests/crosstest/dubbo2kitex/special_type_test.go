package dubbo2kitex

import (
	"context"
	"testing"
)

func TestEchoBinary(t *testing.T) {
	var req = []byte{'1', '2'}
	resp, err := cli.EchoBinary(context.Background(), req)
	assertEcho(t, err, req, resp)
}
