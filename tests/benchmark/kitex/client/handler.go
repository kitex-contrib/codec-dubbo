package main

import (
	"context"

	benchuser "github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

// BenchmarkServiceImpl implements the last service interface defined in the IDL.
type BenchmarkServiceImpl struct {
	cli userservice.Client
}

// GetUser implements the BenchmarkServiceImpl interface.
func (s *BenchmarkServiceImpl) GetUser(ctx context.Context, req *benchuser.Request) (resp *benchuser.User, err error) {
	userResp, err := s.cli.GetUser(ctx, &user.Request{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &benchuser.User{
		ID:   userResp.ID,
		Name: userResp.Name,
		Age:  userResp.Age,
	}, nil
}
