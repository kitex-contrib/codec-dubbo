/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dubbogo/gost/log/logger"

	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

// User transmission struct; for compatibility with java, field names should be consistent with Java class fields
type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User" // Should be same as Java class name for java compatibility
}

type EchoMultiBoolResponse struct {
	BaseResp bool
	ListResp []bool
	MapResp  map[bool]bool
}

func (resp *EchoMultiBoolResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiBoolResponse"
}

type EchoMultiByteResponse struct {
	BaseResp int8
	ListResp []int8
	MapResp  map[int8]int8
}

func (resp *EchoMultiByteResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiByteResponse"
}

type EchoMultiInt16Response struct {
	BaseResp int16
	ListResp []int16
	MapResp  map[int16]int16
}

func (resp *EchoMultiInt16Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt16Response"
}

type EchoMultiInt32Response struct {
	BaseResp int32
	ListResp []int32
	MapResp  map[int32]int32
}

func (resp *EchoMultiInt32Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt32Response"
}

type EchoMultiInt64Response struct {
	BaseResp int64
	ListResp []int64
	MapResp  map[int64]int64
}

func (resp *EchoMultiInt64Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt64Response"
}

type EchoMultiFloatResponse struct {
	BaseResp float64
	ListResp []float64
	MapResp  map[float64]float64
}

func (resp *EchoMultiFloatResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiFloatResponse"
}

type EchoMultiDoubleResponse struct {
	BaseResp float64
	ListResp []float64
	MapResp  map[float64]float64
}

func (resp *EchoMultiDoubleResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiDoubleResponse"
}

type EchoMultiStringResponse struct {
	BaseResp string
	ListResp []string
	MapResp  map[string]string
}

func (resp *EchoMultiStringResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiStringResponse"
}

var UserProviderClient = &UserProvider{} // client pointer

