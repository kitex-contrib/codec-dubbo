package main

import (
	"context"
	"errors"

	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
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

// EchoByte implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoByte(ctx context.Context, req int8) (resp int8, err error) {
	return req, nil
}

// Echo implements the TestServiceImpl interface.
func (s *TestServiceImpl) Echo(ctx context.Context, req *echo.EchoRequest) (resp *echo.EchoResponse, err error) {
	return &echo.EchoResponse{Int32: req.Int32}, nil
}

// EchoBool implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool(ctx context.Context, req bool) (resp bool, err error) {
	return req, nil
}

// EchoInt16 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt16(ctx context.Context, req int16) (resp int16, err error) {
	return req, nil
}

// EchoInt32 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt32(ctx context.Context, req int32) (resp int32, err error) {
	return req, nil
}

// EchoInt64 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt64(ctx context.Context, req int64) (resp int64, err error) {
	return req, nil
}

// EchoDouble implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoDouble(ctx context.Context, req float64) (resp float64, err error) {
	return req, nil
}

// EchoString implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoString(ctx context.Context, req string) (resp string, err error) {
	return req, nil
}

// EchoBinary implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBinary(ctx context.Context, req []byte) (resp []byte, err error) {
	return req, nil
}

// EchoBoolList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBoolList(ctx context.Context, req []bool) (resp []bool, err error) {
	return req, nil
}

// EchoByteList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoByteList(ctx context.Context, req []int8) (resp []int8, err error) {
	return req, nil
}

// EchoInt16List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt16List(ctx context.Context, req []int16) (resp []int16, err error) {
	return req, nil
}

// EchoInt32List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt32List(ctx context.Context, req []int32) (resp []int32, err error) {
	return req, nil
}

// EchoInt64List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoInt64List(ctx context.Context, req []int64) (resp []int64, err error) {
	return req, nil
}

// EchoDoubleList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoDoubleList(ctx context.Context, req []float64) (resp []float64, err error) {
	return req, nil
}

// EchoStringList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoStringList(ctx context.Context, req []string) (resp []string, err error) {
	return req, nil
}

// EchoBinaryList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBinaryList(ctx context.Context, req [][]byte) (resp [][]byte, err error) {
	return req, nil
}
