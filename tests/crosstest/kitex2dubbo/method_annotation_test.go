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

// ----------kitex -> dubbo-java----------

func TestEchoBaseBool_Java(t *testing.T) {
	req := true
	resp, err := cli2Java.EchoBaseBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseByte_Java(t *testing.T) {
	req := int8(1)
	resp, err := cli2Java.EchoBaseByte(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseInt16_Java(t *testing.T) {
	req := int16(1)
	resp, err := cli2Java.EchoBaseInt16(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseInt32_Java(t *testing.T) {
	req := int32(1)
	resp, err := cli2Java.EchoBaseInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseInt64_Java(t *testing.T) {
	req := int64(1)
	resp, err := cli2Java.EchoBaseInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseDouble_Java(t *testing.T) {
	req := 1.0
	resp, err := cli2Java.EchoBaseDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseBoolList_Java(t *testing.T) {
	req := []bool{true, false}
	resp, err := cli2Java.EchoBaseBoolList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// []int8 -> byte[] does not support
// func TestEchoBaseByteList_Java(t *testing.T) {
// 	var req = []int8{1, 2}
// 	resp, err := cli2Java.EchoBaseByteList(context.Background(), req)
// 	assertEcho(t, err, req, resp)
// }

func TestEchoBaseInt16List_Java(t *testing.T) {
	req := []int16{1, 2}
	resp, err := cli2Java.EchoBaseInt16List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseInt32List_Java(t *testing.T) {
	req := []int32{1, 2}
	resp, err := cli2Java.EchoBaseInt32List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseInt64List_Java(t *testing.T) {
	req := []int64{1, 2}
	resp, err := cli2Java.EchoBaseInt64List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBaseDoubleList_Java(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli2Java.EchoBaseDoubleList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2BoolBaseMap_Java(t *testing.T) {
	req := map[bool]bool{
		true: true,
	}
	resp, err := cli2Java.EchoBool2BoolBaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2ByteBaseMap_Java(t *testing.T) {
	req := map[bool]int8{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2ByteBaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2Int16BaseMap_Java(t *testing.T) {
	req := map[bool]int16{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2Int16BaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2Int32BaseMap_Java(t *testing.T) {
	req := map[bool]int32{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2Int32BaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2Int64BaseMap_Java(t *testing.T) {
	req := map[bool]int64{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2Int64BaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2DoubleBaseMap_Java(t *testing.T) {
	req := map[bool]float64{
		true: 12.34,
	}
	resp, err := cli2Java.EchoBool2DoubleBaseMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoMultiBaseBool_Java(t *testing.T) {
	baseReq := true
	listReq := []bool{true, true}
	mapReq := map[bool]bool{
		true: true,
	}
	resp, err := cli2Java.EchoMultiBaseBool(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

// []int8 -> byte[] does not support
//func TestEchoMultiBaseByte_Java(t *testing.T) {
//	baseReq := int8(1)
//	listReq := []int8{12, 34}
//	mapReq := map[int8]int8{
//		12: 34,
//	}
//	resp, err := cli2Java.EchoMultiBaseByte(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

func TestEchoMultiBaseInt16_Java(t *testing.T) {
	baseReq := int16(1)
	listReq := []int16{12, 34}
	mapReq := map[int16]int16{
		12: 34,
	}
	resp, err := cli2Java.EchoMultiBaseInt16(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiBaseInt32_Java(t *testing.T) {
	baseReq := int32(1)
	listReq := []int32{12, 34}
	mapReq := map[int32]int32{
		12: 34,
	}
	resp, err := cli2Java.EchoMultiBaseInt32(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiBaseInt64_Java(t *testing.T) {
	baseReq := int64(1)
	listReq := []int64{12, 34}
	mapReq := map[int64]int64{
		12: 34,
	}
	resp, err := cli2Java.EchoMultiBaseInt64(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}

func TestEchoMultiBaseDouble_Java(t *testing.T) {
	baseReq := 12.34
	listReq := []float64{12.34, 56.78}
	mapReq := map[float64]float64{
		12.34: 56.78,
	}
	resp, err := cli2Java.EchoMultiBaseDouble(context.Background(), baseReq, listReq, mapReq)
	assertEcho(t, err, baseReq, resp.BaseResp)
	assertEcho(t, err, listReq, resp.ListResp)
	assertEcho(t, err, mapReq, resp.MapResp)
}
