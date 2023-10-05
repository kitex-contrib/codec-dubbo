package main

import (
	"context"

	proxyuser "github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

// ProxyServiceImpl implements the last service interface defined in the IDL.
type ProxyServiceImpl struct {
	cli userservice.Client
}

// GetUser implements the ProxyServiceImpl interface.
func (s *ProxyServiceImpl) GetUser(ctx context.Context, req *proxyuser.Request) (resp *proxyuser.User, err error) {
	userResp, err := s.cli.GetUser(ctx, &user.Request{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &proxyuser.User{
		ID:   userResp.ID,
		Name: userResp.Name,
		Age:  userResp.Age,
	}, nil
}
