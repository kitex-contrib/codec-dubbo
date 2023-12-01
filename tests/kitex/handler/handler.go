package handler

import (
	"context"
	"errors"
	"fmt"
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

// EchoFloat implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoFloat(ctx context.Context, req float64) (resp float64, err error) {
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

// EchoFloatList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoFloatList(ctx context.Context, req []float64) (resp []float64, err error) {
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

// EchoBool2BoolMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2BoolMap(ctx context.Context, req map[bool]bool) (resp map[bool]bool, err error) {
	return req, nil
}

// EchoBool2ByteMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2ByteMap(ctx context.Context, req map[bool]int8) (resp map[bool]int8, err error) {
	return req, nil
}

// EchoBool2Int16Map implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int16Map(ctx context.Context, req map[bool]int16) (resp map[bool]int16, err error) {
	return req, nil
}

// EchoBool2Int32Map implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int32Map(ctx context.Context, req map[bool]int32) (resp map[bool]int32, err error) {
	return req, nil
}

// EchoBool2Int64Map implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int64Map(ctx context.Context, req map[bool]int64) (resp map[bool]int64, err error) {
	return req, nil
}

// EchoBool2FloatMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2FloatMap(ctx context.Context, req map[bool]float64) (resp map[bool]float64, err error) {
	return req, nil
}

// EchoBool2DoubleMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2DoubleMap(ctx context.Context, req map[bool]float64) (resp map[bool]float64, err error) {
	return req, nil
}

// EchoBool2StringMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2StringMap(ctx context.Context, req map[bool]string) (resp map[bool]string, err error) {
	return req, nil
}

// EchoBool2BinaryMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte) (resp map[bool][]byte, err error) {
	return req, nil
}

// EchoBool2BoolListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool) (resp map[bool][]bool, err error) {
	return req, nil
}

// EchoBool2ByteListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8) (resp map[bool][]int8, err error) {
	return req, nil
}

// EchoBool2Int16ListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16) (resp map[bool][]int16, err error) {
	return req, nil
}

// EchoBool2Int32ListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32) (resp map[bool][]int32, err error) {
	return req, nil
}

// EchoBool2Int64ListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64) (resp map[bool][]int64, err error) {
	return req, nil
}

// EchoBool2FloatListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2FloatListMap(ctx context.Context, req map[bool][]float64) (resp map[bool][]float64, err error) {
	return req, nil
}

// EchoBool2DoubleListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64) (resp map[bool][]float64, err error) {
	return req, nil
}

// EchoBool2StringListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2StringListMap(ctx context.Context, req map[bool][]string) (resp map[bool][]string, err error) {
	return req, nil
}

// EchoBool2BinaryListMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte) (resp map[bool][][]byte, err error) {
	return req, nil
}

// EchoMultiBool implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool) (resp *echo.EchoMultiBoolResponse, err error) {
	return &echo.EchoMultiBoolResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiByte implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8) (resp *echo.EchoMultiByteResponse, err error) {
	return &echo.EchoMultiByteResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiInt16 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16) (resp *echo.EchoMultiInt16Response, err error) {
	return &echo.EchoMultiInt16Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiInt32 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32) (resp *echo.EchoMultiInt32Response, err error) {
	return &echo.EchoMultiInt32Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiInt64 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64) (resp *echo.EchoMultiInt64Response, err error) {
	return &echo.EchoMultiInt64Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiFloat implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (resp *echo.EchoMultiFloatResponse, err error) {
	return &echo.EchoMultiFloatResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiDouble implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (resp *echo.EchoMultiDoubleResponse, err error) {
	return &echo.EchoMultiDoubleResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiString implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiString(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string) (resp *echo.EchoMultiStringResponse, err error) {
	return &echo.EchoMultiStringResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoBaseBool implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseBool(ctx context.Context, req bool) (resp bool, err error) {
	return req, nil
}

// EchoBaseByte implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseByte(ctx context.Context, req int8) (resp int8, err error) {
	return req, nil
}

// EchoBaseInt16 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt16(ctx context.Context, req int16) (resp int16, err error) {
	return req, nil
}

// EchoBaseInt32 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt32(ctx context.Context, req int32) (resp int32, err error) {
	return req, nil
}

// EchoBaseInt64 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt64(ctx context.Context, req int64) (resp int64, err error) {
	return req, nil
}

// EchoBaseFloat implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseFloat(ctx context.Context, req float64) (resp float64, err error) {
	return req, nil
}

// EchoBaseDouble implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseDouble(ctx context.Context, req float64) (resp float64, err error) {
	return req, nil
}

// EchoBaseBoolList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseBoolList(ctx context.Context, req []bool) (resp []bool, err error) {
	return req, nil
}

// EchoBaseByteList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseByteList(ctx context.Context, req []int8) (resp []int8, err error) {
	return req, nil
}

// EchoBaseInt16List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt16List(ctx context.Context, req []int16) (resp []int16, err error) {
	return req, nil
}

// EchoBaseInt32List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt32List(ctx context.Context, req []int32) (resp []int32, err error) {
	return req, nil
}

// EchoBaseInt64List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseInt64List(ctx context.Context, req []int64) (resp []int64, err error) {
	return req, nil
}

// EchoBaseFloatList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseFloatList(ctx context.Context, req []float64) (resp []float64, err error) {
	return req, nil
}

// EchoBaseDoubleList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBaseDoubleList(ctx context.Context, req []float64) (resp []float64, err error) {
	return req, nil
}

// EchoBool2BoolBaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2BoolBaseMap(ctx context.Context, req map[bool]bool) (resp map[bool]bool, err error) {
	return req, nil
}

