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

package dubbo2kitex

// please see the comments in tests/dubbo-go/api/init()
// related POJO registering statements have been commented so that following cases are commented too.
// following cases with additional comments work well.

//func TestEchoMultiBool(t *testing.T) {
//	baseReq := true
//	listReq := []bool{true, true}
//	mapReq := map[bool]bool{
//		true: true,
//	}
//	resp, err := cli.EchoMultiBool(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

// hessian2.Decode does not support map[int8]int8
//func TestEchoMultiByte(t *testing.T) {
//	baseReq := int8(1)
//	listReq := []int8{12, 34}
//	mapReq := map[int8]int8{
//		12: 34,
//	}
//	resp, err := cli.EchoMultiByte(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

// hessian2.Decode does not support map[int16]int16
//func TestEchoMultiInt16(t *testing.T) {
//	baseReq := int16(1)
//	listReq := []int16{12, 34}
//	mapReq := map[int16]int16{
//		12: 34,
//	}
//	resp, err := cli.EchoMultiInt16(context.Background(), baseReq, listReq, mapReq)
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
//	resp, err := cli.EchoMultiInt32(context.Background(), baseReq, listReq, mapReq)
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
//	resp, err := cli.EchoMultiInt64(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}

//func TestEchoMultiFloat(t *testing.T) {
//	baseReq := 1.0
//	listReq := []float64{1.0, 2.0}
//	mapReq := map[float64]float64{
//		1.0: 2.0,
//	}
//	resp, err := cli.EchoMultiFloat(context.Background(), baseReq, listReq, mapReq)
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
//	resp, err := cli.EchoMultiDouble(context.Background(), baseReq, listReq, mapReq)
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
//	resp, err := cli.EchoMultiString(context.Background(), baseReq, listReq, mapReq)
//	assertEcho(t, err, baseReq, resp.BaseResp)
//	assertEcho(t, err, listReq, resp.ListResp)
//	assertEcho(t, err, mapReq, resp.MapResp)
//}
