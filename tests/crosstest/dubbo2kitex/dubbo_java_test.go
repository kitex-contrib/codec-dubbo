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

import (
	"testing"

	"github.com/kitex-contrib/codec-dubbo/tests/util"
)

func TestDubboJava(t *testing.T) {
	util.RunAndTestDubboJavaClient(t, "../../dubbo-java", "org.apache.dubbo.tests.client.Application",
		nil,
		[]string{
			// comment lines mean dubbo-java can not support
			"EchoBool",
			"EchoByte",
			"EchoInt16",
			"EchoInt32",
			"EchoInt64",
			"EchoFloat",
			"EchoDouble",
			"EchoString",
			"EchoBinary",
			"EchoBoolList",
			//"EchoByteList",
			//"EchoInt16List",
			"EchoInt32List",
			"EchoInt64List",
			//"EchoFloatList",
			"EchoDoubleList",
			"EchoStringList",
			// hessian2 can not support encoding [][]byte
			// dubbo-java can not support
			//"EchoBinaryList",
			"EchoBool2BoolMap",
			//"EchoBool2ByteMap",
			//"EchoBool2Int16Map",
			"EchoBool2Int32Map",
			"EchoBool2Int64Map",
			//"EchoBool2FloatMap",
			"EchoBool2DoubleMap",
			"EchoBool2StringMap",
			//"EchoBool2BinaryMap",
			"EchoMultiBool",
			"EchoMultiByte",
			"EchoMultiInt16",
			"EchoMultiInt32",
			"EchoMultiInt64",
			"EchoMultiFloat",
			"EchoMultiDouble",
			"EchoMultiString",
			"EchoMethodA",
			"EchoMethodB",
			"EchoMethodC",
			"EchoMethodD",
		},
	)
}
