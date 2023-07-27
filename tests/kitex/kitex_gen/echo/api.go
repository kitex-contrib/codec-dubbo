package echo

import (
	"context"
	"fmt"
)

type EchoRequest struct {
	Int32 int32 `json:"int32"`
}

func (p *EchoRequest) JavaClassName() string {
	return "kitex.echo.EchoRequest"
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{}
}

func (p *EchoRequest) GetInt32() (v int32) {
	return p.Int32
}

func (p *EchoRequest) SetInt32(val int32) {
	p.Int32 = val
}

func (p *EchoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRequest(%+v)", *p)
}

type EchoResponse struct {
	Int32 int32 `json:"int32"`
}

func (p *EchoResponse) JavaClassName() string {
	return "kitex.echo.EchoResponse"
}

func NewEchoResponse() *EchoResponse {
	return &EchoResponse{}
}

func (p *EchoResponse) GetInt32() (v int32) {
	return p.Int32
}

func (p *EchoResponse) SetInt32(val int32) {
	p.Int32 = val
}

func (p *EchoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type TestService interface {
	EchoInt(ctx context.Context, req int32) (r int32, err error)
	Echo(ctx context.Context, req *EchoRequest) (r *EchoResponse, err error)
}