// UserProvider client interface
type UserProvider struct {
	// dubbo tag is necessary to map go function name to java function name
	GetUser func(ctx context.Context, req int32) (*User, error) //`dubbo:"getUser"`
	EchoInt func(ctx context.Context, req int32) (int32, error) //`dubbo:"echoInt"`

	// base types
	EchoBool   func(ctx context.Context, req bool) (bool, error)       //`dubbo:"echoBool"`
	EchoByte   func(ctx context.Context, req int8) (int8, error)       //`dubbo:"echoByte"`
	EchoInt16  func(ctx context.Context, req int16) (int16, error)     //`dubbo:"echoInt16"`
	EchoInt32  func(ctx context.Context, req int32) (int32, error)     //`dubbo:"echoInt32"`
	EchoInt64  func(ctx context.Context, req int64) (int64, error)     //`dubbo:"echoInt64"`
	EchoFloat  func(ctx context.Context, req float64) (float64, error) //`dubbo:"echoFloat"`
	EchoDouble func(ctx context.Context, req float64) (float64, error) //`dubbo:"echoDouble"`
	EchoString func(ctx context.Context, req string) (string, error)   //`dubbo:"echoString"`

	// special types
	EchoBinary func(ctx context.Context, req []byte) ([]byte, error) //`dubbo:"echoBinary"`

	// container list
	EchoBoolList   func(ctx context.Context, req []bool) ([]bool, error)       //`dubbo:"echoBoolList"`
	EchoByteList   func(ctx context.Context, req []int8) ([]int8, error)       //`dubbo:"echoByteList"`
	EchoInt16List  func(ctx context.Context, req []int16) ([]int16, error)     //`dubbo:"echoInt16List"`
	EchoInt32List  func(ctx context.Context, req []int32) ([]int32, error)     //`dubbo:"echoInt32List"`
	EchoInt64List  func(ctx context.Context, req []int64) ([]int64, error)     //`dubbo:"echoInt64List"`
	EchoFloatList  func(ctx context.Context, req []float64) ([]float64, error) //`dubbo:"echoFloatList"`
	EchoDoubleList func(ctx context.Context, req []float64) ([]float64, error) //`dubbo:"echoDoubleList"`
	EchoStringList func(ctx context.Context, req []string) ([]string, error)   //`dubbo:"echoStringList"`
	EchoBinaryList func(ctx context.Context, req [][]byte) ([][]byte, error)   //`dubbo:"echoBinaryList"`

	// container map
	EchoBool2BoolMap   func(ctx context.Context, req map[bool]bool) (map[bool]bool, error)       //`dubbo:"echoBool2BoolMap"`
	EchoBool2ByteMap   func(ctx context.Context, req map[bool]int8) (map[bool]int8, error)       //`dubbo:"echoBool2ByteMap"`
	EchoBool2Int16Map  func(ctx context.Context, req map[bool]int16) (map[bool]int16, error)     //`dubbo:"echoBool2Int16Map"`
	EchoBool2Int32Map  func(ctx context.Context, req map[bool]int32) (map[bool]int32, error)     //`dubbo:"echoBool2Int32Map"`
	EchoBool2Int64Map  func(ctx context.Context, req map[bool]int64) (map[bool]int64, error)     //`dubbo:"echoBool2Int64Map"`
	EchoBool2FloatMap  func(ctx context.Context, req map[bool]float64) (map[bool]float64, error) //`dubbo:"echoBool2FloatMap"`
	EchoBool2DoubleMap func(ctx context.Context, req map[bool]float64) (map[bool]float64, error) //`dubbo:"echoBool2DoubleMap"`
	EchoBool2StringMap func(ctx context.Context, req map[bool]string) (map[bool]string, error)   //`dubbo:"echoBool2StringMap"`
	EchoBool2BinaryMap func(ctx context.Context, req map[bool][]byte) (map[bool][]byte, error)   //`dubbo:"echoBool2BinaryMap"`

	EchoMultiBool   func(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool) (*EchoMultiBoolResponse, error)               //`dubbo:"echoMultiBool"`
	EchoMultiByte   func(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8) (*EchoMultiByteResponse, error)               //`dubbo:"echoMultiByte"`
	EchoMultiInt16  func(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16) (*EchoMultiInt16Response, error)          //`dubbo:"echoMultiInt16"`
	EchoMultiInt32  func(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32) (*EchoMultiInt32Response, error)          //`dubbo:"echoMultiInt32"`
	EchoMultiInt64  func(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64) (*EchoMultiInt64Response, error)          //`dubbo:"echoMultiInt64"`
	EchoMultiFloat  func(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (*EchoMultiDoubleResponse, error) //`dubbo:"echoMultiFloat"`
	EchoMultiDouble func(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (*EchoMultiDoubleResponse, error) //`dubbo:"echoMultiDouble"`
	EchoMultiString func(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string) (*EchoMultiStringResponse, error)     //`dubbo:"echoMultiString"`
}

type UserProviderImpl struct{}

func (u *UserProviderImpl) EchoBool(ctx context.Context, req bool) (bool, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt16(ctx context.Context, req int16) (int16, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt32(ctx context.Context, req int32) (int32, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt64(ctx context.Context, req int64) (int64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoFloat(ctx context.Context, req float64) (float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoDouble(ctx context.Context, req float64) (float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoString(ctx context.Context, req string) (string, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBinary(ctx context.Context, req []byte) ([]byte, error) {
	return req, nil
}

//func (u *UserProviderImpl) Echo(ctx context.Context, req *echo.EchoRequest) (r *echo.EchoResponse, err error) {
//	//TODO implement me
//	panic("implement me")
//}

func (u *UserProviderImpl) EchoBoolList(ctx context.Context, req []bool) ([]bool, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoByteList(ctx context.Context, req []int8) ([]int8, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt16List(ctx context.Context, req []int16) ([]int16, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt32List(ctx context.Context, req []int32) ([]int32, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoInt64List(ctx context.Context, req []int64) ([]int64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoFloatList(ctx context.Context, req []float64) ([]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoDoubleList(ctx context.Context, req []float64) ([]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoStringList(ctx context.Context, req []string) ([]string, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBinaryList(ctx context.Context, req [][]byte) ([][]byte, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2BoolMap(ctx context.Context, req map[bool]bool) (map[bool]bool, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2ByteMap(ctx context.Context, req map[bool]int8) (map[bool]int8, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int16Map(ctx context.Context, req map[bool]int16) (map[bool]int16, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int32Map(ctx context.Context, req map[bool]int32) (map[bool]int32, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int64Map(ctx context.Context, req map[bool]int64) (map[bool]int64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2FloatMap(ctx context.Context, req map[bool]float64) (map[bool]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2DoubleMap(ctx context.Context, req map[bool]float64) (map[bool]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2StringMap(ctx context.Context, req map[bool]string) (map[bool]string, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte) (map[bool][]byte, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool) (map[bool][]bool, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8) (map[bool][]int8, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16) (map[bool][]int16, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32) (map[bool][]int32, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64) (map[bool][]int64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2FloatListMap(ctx context.Context, req map[bool][]float64) (map[bool][]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64) (map[bool][]float64, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2StringListMap(ctx context.Context, req map[bool][]string) (map[bool][]string, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte) (map[bool][][]byte, error) {
	return req, nil
}

func (u *UserProviderImpl) EchoMultiBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool) (*EchoMultiBoolResponse, error) {
	return &EchoMultiBoolResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8) (*EchoMultiByteResponse, error) {
	return &EchoMultiByteResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16) (*EchoMultiInt16Response, error) {
	return &EchoMultiInt16Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32) (*EchoMultiInt32Response, error) {
	return &EchoMultiInt32Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64) (*EchoMultiInt64Response, error) {
	return &EchoMultiInt64Response{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiFloat(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (*EchoMultiFloatResponse, error) {
	return &EchoMultiFloatResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (*EchoMultiDoubleResponse, error) {
	return &EchoMultiDoubleResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

func (u *UserProviderImpl) EchoMultiString(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string) (*EchoMultiStringResponse, error) {
	return &EchoMultiStringResponse{
		BaseResp: baseReq,
		ListResp: listReq,
		MapResp:  mapReq,
	}, nil
}

// GetUser implements the interface
func (u *UserProviderImpl) GetUser(ctx context.Context, req int32) (*User, error) {
	var err error
	logger.Infof("req:%#v", req)
	user := &User{}
	user.ID = strconv.Itoa(int(req))
	user.Name = "laurence"
	user.Age = 22
	user.Time = time.Now()
	return user, err
}

func (u *UserProviderImpl) EchoInt(ctx context.Context, req int32) (int32, error) {
	// for exception test
	if req == 400 {
		return 0, errors.New("EchoInt failed without reason")
	}

	return req, nil
}

func (u *UserProviderImpl) EchoByte(ctx context.Context, req int8) (int8, error) {
	// for exception test
	return req, nil
}

// MethodMapper is for mapping go func name to java func name.
// Not necessary for go client -> go server
// func (s *UserProviderImpl) MethodMapper() map[string]string {
// 	return map[string]string{
// 		"GetUser": "getUser",
// 	}
//

type UserProviderImplV1 struct{}

func (u *UserProviderImplV1) EchoBool(ctx context.Context, req bool) (bool, error) {
	return !req, nil
}

func init() {
	hessian.RegisterPOJO(&User{}) // Register all transmission struct to hessian lib
	// these POJOs would override POJOs registered by Kitex generation with same JavaClassName.
	// since we do not enforce interoperability with dubbo-go, comment these statements.
	// hessian.RegisterPOJO(&EchoMultiBoolResponse{})
	// hessian.RegisterPOJO(&EchoMultiByteResponse{})
	// hessian.RegisterPOJO(&EchoMultiInt16Response{})
	// hessian.RegisterPOJO(&EchoMultiInt32Response{})
	// hessian.RegisterPOJO(&EchoMultiInt64Response{})
	// hessian.RegisterPOJO(&EchoMultiDoubleResponse{})
	// hessian.RegisterPOJO(&EchoMultiStringResponse{})
	// Register client interface to the framework
	config.SetConsumerService(UserProviderClient)
}
