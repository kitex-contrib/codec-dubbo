package hello

import (
	"fmt"
	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2/enum"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi = []interface{}{
	&GreetRequest{},
	&GreetResponse{},
	KitexEnum_ONE,
	KitexEnum_TWO,
	KitexEnum_THREE,
	KitexEnum_FOUR,
	KitexEnum_FIVE,
}

func init() {
	hessian2.Register(objectsApi)
}

func GetGreetServiceIDLAnnotations() map[string][]string {
	return map[string][]string{}
}
func GetGreetEnumServiceIDLAnnotations() map[string][]string {
	return map[string][]string{}
}

func (p *GreetRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetRequest) JavaClassName() string {
	return "org.cloudwego.kitex.samples.api.GreetRequest"
}

func (p *GreetResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Resp)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Resp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetResponse) JavaClassName() string {
	return "org.cloudwego.kitex.samples.api.GreetResponse"
}

var KitexEnumValues = map[string]KitexEnum{
	"ONE":   KitexEnum_ONE,
	"TWO":   KitexEnum_TWO,
	"THREE": KitexEnum_THREE,
	"FOUR":  KitexEnum_FOUR,
	"FIVE":  KitexEnum_FIVE,
}

func (KitexEnum) JavaClassName() string {
	return "org.cloudwego.kitex.samples.enumeration.KitexEnum"
}

func (KitexEnum) EnumValue(s string) enum.JavaEnum {
	v, ok := KitexEnumValues[s]
	if ok {
		return enum.JavaEnum(v)
	}
	return enum.InvalidJavaEnum
}

func (p *GreetServiceGreetArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetServiceGreetArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetServiceGreetResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetServiceGreetResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetServiceGreetWithStructArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetServiceGreetWithStructArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetServiceGreetWithStructResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetServiceGreetWithStructResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetEnumServiceGreetEnumArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetEnumServiceGreetEnumArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *GreetEnumServiceGreetEnumResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *GreetEnumServiceGreetEnumResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}
