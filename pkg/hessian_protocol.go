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

package dubbo

import (
	"bytes"
	"encoding/binary"
	"sync"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

// must be strict read & strict write
var (
	bpPool sync.Pool
	_      iface.BaseProtocol = (*BinaryProtocol)(nil)
)

func init() {
	bpPool.New = newBP
}

func newBP() interface{} {
	return &BinaryProtocol{}
}

// NewBinaryProtocol ...
func NewBinaryProtocol(t *bytes.Buffer) *BinaryProtocol {
	bp := bpPool.Get().(*BinaryProtocol)
	bp.trans = t
	return bp
}

// BinaryProtocol ...
type BinaryProtocol struct {
	trans *bytes.Buffer
	enc   Encoder
	dec   Decoder
}

func (p *BinaryProtocol) WriteString(s string) error {
	// TODO implement me: compact format
	b := utils.StringToSliceByte(s)
	for {
		if len(b) <= 15 {
			if err := p.trans.WriteByte(0x20 + uint8(len(b))); err != nil {
				return err
			}
			if _, err := p.trans.Write(b); err != nil {
				return err
			}
			return nil
		}

		var trunk []byte
		var tag byte
		if len(b) > 65535 { // non-final trunk
			tag = 'A'
			trunk = b[:65535]
			b = b[65535:]
		} else { // final trunk
			tag = 'B'
		}

		buf := []byte{tag, byte(len(trunk) >> 8), byte(len(trunk))}
		if _, err := p.trans.Write(buf); err != nil {
			return err
		}
		if _, err := p.trans.Write(trunk); err != nil {
			return err
		}
		if tag == 'B' {
			break
		}
	}
	return nil
}

func (p *BinaryProtocol) WriteInt32(i int32) error {
	// TODO implement me: compact encoding
	buf := make([]byte, 5)
	buf[0] = 'I'
	binary.BigEndian.PutUint32(buf[1:], uint32(i))
	_, err := p.trans.Write(buf)
	return err
}

// WriteByte ...
func (p *BinaryProtocol) WriteByte(value byte) error {
	err := p.enc.Encode(value)
	return err
}

// ReadByte ...
func (p *BinaryProtocol) ReadByte() (value byte, err error) {
	err = p.dec.Decode(p.trans)
	if err != nil {
		return 0, err
	}
	readByte, _ := p.dec.buffer.ReadByte()
	return byte(readByte), err
}
