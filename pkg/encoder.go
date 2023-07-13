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
	"bufio"
	"bytes"
	"fmt"
	"reflect"

	hessian2 "github.com/kitex-contrib/codec-hessian2/pkg/data"
)

func NewEncoder(w bufio.Writer) *Encoder {
	return &Encoder{
		writer: w,
		buffer: bytes.NewBuffer(nil),
	}
}

type Encoder struct {
	writer bufio.Writer  // output stream
	buffer *bytes.Buffer // buffer cache
}

func (e *Encoder) Encode(obj interface{}) error {
	// nil or nil pointer
	if obj == nil || reflect.ValueOf(obj).IsNil() {
		hessian2.EncodeNull(obj)
	}

	// judge class type
	switch v := obj.(type) {
	case bool:
		panic("this type not implemented")
	case int:
		panic("this type not implemented")
	case int8:
		panic("this type not implemented")
	case int16:
		panic("this type not implemented")
	case int32:
		panic("this type not implemented")
	case int64:
		panic("this type not implemented")
	case uint:
		panic("this type not implemented")
	case uint8:
		panic("this type not implemented")
	case uint16:
		panic("this type not implemented")
	case uint32:
		panic("this type not implemented")
	case uint64:
		panic("this type not implemented")
	case float32:
		panic("this type not implemented")
	case float64:
		panic("this type not implemented")
	case string:
		panic("this type not implemented")
	case []byte:
		panic("this type not implemented")
	case []interface{}:
		panic("this type not implemented")
	case map[string]interface{}:
		panic("this type not implemented")
	default:
		return fmt.Errorf("type not supported! %s", v)
	}

	return nil
}
