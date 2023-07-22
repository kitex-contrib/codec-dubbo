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
	"context"
	"fmt"

	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/kitex-contrib/codec-hessian2/pkg/dubbo"
	"github.com/kitex-contrib/codec-hessian2/pkg/iface"
)

var _ remote.PayloadCodec = (*Hessian2Codec)(nil)

// Hessian2Codec NewHessian2Codec creates the hessian2 codec.
type Hessian2Codec struct{}

// NewHessian2Codec creates a new codec instance.
func NewHessian2Codec() *Hessian2Codec {
	return &Hessian2Codec{}
}

// Name codec name
func (m *Hessian2Codec) Name() string {
	return "hessian2"
}

// Marshal encode method
func (m *Hessian2Codec) Marshal(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	payload, err := m.buildPayload(ctx, message)
	if err != nil {
		return err
	}

	header := m.buildDubboHeader(message, len(payload))

	// write header
	if err := header.Encode(out); err != nil {
		return err
	}

	// write payload
	if _, err := out.WriteBinary(payload); err != nil {
		return err
	}
	return nil
}

func (m *Hessian2Codec) buildPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian.NewEncoder()

	if err = m.messageBegin(ctx, message, encoder); err != nil {
		return nil, err
	}

	if err = m.messageData(message, encoder); err != nil {
		return nil, err
	}

	if err = m.messageEnd(ctx, message, encoder); err != nil {
		return nil, err
	}

	return encoder.Buffer(), nil
}

func (m *Hessian2Codec) buildDubboHeader(message remote.Message, size int) *dubbo.DubboHeader {
	msgType := message.MessageType()
	return &dubbo.DubboHeader{
		IsRequest:       msgType == remote.Call || msgType == remote.Oneway,
		IsEvent:         true,
		IsOneWay:        msgType == remote.Oneway,
		SerializationID: dubbo.SERIALIZATION_ID_HESSIAN,
		RequestID:       uint64(message.RPCInfo().Invocation().SeqID()),
		DataLength:      uint32(size),
	}
}

func (m *Hessian2Codec) messageData(message remote.Message, e iface.Encoder) error {
	data, ok := message.Data().(iface.Message)
	if !ok {
		return fmt.Errorf("invalid data: not hessian2.MessageWriter")
	}
	//TODO: e.Encode(data.GetTypes()), return err if err != nil
	return data.Encode(e)
}

func (m *Hessian2Codec) messageBegin(ctx context.Context, message remote.Message, e iface.Encoder) error {
	service := &dubbo.Service{
		ProtocolVersion: dubbo.DEFAULT_DUBBO_PROTOCOL_VERSION,
		Path:            "<PATH:TODO>",
		Version:         "<VERSION:TODO>",
		Method:          message.RPCInfo().Invocation().MethodName(),
	}
	if err := e.Encode(service.ProtocolVersion); err != nil {
		return err
	}
	if err := e.Encode(service.Path); err != nil {
		return err
	}
	if err := e.Encode(service.Version); err != nil {
		return err
	}
	if err := e.Encode(service.Method); err != nil {
		return err
	}
	return nil
}

func (m *Hessian2Codec) messageEnd(ctx context.Context, message remote.Message, e iface.Encoder) error {
	// TODO: WIP
	attachment := dubbo.NewAttachment(
		"<PATH:TODO>",
		"<GROUP:TODO>",
		"<IFACE:TODO>",
		"<VERSION:TODO>",
		0, // TODO
	)
	return e.Encode(attachment)
}

// Unmarshal decode method
func (m *Hessian2Codec) Unmarshal(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	// TODO: WIP
	panic("not implemented")
}
