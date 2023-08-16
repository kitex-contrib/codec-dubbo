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
	"encoding/binary"
	"fmt"
	"io"
)

/*
 * Dubbo Protocol detail:
 * https://dubbo.apache.org/zh-cn/blog/2018/10/05/dubbo-%E5%8D%8F%E8%AE%AE%E8%AF%A6%E8%A7%A3/
 */

const (
	HEADER_SIZE = 16

	MAGIC_HIGH = 0xda
	MAGIC_LOW  = 0xbb

	IS_REQUEST        = 1
	IS_RESPONSE       = 0
	REQUEST_BIT_SHIFT = 7

	IS_ONEWAY        = 0
	IS_PINGPONG      = 1
	ONEWAY_BIT_SHIFT = 6

	IS_EVENT        = 1
	EVENT_BIT_SHIFT = 5

	SERIALIZATION_ID_HESSIAN = 2
	SERIALIZATION_ID_MASK    = 0x1F

	StatusOK                  StatusCode = 20
	StatusClientTimeout       StatusCode = 30
	StatusServerTimeout       StatusCode = 31
	StatusBadRequest          StatusCode = 40
	StatusBadResponse         StatusCode = 50
	StatusServiceNotFound     StatusCode = 60
	StatusServiceError        StatusCode = 70
	StatusServerError         StatusCode = 80
	StatusClientError         StatusCode = 90
	StatusServerPoolExhausted StatusCode = 100
)

var ErrInvalidHeader = fmt.Errorf("INVALID HEADER")

type StatusCode uint8

type DubboHeader struct {
	IsRequest       bool       // 1 bit
	IsOneWay        bool       // 1 bit
	IsEvent         bool       // 1 bit
	SerializationID uint8      // 5 bits
	Status          StatusCode // 8 bits
	RequestID       uint64     // 8 bytes
	DataLength      uint32     // 4 bytes
}

func (h *DubboHeader) RequestResponseByte() byte {
	if h.IsRequest {
		return IS_REQUEST << REQUEST_BIT_SHIFT
	}
	return 0
}

func (h *DubboHeader) OnewayByte() byte {
	if !h.IsOneWay {
		return IS_PINGPONG << ONEWAY_BIT_SHIFT
	}
	return IS_ONEWAY
}

func (h *DubboHeader) EventByte() byte {
	if h.IsEvent {
		return IS_EVENT << EVENT_BIT_SHIFT
	}
	return 0
}

func (h *DubboHeader) EncodeToByteSlice() []byte {
	buf := make([]byte, HEADER_SIZE)
	buf[0] = MAGIC_HIGH
	buf[1] = MAGIC_LOW
	buf[2] = h.RequestResponseByte() | h.OnewayByte() | h.EventByte() | getSerializationID(h.SerializationID)
	buf[3] = byte(h.Status)
	binary.BigEndian.PutUint64(buf[4:12], h.RequestID)
	binary.BigEndian.PutUint32(buf[12:HEADER_SIZE], h.DataLength)
	return buf
}

func (h *DubboHeader) Encode(w io.Writer) error {
	_, err := w.Write(h.EncodeToByteSlice())
	return err
}

func (h *DubboHeader) DecodeFromByteSlice(buf []byte) error {
	if buf[0] != MAGIC_HIGH || buf[1] != MAGIC_LOW {
		return ErrInvalidHeader
	}
	h.IsRequest = isRequest(buf[2])
	h.IsOneWay = isOneWay(buf[2])
	h.IsEvent = isEvent(buf[2])
	h.SerializationID = getSerializationID(buf[2])
	h.Status = StatusCode(buf[3])
	h.RequestID = binary.BigEndian.Uint64(buf[4:12])
	h.DataLength = binary.BigEndian.Uint32(buf[12:])
	return nil
}

func (h *DubboHeader) Decode(r io.Reader) error {
	buf := make([]byte, HEADER_SIZE)
	if _, err := r.Read(buf); err != nil {
		return err
	}
	return h.DecodeFromByteSlice(buf)
}

func getSerializationID(b byte) uint8 {
	return b & SERIALIZATION_ID_MASK
}

func BitTest(b byte, shift int, expected byte) bool {
	return (b & (1 << shift) >> shift) == expected
}

func isRequest(b byte) bool {
	return BitTest(b, REQUEST_BIT_SHIFT, IS_REQUEST)
}

func isOneWay(b byte) bool {
	return BitTest(b, ONEWAY_BIT_SHIFT, IS_ONEWAY)
}

func isEvent(b byte) bool {
	return BitTest(b, EVENT_BIT_SHIFT, IS_EVENT)
}
