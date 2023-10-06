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
	&EchoMultiBoolResponse{},
	&EchoMultiByteResponse{},
	&EchoMultiInt16Response{},
	&EchoMultiInt32Response{},
	&EchoMultiInt64Response{},
	&EchoMultiDoubleResponse{},
	&EchoMultiStringResponse{},
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

func (p *EchoMultiBoolResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiBoolResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiBoolResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiBoolResponse"
}

func (p *EchoMultiByteResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiByteResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiByteResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiByteResponse"
}

func (p *EchoMultiInt16Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt16Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt16Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt16Response"
}

func (p *EchoMultiInt32Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt32Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt32Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt32Response"
}

func (p *EchoMultiInt64Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt64Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt64Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt64Response"
}

func (p *EchoMultiDoubleResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiDoubleResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiDoubleResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiDoubleResponse"
}

func (p *EchoMultiStringResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiStringResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiStringResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiStringResponse"
}
