package main

import (
	"context"

	hessian "github.com/apache/dubbo-go-hessian2"
)

func init() {
	hessian.RegisterPOJO(&Request{})
	hessian.RegisterPOJO(&User{})
}

type Request struct {
	Name string
}

func (r *Request) JavaClassName() string {
	return "org.apache.dubbo.Request"
}

type User struct {
	ID   string
	Name string
	Age  int32
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

type UserProvider struct{}

func (up *UserProvider) GetUser(ctx context.Context, req *Request) (*User, error) {
	return &User{
		ID:   "12345",
		Name: "Hello " + req.Name,
		Age:  21,
	}, nil
}
