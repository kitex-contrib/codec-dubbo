package main

import (
	"context"
	"errors"

	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// EchoInt implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt(ctx context.Context, req int32) (resp int32, err error) {
	// for exception test
	if req == 400 {
		return 0, errors.New("EchoInt failed without reason")
	}

	return req, nil
}

func (s *TestServiceImpl) EchoByte(ctx context.Context, req byte) (resp byte, err error) {
	return req, nil
}

func (s *TestServiceImpl) EchoBytes(ctx context.Context, req []byte) (resp []byte, err error) {
	return req, nil
}

func (s *TestServiceImpl) EchoInt8(ctx context.Context, req int8) (resp int8, err error) {
	return req, nil
}

func (s *TestServiceImpl) EchoInt8s(ctx context.Context, req []int8) (resp []int8, err error) {
	return req, nil
}

// Echo implements the TestServiceImpl interface.
func (s *TestServiceImpl) Echo(ctx context.Context, req *echo.EchoRequest) (resp *echo.EchoResponse, err error) {
	return &echo.EchoResponse{
		Int32: req.Int32,
	}, nil
}
