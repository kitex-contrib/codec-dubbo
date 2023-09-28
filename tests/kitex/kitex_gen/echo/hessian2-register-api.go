package echo

import (
	"fmt"

	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi = []interface{}{
	&EchoRequest{},
	&EchoResponse{},
}

func init() {
	hessian2.Register(objectsApi)
}

func (p *EchoRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Int32)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Int32)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoRequest) JavaClassName() string {
	return "kitex.echo.EchoRequest"
}

func (p *EchoResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Int32)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Int32)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoResponse) JavaClassName() string {
	return "kitex.echo.EchoResponse"
}
