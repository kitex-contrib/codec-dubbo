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
	"reflect"
	"testing"
)

func TestDubboHeader_RequestResponseByte(t *testing.T) {
	type fields struct {
		IsRequest bool
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name: "request",
			fields: fields{
				IsRequest: true,
			},
			want: IS_REQUEST << REQUEST_BIT_SHIFT,
		},
		{
			name: "response",
			fields: fields{
				IsRequest: false,
			},
			want: IS_RESPONSE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &DubboHeader{
				IsRequest: tt.fields.IsRequest,
			}
			if got := h.RequestResponseByte(); got != tt.want {
				t.Errorf("RequestResponseByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDubboHeader_EncodeToByteSlice(t *testing.T) {
	type fields struct {
		IsRequest       bool
		IsOneWay        bool
		IsEvent         bool
		SerializationID uint8
		Status          StatusCode
		RequestID       uint64
		DataLength      uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "Request/PingPong/Event/Hessian/OK/0x1234567887654321/0x12344321",
			fields: fields{
				IsRequest:       true,
				IsOneWay:        false,
				IsEvent:         true,
				SerializationID: SERIALIZATION_ID_HESSIAN,
				Status:          StatusOK,
				RequestID:       0x1234567887654321,
				DataLength:      0x12344321,
			},
			want: []byte{
				MAGIC_HIGH,
				MAGIC_LOW,
				(IS_REQUEST << REQUEST_BIT_SHIFT) | (IS_PINGPONG << ONEWAY_BIT_SHIFT) | (IS_EVENT << EVENT_BIT_SHIFT) | SERIALIZATION_ID_HESSIAN,
				byte(StatusOK),
				0x12, 0x34, 0x56, 0x78, 0x87, 0x65, 0x43, 0x21,
				0x12, 0x34, 0x43, 0x21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &DubboHeader{
				IsRequest:       tt.fields.IsRequest,
				IsOneWay:        tt.fields.IsOneWay,
				IsEvent:         tt.fields.IsEvent,
				SerializationID: tt.fields.SerializationID,
				Status:          tt.fields.Status,
				RequestID:       tt.fields.RequestID,
				DataLength:      tt.fields.DataLength,
			}
			if got := h.EncodeToByteSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeToByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
