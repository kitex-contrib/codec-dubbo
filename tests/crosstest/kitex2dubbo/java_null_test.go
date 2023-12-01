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

	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

func TestEchoOptionalBool_Java(t *testing.T) {
	req := false
	resp, err := cli2Java.EchoOptionalBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalInt32_Java(t *testing.T) {
	req := int32(0)
	resp, err := cli2Java.EchoOptionalInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalString_Java(t *testing.T) {
	req := ""
	resp, err := cli2Java.EchoOptionalString(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalBoolList_Java(t *testing.T) {
	var req []bool = nil
	resp, err := cli2Java.EchoOptionalBoolList(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalInt32List_Java(t *testing.T) {
	var req []int32 = nil
	resp, err := cli2Java.EchoOptionalInt32List(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalStringList_Java(t *testing.T) {
	var req []string = nil
	resp, err := cli2Java.EchoOptionalStringList(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalBool2BoolMap_Java(t *testing.T) {
	var req map[bool]bool = nil
	resp, err := cli2Java.EchoOptionalBool2BoolMap(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalBool2Int32Map_Java(t *testing.T) {
	var req map[bool]int32 = nil
	resp, err := cli2Java.EchoOptionalBool2Int32Map(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalBool2StringMap_Java(t *testing.T) {
	var req map[bool]string = nil
	resp, err := cli2Java.EchoOptionalBool2StringMap(context.Background(), nil)
	assertEcho(t, err, req, resp)
}

func TestEchoOptionalStruct_Java(t *testing.T) {
	req := &echo.EchoOptionalStructRequest{}
	resp, err := cli2Java.EchoOptionalStruct(context.Background(), req)
	var expect *echo.EchoOptionalStructResponse = nil
	assertEcho(t, err, expect, resp)
}

func TestEchoOptionalMultiBoolRequest_Java(t *testing.T) {
	req := &echo.EchoOptionalMultiBoolRequest{}
	resp, err := cli2Java.EchoOptionalMultiBoolRequest(context.Background(), req)
	assertEcho(t, err, false, resp)
}

func TestEchoOptionalMultiInt32Request_Java(t *testing.T) {
	req := &echo.EchoOptionalMultiInt32Request{}
	resp, err := cli2Java.EchoOptionalMultiInt32Request(context.Background(), req)
	assertEcho(t, err, int32(0), resp)
}

func TestEchoOptionalMultiStringRequest_Java(t *testing.T) {
	req := &echo.EchoOptionalMultiStringRequest{}
	resp, err := cli2Java.EchoOptionalMultiStringRequest(context.Background(), req)
	assertEcho(t, err, "", resp)
}

func TestEchoOptionalMultiBoolResponse_Java(t *testing.T) {
	resp, err := cli2Java.EchoOptionalMultiBoolResponse(context.Background(), false)
	expect := &echo.EchoOptionalMultiBoolResponse{}
	assertEcho(t, err, expect.GetBasicResp(), resp.GetBasicResp())
	assertEcho(t, err, expect.PackResp, resp.PackResp)
	assertEcho(t, err, expect.ListResp, resp.ListResp)
	assertEcho(t, err, expect.MapResp, resp.MapResp)
}

func TestEchoOptionalMultiInt32Response_Java(t *testing.T) {
	resp, err := cli2Java.EchoOptionalMultiInt32Response(context.Background(), 0)
	expect := &echo.EchoOptionalMultiInt32Response{}
	assertEcho(t, err, expect.GetBasicResp(), resp.GetBasicResp())
	assertEcho(t, err, expect.PackResp, resp.PackResp)
	assertEcho(t, err, expect.ListResp, resp.ListResp)
	assertEcho(t, err, expect.MapResp, resp.MapResp)
}

func TestEchoOptionalMultiStringResponse_Java(t *testing.T) {
	resp, err := cli2Java.EchoOptionalMultiStringResponse(context.Background(), "")
	expect := &echo.EchoOptionalMultiStringResponse{}
	assertEcho(t, err, expect.BaseResp, resp.BaseResp)
	assertEcho(t, err, expect.ListResp, resp.ListResp)
	assertEcho(t, err, expect.MapResp, resp.MapResp)
}
