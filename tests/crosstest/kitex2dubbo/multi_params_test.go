/*
 * Copyright 2023 CloudWeGo Authors
 *
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

package kitex2dubbo

import (
	"context"
	"testing"
)

// We have supported other map types (refer to /pkg/hessian2/response_test),
// but ReflectResponse in dubbo-go side could not support. As a result dubbo-go could not parse
// map types request correctly.

// dubbo-go does not support Byte, Int16,  ByteList and Int16List

//func TestEchoMultiBool(t *testing.T) {
//	baseReq := true
//	listReq := []bool{true, true}
//	mapReq := map[bool]bool{
//		true: true,
//	}
//	resp, err := cli2Go.EchoMultiBool(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiByte(t *testing.T) {
//	baseReq := int8(1)
//	listReq := []int8{12, 34}
//	mapReq := map[int8]int8{
//		12: 34,
//	}
//	resp, err := cli2Go.EchoMultiByte(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiInt16(t *testing.T) {
//	baseReq := int16(1)
//	listReq := []int16{12, 34}
//	mapReq := map[int16]int16{
//		12: 34,
//	}
//	resp, err := cli2Go.EchoMultiInt16(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiInt32(t *testing.T) {
//	baseReq := int32(1)
//	listReq := []int32{12, 34}
//	mapReq := map[int32]int32{
//		12: 34,
//	}
//	resp, err := cli2Go.EchoMultiInt32(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiInt64(t *testing.T) {
//	baseReq := int64(1)
//	listReq := []int64{12, 34}
//	mapReq := map[int64]int64{
//		12: 34,
//	}
//	resp, err := cli2Go.EchoMultiInt64(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiDouble(t *testing.T) {
//	baseReq := 12.34
//	listReq := []float64{12.34, 56.78}
//	mapReq := map[float64]float64{
//		12.34: 56.78,
//	}
//	resp, err := cli2Go.EchoMultiDouble(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiString(t *testing.T) {
//	baseReq := "1"
//	listReq := []string{"12", "34"}
//	mapReq := map[string]string{
//		"12": "34",
//	}
//	resp, err := cli2Go.EchoMultiString(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

// ----------kitex -> dubbo-java----------

func TestEchoMultiBool_Java(t *testing.T) {
	baseReq := true
	listReq := []bool{true, true}
	mapReq := map[bool]bool{
		true: true,
	}
	resp, err := cli2Java.EchoMultiBool(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

// hessian2.Decode does not support but dubbo-java supports
//func TestEchoMultiByte_Java(t *testing.T) {
//	baseReq := int8(1)
//	listReq := []int8{12, 34}
//	mapReq := map[int8]int8{
//		12: 34,
//	}
//	resp, err := cli2Java.EchoMultiByte(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

// hessian2.Decode does not support but dubbo-java supports
//func TestEchoMultiInt16_Java(t *testing.T) {
//	baseReq := int16(1)
//	listReq := []int16{12, 34}
//	mapReq := map[int16]int16{
//		12: 34,
//	}
//	resp, err := cli2Java.EchoMultiInt16(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

func TestEchoMultiInt32_Java(t *testing.T) {
	baseReq := int32(1)
	listReq := []int32{12, 34}
	mapReq := map[int32]int32{
		12: 34,
	}
	resp, err := cli2Java.EchoMultiInt32(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiInt64_Java(t *testing.T) {
	baseReq := int64(1)
	listReq := []int64{12, 34}
	mapReq := map[int64]int64{
		12: 34,
	}
	resp, err := cli2Java.EchoMultiInt64(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiDouble_Java(t *testing.T) {
	baseReq := 12.34
	listReq := []float64{12.34, 56.78}
	mapReq := map[float64]float64{
		12.34: 56.78,
	}
	resp, err := cli2Java.EchoMultiDouble(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiString_Java(t *testing.T) {
	baseReq := "1"
	listReq := []string{"12", "34"}
	mapReq := map[string]string{
		"12": "34",
	}
	resp, err := cli2Java.EchoMultiString(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}
