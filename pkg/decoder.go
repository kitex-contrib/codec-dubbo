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

	commons "github.com/kitex-contrib/codec-hessian2/pkg/common"
)

func NewDecoder(r bufio.Reader) *Decoder {
	return &Decoder{
		reader: r,
		buffer: bytes.NewBuffer(nil),
	}
}

type Decoder struct {
	reader bufio.Reader  // input stream
	buffer *bytes.Buffer // buffer cache
}

func (d *Decoder) Decode(obj interface{}) error {
	tag, err := d.buffer.ReadByte()
	if err != nil {
		return err
	}
	// judge class type
	switch tag {
	case commons.BC_NULL:
		return nil
	case commons.BC_INT:
		panic("this type not implemented")
	case commons.BC_BINARY:
		panic("this type not implemented")
	case commons.BC_BINARY_CHUNK:
		panic("this type not implemented")
	case commons.BC_BINARY_DIRECT:
		panic("this type not implemented")
	case commons.BC_FALSE:
		panic("this type not implemented")
	case commons.BC_TRUE:
		panic("this type not implemented")
	case commons.BC_LONG:
		panic("this type not implemented")
	case commons.BC_LIST:
		panic("this type not implemented")
	case commons.BC_MAP:
		panic("this type not implemented")
	case commons.BC_DOUBLE:
		panic("this type not implemented")
	case commons.BC_LIST_FIXED:
		panic("this type not implemented")
	case commons.BC_OBJECT:
		panic("this type not implemented")
	case commons.BC_MAP_NON_TYPE:
		panic("this type not implemented")
	default:
		return fmt.Errorf("type not supported! %s", tag)
	}

	return nil
}
