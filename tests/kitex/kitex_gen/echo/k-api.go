package echo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	hessian2Api "github.com/kitex-contrib/codec-dubbo/pkg/hessian2"

	hessian2 "github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)

	_ hessian2.Message = (*TestServiceEchoIntArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoIntResult)(nil)
	_ hessian2.Message = (*TestServiceEchoByteArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoByteResult)(nil)
	_ hessian2.Message = (*TestServiceEchoBytesArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoBytesResult)(nil)
	_ hessian2.Message = (*TestServiceEchoInt8Args)(nil)
	_ hessian2.Message = (*TestServiceEchoInt8Result)(nil)
	_ hessian2.Message = (*TestServiceEchoInt8sArgs)(nil)
	_ hessian2.Message = (*TestServiceEchoInt8sResult)(nil)
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
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
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
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
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

type TestServiceEchoByteArgs struct {
	Req byte `json:"req"`
}

func (p *TestServiceEchoByteArgs) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoByteArgs) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoByteArgs() *TestServiceEchoByteArgs {
	return &TestServiceEchoByteArgs{}
}

func (p *TestServiceEchoByteArgs) GetReq() (v byte) {
	return p.Req
}

func (p *TestServiceEchoByteArgs) SetReq(val byte) {
	p.Req = val
}

func (p *TestServiceEchoByteArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoByteArgs(%+v)", *p)
}

func (p *TestServiceEchoByteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoByteResult struct {
	Success *byte `json:"success,omitempty"`
}

func (p *TestServiceEchoByteResult) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoByteResult) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoByteResult() *TestServiceEchoByteResult {
	return &TestServiceEchoByteResult{}
}

var TestServiceEchoByteResult_Success_DEFAULT byte

func (p *TestServiceEchoByteResult) GetSuccess() (v byte) {
	if !p.IsSetSuccess() {
		return TestServiceEchoByteResult_Success_DEFAULT
	}
	return *p.Success
}

func (p *TestServiceEchoByteResult) SetSuccess(x interface{}) {
	p.Success = x.(*byte)
}

func (p *TestServiceEchoByteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoByteResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoByteResult(%+v)", *p)
}

func (p *TestServiceEchoByteResult) GetResult() interface{} {
	return p.Success
}

type TestServiceEchoBytesArgs struct {
	Req []byte `json:"req"`
}

func (p *TestServiceEchoBytesArgs) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoBytesArgs) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoBytesArgs() *TestServiceEchoBytesArgs {
	return &TestServiceEchoBytesArgs{}
}

func (p *TestServiceEchoBytesArgs) GetReq() (v []byte) {
	return p.Req
}

func (p *TestServiceEchoBytesArgs) SetReq(val []byte) {
	p.Req = val
}

func (p *TestServiceEchoBytesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoBytesArgs(%+v)", *p)
}

func (p *TestServiceEchoBytesArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoBytesResult struct {
	Success []byte `json:"success,omitempty"`
}

func (p *TestServiceEchoBytesResult) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoBytesResult) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoBytesResult() *TestServiceEchoBytesResult {
	return &TestServiceEchoBytesResult{}
}

var TestServiceEchoBytesResult_Success_DEFAULT []byte

func (p *TestServiceEchoBytesResult) GetSuccess() (v []byte) {
	if !p.IsSetSuccess() {
		return TestServiceEchoBytesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *TestServiceEchoBytesResult) SetSuccess(x interface{}) {
	p.Success = x.([]byte)
}

func (p *TestServiceEchoBytesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoBytesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoBytesResult(%+v)", *p)
}

func (p *TestServiceEchoBytesResult) GetResult() interface{} {
	return p.Success
}

type TestServiceEchoInt8Args struct {
	Req int8 `json:"req"`
}

func (p *TestServiceEchoInt8Args) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoInt8Args) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoInt8Args() *TestServiceEchoInt8Args {
	return &TestServiceEchoInt8Args{}
}

func (p *TestServiceEchoInt8Args) GetReq() (v int8) {
	return p.Req
}

func (p *TestServiceEchoInt8Args) SetReq(val int8) {
	p.Req = val
}

func (p *TestServiceEchoInt8Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoInt8Args(%+v)", *p)
}

func (p *TestServiceEchoInt8Args) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoInt8Result struct {
	Success *int8 `json:"success,omitempty"`
}

func (p *TestServiceEchoInt8Result) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoInt8Result) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoInt8Result() *TestServiceEchoInt8Result {
	return &TestServiceEchoInt8Result{}
}

var TestServiceEchoInt8Result_Success_DEFAULT int8

func (p *TestServiceEchoInt8Result) GetSuccess() (v int8) {
	if !p.IsSetSuccess() {
		return TestServiceEchoInt8Result_Success_DEFAULT
	}
	return *p.Success
}

func (p *TestServiceEchoInt8Result) SetSuccess(x interface{}) {
	p.Success = x.(*int8)
}

func (p *TestServiceEchoInt8Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoInt8Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoInt8Result(%+v)", *p)
}

func (p *TestServiceEchoInt8Result) GetResult() interface{} {
	return p.Success
}

type TestServiceEchoInt8sArgs struct {
	Req []int8 `json:"req"`
}

func (p *TestServiceEchoInt8sArgs) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Req)
}

func (p *TestServiceEchoInt8sArgs) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoInt8sArgs() *TestServiceEchoInt8sArgs {
	return &TestServiceEchoInt8sArgs{}
}

func (p *TestServiceEchoInt8sArgs) GetReq() (v []int8) {
	return p.Req
}

func (p *TestServiceEchoInt8sArgs) SetReq(val []int8) {
	p.Req = val
}

func (p *TestServiceEchoInt8sArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoInt8sArgs(%+v)", *p)
}

func (p *TestServiceEchoInt8sArgs) GetFirstArgument() interface{} {
	return p.Req
}

type TestServiceEchoInt8sResult struct {
	Success []int8 `json:"success,omitempty"`
}

func (p *TestServiceEchoInt8sResult) Encode(e hessian2.Encoder) error {
	return e.Encode(p.Success)
}

func (p *TestServiceEchoInt8sResult) Decode(d hessian2.Decoder) error {
	v, err := d.Decode()
	if err != nil {
		return err
	}
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
	return nil
}

func NewTestServiceEchoInt8sResult() *TestServiceEchoInt8sResult {
	return &TestServiceEchoInt8sResult{}
}

var TestServiceEchoInt8sResult_Success_DEFAULT []int8

func (p *TestServiceEchoInt8sResult) GetSuccess() (v []int8) {
	if !p.IsSetSuccess() {
		return TestServiceEchoInt8sResult_Success_DEFAULT
	}
	return p.Success
}

func (p *TestServiceEchoInt8sResult) SetSuccess(x interface{}) {
	p.Success = x.([]int8)
}

func (p *TestServiceEchoInt8sResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceEchoInt8sResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceEchoInt8sResult(%+v)", *p)
}

func (p *TestServiceEchoInt8sResult) GetResult() interface{} {
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
	if err = hessian2Api.ReflectResponse(v, &p.Req); err != nil {
		return err
	}
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
	if err = hessian2Api.ReflectResponse(v, &p.Success); err != nil {
		return err
	}
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
