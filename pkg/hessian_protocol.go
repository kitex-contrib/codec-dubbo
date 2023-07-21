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
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

// must be strict read & strict write
var (
	bpPool sync.Pool
	_      BaseProtocol = (*BinaryProtocol)(nil)
)

func init() {
	bpPool.New = newBP
}

func newBP() interface{} {
	return &BinaryProtocol{}
}

// NewBinaryProtocol ...
func NewBinaryProtocol(t remote.ByteBuffer) *BinaryProtocol {
	bp := bpPool.Get().(*BinaryProtocol)
	bp.trans = t
	return bp
}

// BinaryProtocol ...
type BinaryProtocol struct {
	trans remote.ByteBuffer
	enc   Encoder
	dec   Decoder
}

// WriteByte ...
func (p *BinaryProtocol) WriteByte(value int8) error {
	err := p.enc.Encode(value)
	return err
}

// ReadByte ...
func (p *BinaryProtocol) ReadByte() (value int8, err error) {
	err = p.dec.Decode(p.trans)
	if err != nil {
		return 0, err
	}
	readByte, _ := p.dec.buffer.ReadByte()
	return int8(readByte), err
}
