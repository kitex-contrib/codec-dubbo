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

package hessian2

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

func NewEncoder() iface.Encoder {
	return hessian.NewEncoder()
}

func NewDecoder(b []byte) iface.Decoder {
	return hessian.NewDecoder(b)
}

type (
	Encoder struct {
		hessian.Encoder
	}
	Decoder struct {
		hessian.Decoder
	}
)

func Register(pojos []interface{}) {
	for _, i := range pojos {
		pojo, ok := i.(hessian.POJOEnum)
		if ok {
			hessian.RegisterJavaEnum(pojo)
		} else {
			hessian.RegisterPOJO(i.(hessian.POJO))
		}
	}
}
