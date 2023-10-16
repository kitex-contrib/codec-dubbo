package user

import (
	"fmt"

	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi = []interface{}{
	&Request{},
	&User{},
}

func init() {
	hessian2.Register(objectsApi)
}

func GetUserServiceIDLAnnotations() map[string][]string {
	return map[string][]string{
		"JavaClassName": {"org.apache.dubbo.UserProvider"},
	}
}

func (p *Request) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (p *Request) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *Request) JavaClassName() string {
	return "org.apache.dubbo.Request"
}

func (p *User) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.ID)
	if err != nil {
		return err
	}

	err = e.Encode(p.Name)
	if err != nil {
		return err
	}

	err = e.Encode(p.Age)
	if err != nil {
		return err
	}

	return nil
}

func (p *User) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Age)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

func (p *UserServiceGetUserArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserServiceGetUserArgs) Decode(d codec.Decoder) error {
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

func (p *UserServiceGetUserResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserServiceGetUserResult) Decode(d codec.Decoder) error {
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