// EchoBool2ByteBaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2ByteBaseMap(ctx context.Context, req map[bool]int8) (resp map[bool]int8, err error) {
	return req, nil
}

// EchoBool2Int16BaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int16BaseMap(ctx context.Context, req map[bool]int16) (resp map[bool]int16, err error) {
	return req, nil
}

// EchoBool2Int32BaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int32BaseMap(ctx context.Context, req map[bool]int32) (resp map[bool]int32, err error) {
	return req, nil
}

// EchoBool2Int64BaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2Int64BaseMap(ctx context.Context, req map[bool]int64) (resp map[bool]int64, err error) {
	return req, nil
}

// EchoBool2FloatBaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2FloatBaseMap(ctx context.Context, req map[bool]float64) (resp map[bool]float64, err error) {
	return req, nil
}

// EchoBool2DoubleBaseMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoBool2DoubleBaseMap(ctx context.Context, req map[bool]float64) (resp map[bool]float64, err error) {
	return req, nil
}

// EchoMultiBaseBool implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool) (resp *echo.EchoMultiBoolResponse, err error) {
	return &echo.EchoMultiBoolResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseByte implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8) (resp *echo.EchoMultiByteResponse, err error) {
	return &echo.EchoMultiByteResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseInt16 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16) (resp *echo.EchoMultiInt16Response, err error) {
	return &echo.EchoMultiInt16Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseInt32 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32) (resp *echo.EchoMultiInt32Response, err error) {
	return &echo.EchoMultiInt32Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseInt64 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64) (resp *echo.EchoMultiInt64Response, err error) {
	return &echo.EchoMultiInt64Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseFloat implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (resp *echo.EchoMultiFloatResponse, err error) {
	return &echo.EchoMultiFloatResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMultiBaseDouble implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMultiBaseDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (resp *echo.EchoMultiDoubleResponse, err error) {
	return &echo.EchoMultiDoubleResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// EchoMethodA implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMethodA(ctx context.Context, req bool) (resp string, err error) {
	return fmt.Sprintf("A:%v", req), nil
}

// EchoMethodB implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMethodB(ctx context.Context, req int32) (resp string, err error) {
	return fmt.Sprintf("B:%v", req), nil
}

// EchoMethodC implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMethodC(ctx context.Context, req int32) (resp string, err error) {
	return fmt.Sprintf("C:%v", req), nil
}

// EchoMethodD implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoMethodD(ctx context.Context, req1 bool, req2 int32) (resp string, err error) {
	return fmt.Sprintf("D:%v,%v", req1, req2), nil
}

// EchoOptionalBool implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalBool(ctx context.Context, req bool) (resp bool, err error) {
	return req, nil
}

// EchoOptionalInt32 implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalInt32(ctx context.Context, req int32) (resp int32, err error) {
	return req, nil
}

// EchoOptionalString implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalString(ctx context.Context, req string) (resp string, err error) {
	return req, nil
}

// EchoOptionalBoolList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalBoolList(ctx context.Context, req []bool) (resp []bool, err error) {
	return req, nil
}

// EchoOptionalInt32List implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalInt32List(ctx context.Context, req []int32) (resp []int32, err error) {
	return req, nil
}

// EchoOptionalStringList implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalStringList(ctx context.Context, req []string) (resp []string, err error) {
	return req, nil
}

// EchoOptionalBool2BoolMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalBool2BoolMap(ctx context.Context, req map[bool]bool) (resp map[bool]bool, err error) {
	return req, nil
}

// EchoOptionalBool2Int32Map implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalBool2Int32Map(ctx context.Context, req map[bool]int32) (resp map[bool]int32, err error) {
	return req, nil
}

// EchoOptionalBool2StringMap implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalBool2StringMap(ctx context.Context, req map[bool]string) (resp map[bool]string, err error) {
	return req, nil
}

// EchoOptionalStruct implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalStruct(ctx context.Context, req *echo.EchoOptionalStructRequest) (resp *echo.EchoOptionalStructResponse, err error) {
	return nil, nil
}

// EchoOptionalMultiBoolRequest implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiBoolRequest(ctx context.Context, req *echo.EchoOptionalMultiBoolRequest) (resp bool, err error) {
	return req.GetBasicReq(), nil
}

// EchoOptionalMultiInt32Request implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiInt32Request(ctx context.Context, req *echo.EchoOptionalMultiInt32Request) (resp int32, err error) {
	return req.GetBasicReq(), nil
}

// EchoOptionalMultiStringRequest implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiStringRequest(ctx context.Context, req *echo.EchoOptionalMultiStringRequest) (resp string, err error) {
	return req.GetBaseReq(), nil
}

// EchoOptionalMultiBoolResponse implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiBoolResponse(ctx context.Context, req bool) (resp *echo.EchoOptionalMultiBoolResponse, err error) {
	return &echo.EchoOptionalMultiBoolResponse{
		BasicResp: nil,
		PackResp:  nil,
		ListResp:  nil,
		MapResp:   nil,
	}, nil
}

// EchoOptionalMultiInt32Response implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiInt32Response(ctx context.Context, req int32) (resp *echo.EchoOptionalMultiInt32Response, err error) {
	return &echo.EchoOptionalMultiInt32Response{
		BasicResp: nil,
		PackResp:  nil,
		ListResp:  nil,
		MapResp:   nil,
	}, nil
}

// EchoOptionalMultiStringResponse implements the TestServiceImpl interface.
func (s *TestServiceImpl) EchoOptionalMultiStringResponse(ctx context.Context, req string) (resp *echo.EchoOptionalMultiStringResponse, err error) {
	return &echo.EchoOptionalMultiStringResponse{
		BaseResp: nil,
		ListResp: nil,
		MapResp:  nil,
	}, nil
}
