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

func TestEchoBoolList(t *testing.T) {
	req := []bool{true, false}
	resp, err := cli2Go.EchoBoolList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-go does not support
//func TestEchoByteList(t *testing.T) {
//	var req = []int8{1, 2}
//	resp, err := cli2Go.EchoByteList(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// dubbo-go does not support
//func TestEchoInt16List(t *testing.T) {
//	var req = []int16{1, 2}
//	resp, err := cli2Go.EchoInt16List(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

func TestEchoInt32List(t *testing.T) {
	req := []int32{1, 2}
	resp, err := cli2Go.EchoInt32List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64List(t *testing.T) {
	req := []int64{1, 2}
	resp, err := cli2Go.EchoInt64List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoFloatList(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli2Go.EchoFloatList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDoubleList(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli2Go.EchoDoubleList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoStringList(t *testing.T) {
	req := []string{"1", "2"}
	resp, err := cli2Go.EchoStringList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-go hessian2 does not support [][]byte, please refer to github.com/apache/dubbo-go-hessian2/list.go
//func TestEchoBinaryList(t *testing.T) {
//	var req = [][]byte{{'1'}, {'2'}}
//	resp, err := cli2Go.EchoBinaryList(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// We have supported other map types (refer to /pkg/hessian2/response_test),
// but ReflectResponse in dubbo-go side could not support. As a result dubbo-go could not parse
// map types request correctly.
// We would finish this part of tests in kitex -> dubbo-java.

//func TestEchoBool2BoolMap(t *testing.T) {
//	req := map[bool]bool{
//		true: true,
//	}
//	resp, err := cli2Go.EchoBool2BoolMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2ByteMap(t *testing.T) {
//	req := map[bool]int8{
//		true: 1,
//	}
//	resp, err := cli2Go.EchoBool2ByteMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int16Map(t *testing.T) {
//	req := map[bool]int16{
//		true: 1,
//	}
//	resp, err := cli2Go.EchoBool2Int16Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int32Map(t *testing.T) {
//	req := map[bool]int32{
//		true: 1,
//	}
//	resp, err := cli2Go.EchoBool2Int32Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2Int64Map(t *testing.T) {
//	req := map[bool]int64{
//		true: 1,
//	}
//	resp, err := cli2Go.EchoBool2Int64Map(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2FloatMap(t *testing.T) {
//	req := map[bool]float64{
//		true: 12.34,
//	}
//	resp, err := cli2Go.EchoBool2FloatMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2DoubleMap(t *testing.T) {
//	req := map[bool]float64{
//		true: 12.34,
//	}
//	resp, err := cli2Go.EchoBool2DoubleMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2StringMap(t *testing.T) {
//	req := map[bool]string{
//		true: "1",
//	}
//	resp, err := cli2Go.EchoBool2StringMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

//func TestEchoBool2BinaryMap(t *testing.T) {
//	req := map[bool][]byte{
//		true: {'1', '2'},
//	}
//	resp, err := cli2Go.EchoBool2BinaryMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// ----------kitex -> dubbo-java----------

// dubbo-go-hessian2 Encode can not produce correct types for List and Map.
// todo: improve hessian2 Encode

func TestEchoBoolList_Java(t *testing.T) {
	req := []bool{true, false}
	resp, err := cli2Java.EchoBoolList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-java -> dubbo-java does not support
// func TestEchoByteList_Java(t *testing.T) {
// 	var req = []int8{1, 2}
// 	resp, err := cli2Java.EchoByteList(context.Background(), req)
// 	assertEcho(t, err, req, resp)
// }

// dubbo-java -> dubbo-java does not support
// func TestEchoInt16List_Java(t *testing.T) {
// 	var req = []int16{1, 2}
// 	resp, err := cli2Java.EchoInt16List(context.Background(), req)
// 	assertEcho(t, err, req, resp)
// }

func TestEchoInt32List_Java(t *testing.T) {
	req := []int32{1, 2}
	resp, err := cli2Java.EchoInt32List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64List_Java(t *testing.T) {
	req := []int64{1, 2}
	resp, err := cli2Java.EchoInt64List(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoFloatList_Java(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli2Java.EchoFloatList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDoubleList_Java(t *testing.T) {
	req := []float64{12.3456, 78.9012}
	resp, err := cli2Java.EchoDoubleList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoStringList_Java(t *testing.T) {
	req := []string{"1", "2"}
	resp, err := cli2Java.EchoStringList(context.Background(), req)
	assertEcho(t, err, req, resp)
}

//func TestEchoBinaryList_Java(t *testing.T) {
//	var req = [][]byte{{'1'}, {'2'}}
//	resp, err := cli2Java.EchoBinaryList(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

func TestEchoBool2BoolMap_Java(t *testing.T) {
	req := map[bool]bool{
		true: true,
	}
	resp, err := cli2Java.EchoBool2BoolMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-java -> dubbo-java does not support
// func TestEchoBool2ByteMap_Java(t *testing.T) {
// 	req := map[bool]int8{
// 		true: 1,
// 	}
// 	resp, err := cli2Java.EchoBool2ByteMap(context.Background(), req)
// 	assertEcho(t, err, req, resp)
// }

// dubbo-java -> dubbo-java does not support
// func TestEchoBool2Int16Map_Java(t *testing.T) {
// 	req := map[bool]int16{
// 		true: 1,
// 	}
// 	resp, err := cli2Java.EchoBool2Int16Map(context.Background(), req)
// 	assertEcho(t, err, req, resp)
// }

func TestEchoBool2Int32Map_Java(t *testing.T) {
	req := map[bool]int32{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2Int32Map(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2Int64Map_Java(t *testing.T) {
	req := map[bool]int64{
		true: 1,
	}
	resp, err := cli2Java.EchoBool2Int64Map(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2FloatMap_Java(t *testing.T) {
	req := map[bool]float64{
		true: 12.34,
	}
	resp, err := cli2Java.EchoBool2FloatMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2DoubleMap_Java(t *testing.T) {
	req := map[bool]float64{
		true: 12.34,
	}
	resp, err := cli2Java.EchoBool2DoubleMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBool2StringMap_Java(t *testing.T) {
	req := map[bool]string{
		true: "1",
	}
	resp, err := cli2Java.EchoBool2StringMap(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-java -> dubbo-java does not support
//func TestEchoBool2BinaryMap_Java(t *testing.T) {
//	req := map[bool][]byte{
//		true: {'1', '2'},
//	}
//	resp, err := cli2Java.EchoBool2BinaryMap(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}
