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

var _ remote.Codec = (*Hessian2Codec)(nil)

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
func (m *Hessian2Codec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
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

	service := &dubbo.Service{
		ProtocolVersion: dubbo.DEFAULT_DUBBO_PROTOCOL_VERSION,
		// todo: should be message.RPCInfo().Invocation.ServiceName
		Path: message.RPCInfo().To().ServiceName(),
		// todo: kitex mapping
		Version: "",
		Method:  message.RPCInfo().Invocation().MethodName(),
		Timeout: message.RPCInfo().Config().RPCTimeout(),
		// todo: kitex mapping
		Group: "",
	}
	if err = m.messageServiceInfo(ctx, service, encoder); err != nil {
		return nil, err
	}

	if err = m.messageData(message, encoder); err != nil {
		return nil, err
	}

	if err = m.messageAttachment(ctx, service, encoder); err != nil {
		return nil, err
	}

	return encoder.Buffer(), nil
}

func (m *Hessian2Codec) buildDubboHeader(message remote.Message, size int) *dubbo.DubboHeader {
	msgType := message.MessageType()
	return &dubbo.DubboHeader{
		IsRequest:       msgType == remote.Call || msgType == remote.Oneway,
		IsEvent:         false,
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
	if err := e.Encode(data.GetTypes()); err != nil {
		return err
	}
	return data.Encode(e)
}

func (m *Hessian2Codec) messageServiceInfo(ctx context.Context, service *dubbo.Service, e iface.Encoder) error {
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

func (m *Hessian2Codec) messageAttachment(ctx context.Context, service *dubbo.Service, e iface.Encoder) error {
	attachment := dubbo.NewAttachment(
		service.Path,
		service.Group,
		service.Path,
		service.Version,
		service.Timeout,
	)
	return e.Encode(attachment)
}

// Unmarshal decode method
func (m *Hessian2Codec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	// parse header part
	header := new(dubbo.DubboHeader)
	if err := header.Decode(in); err != nil {
		return err
	}

	// parse body part
	body, err := in.Peek(int(header.DataLength))
	if err != nil {
		return err
	}
	decoder := hessian.NewDecoder(body)
	payloadType, err := decoder.Decode()
	if err != nil {
		return err
	}
	switch payloadType {
	// todo: processing other payload types with attachments
	case dubbo.RESPONSE_VALUE:
		msg, ok := message.Data().(iface.Message)
		if !ok {
			return fmt.Errorf("invalid data %v: not hessian2.MessageReader", msg)
		}
		if err := msg.Decode(decoder); err != nil {
			return err
		}
	case dubbo.RESPONSE_WITH_EXCEPTION:
		exception, err := decoder.Decode()
		if err != nil {
			return err
		}
		if exceptionErr, ok := exception.(error); ok {
			return exceptionErr
		}
		return fmt.Errorf("dubbo side exception: %v", exception)
	default:
		return fmt.Errorf("unsupported payloadType: %v", payloadType)
	}
	return nil
}
