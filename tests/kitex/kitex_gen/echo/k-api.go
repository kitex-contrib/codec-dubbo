package echo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	hessian2 "github.com/kitex-contrib/codec-hessian2/pkg/iface"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)

	_ hessian2.Message = (*TestServiceEchoIntArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoIntResult)(nil)
	_ hessian2.Message = (*TestServiceEchoArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoResult)(nil)
)

type TestServiceEchoIntArgs struct {
	Req int32 `json:"req"`
}

func (p *TestServiceEchoIntArgs) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoIntArgs) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	i, ok := v.(int32)
	if !ok {
		return fmt.Errorf("invalid data type: %T", v)
	}
	p.Req = i
	return nil
}

func NewTestServiceEchoIntArgs() *TestServiceEchoIntArgs {
	return &TestServiceEchoIntArgs{}
}

func (p *TestServiceEchoIntArgs) GetReq() (v int32) {
	return p.Req
}

func (p *TestServiceEchoIntArgs) SetReq(val int32) {
	p.Req = val
}

func (p *TestServiceEchoIntArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoIntArgs(%+v)", *p)
}

func (p *TestServiceEchoIntArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoIntResult struct {
	Success *int32 `json:"success,omitempty"`
}

func (p *TestServiceEchoIntResult) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoIntResult) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	i, ok := v.(int32)
	if !ok {
		return fmt.Errorf("invalid data type: %T", v)
	}
	p.Success = &i
	return nil
}

func NewTestServiceEchoIntResult() *TestServiceEchoIntResult {
	return &TestServiceEchoIntResult{}
}

var TestServiceEchoIntResult_Success_DEFAULT int32

func (p *TestServiceEchoIntResult) GetSuccess() (v int32) {
	if !p.IsSetSuccess() {
		return TestServiceEchoIntResult_Success_DEFAULT
	}
	return *p.Success
}

func (p *TestServiceEchoIntResult) SetSuccess(x interface{}) {
	p.Success = x.(*int32)
}

func (p *TestServiceEchoIntResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoIntResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoIntResult(%+v)", *p)
}

func (p *TestServiceEchoIntResult) GetResult() interface{} {
	return p.Success
}

type TestServiceEchoArgs struct {
	Req *EchoRequest `json:"req"`
}

func (p *TestServiceEchoArgs) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoArgs) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	i, ok := v.(*EchoRequest)
	if !ok {
		return fmt.Errorf("invalid data type: %T", v)
	}
	p.Req = i
	return nil
}

func NewTestServiceEchoArgs() *TestServiceEchoArgs {
	return &TestServiceEchoArgs{}
}

var TestServiceEchoArgs_Req_DEFAULT *EchoRequest

func (p *TestServiceEchoArgs) GetReq() (v *EchoRequest) {
	if !p.IsSetReq() {
		return TestServiceEchoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *TestServiceEchoArgs) SetReq(val *EchoRequest) {
	p.Req = val
}

func (p *TestServiceEchoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *TestServiceEchoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoArgs(%+v)", *p)
}

func (p *TestServiceEchoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoResult struct {
	Success *EchoResponse `json:"success,omitempty"`
}

func (p *TestServiceEchoResult) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoResult) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	i, ok := v.(*EchoResponse)
	if !ok {
		return fmt.Errorf("invalid data type: %T", v)
	}
	p.Success = i
	return nil
}

func NewTestServiceEchoResult() *TestServiceEchoResult {
	return &TestServiceEchoResult{}
}

var TestServiceEchoResult_Success_DEFAULT *EchoResponse

func (p *TestServiceEchoResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return TestServiceEchoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *TestServiceEchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}

func (p *TestServiceEchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoResult(%+v)", *p)
}

func (p *TestServiceEchoResult) GetResult() interface{} {
	return p.Success
}
