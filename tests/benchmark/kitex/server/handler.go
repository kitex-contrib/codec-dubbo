package main

import (
	"context"

	user "github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.Request) (resp *user.User, err error) {
	return &user.User{
		ID:   "12345",
		Name: "Hello " + req.Name,
		Age:  21,
	}, nil
}
