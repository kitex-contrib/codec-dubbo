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
	"errors"
	"fmt"

	commons "github.com/kitex-contrib/codec-hessian2/pkg/common"

	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go-hessian2/java_exception"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
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
	var payload []byte
	var err error
	var status dubbo.StatusCode
	msgType := message.MessageType()
	switch msgType {
	case remote.Call, remote.Oneway:
		payload, err = m.encodeRequestPayload(ctx, message)
	case remote.Exception:
		payload, err = m.encodeExceptionPayload(ctx, message)
		// use StatusOK by default, regardless of whether it is Reply or Exception
		status = dubbo.StatusOK
	case remote.Reply:
		payload, err = m.encodeResponsePayload(ctx, message)
		status = dubbo.StatusOK
	case remote.Heartbeat:
		payload, err = m.encodeHeartbeatPayload(ctx, message)
		status = dubbo.StatusOK
	default:
		return fmt.Errorf("unsupported MessageType: %v", msgType)
	}

	if err != nil {
		return err
	}

	header := m.buildDubboHeader(message, status, len(payload))

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

func (m *Hessian2Codec) encodeRequestPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
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

func (m *Hessian2Codec) encodeResponsePayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian.NewEncoder()
	var payloadType dubbo.PayloadType
	if len(message.Tags()) != 0 {
		payloadType = dubbo.RESPONSE_VALUE_WITH_ATTACHMENTS
	} else {
		payloadType = dubbo.RESPONSE_VALUE
	}

	if err := encoder.Encode(payloadType); err != nil {
		return nil, err
	}

	// encode data
	data, ok := message.Data().(iface.Message)
	if !ok {
		return nil, fmt.Errorf("invalid data: not hessian2.MessageWriter")
	}

	if err := data.Encode(encoder); err != nil {
		return nil, err
	}

	// encode attachments if needed
	if dubbo.IsAttachmentsPayloadType(payloadType) {
		if err := encoder.Encode(message.Tags()); err != nil {
			return nil, err
		}
	}

	// java client needs this Null as the sign of end of file
	return hessian.EncNull(encoder.Buffer()), nil
}

func (m *Hessian2Codec) encodeExceptionPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian.NewEncoder()
	var payloadType dubbo.PayloadType
	if len(message.Tags()) != 0 {
		payloadType = dubbo.RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS
	} else {
		payloadType = dubbo.RESPONSE_WITH_EXCEPTION
	}

	if err := encoder.Encode(payloadType); err != nil {
		return nil, err
	}

	// encode exception
	data := message.Data()
	errRaw, ok := data.(error)
	if !ok {
		return nil, fmt.Errorf("%v exception does not implement Error", data)
	}
	if exception, ok := data.(java_exception.Throwabler); ok {
		if err := encoder.Encode(exception); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(java_exception.NewException(errRaw.Error())); err != nil {
			return nil, err
		}
	}

	if dubbo.IsAttachmentsPayloadType(payloadType) {
		if err := encoder.Encode(message.Tags()); err != nil {
			return nil, err
		}
	}

	// java client needs this Null as the sign of end of file
	return hessian.EncNull(encoder.Buffer()), nil
}

func (m *Hessian2Codec) encodeHeartbeatPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian.NewEncoder()
	// nil does not mean body is empty. after encoding, body contains 'N'
	if err := encoder.Encode(nil); err != nil {
		return nil, err
	}

	// java client needs this Null as the sign of end of file
	return hessian.EncNull(encoder.Buffer()), nil
}

func (m *Hessian2Codec) buildDubboHeader(message remote.Message, status dubbo.StatusCode, size int) *dubbo.DubboHeader {
	msgType := message.MessageType()
	return &dubbo.DubboHeader{
		IsRequest: msgType == remote.Call || msgType == remote.Oneway,
		// todo(DMwangnima): message contains heartbeat information or heartbeat flag passed in
		IsEvent:         false,
		IsOneWay:        msgType == remote.Oneway,
		SerializationID: dubbo.SERIALIZATION_ID_HESSIAN,
		Status:          status,
		RequestID:       uint64(message.RPCInfo().Invocation().SeqID()),
		DataLength:      uint32(size),
	}
}

func (m *Hessian2Codec) messageData(message remote.Message, e iface.Encoder) error {
	data, ok := message.Data().(iface.Message)
	if !ok {
		return fmt.Errorf("invalid data: not hessian2.MessageWriter")
	}
	types, err := dubbo.GetTypes(data)
	if err != nil {
		return err
	}
	if err := e.Encode(types); err != nil {
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
	if err := codec.SetOrCheckSeqID(int32(header.RequestID), message); err != nil {
		return err
	}

	// parse body part
	if header.IsRequest {
		// heartbeat package
		if header.IsEvent {
			return m.decodeEventBody(ctx, header, message, in)
		}
		return m.decodeRequestBody(ctx, header, message, in)
	}
	return m.decodeResponseBody(ctx, header, message, in)
}

func (m *Hessian2Codec) decodeEventBody(ctx context.Context, header *dubbo.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	// entire body equals to BC_NULL determines that this request is a heartbeat
	if len(body) == 1 && body[0] == commons.BC_NULL {
		message.SetMessageType(remote.Heartbeat)
	}
	// there are other events(READONLY_EVENT, WRITABLE_EVENT) in dubbo-java that we are not planing to implement

	return nil
}

func (m *Hessian2Codec) decodeRequestBody(ctx context.Context, header *dubbo.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	decoder := hessian.NewDecoder(body)
	service := new(dubbo.Service)
	if err := service.Decode(decoder); err != nil {
		return err
	}
	if serviceName := message.ServiceInfo().ServiceName; service.Path != serviceName {
		return fmt.Errorf("dubbo requested Path: %s, kitex ServiceName: %s", service.Path, serviceName)
	}

	// decode payload
	// there is no need to make use of types
	if _, err = decoder.Decode(); err != nil {
		return err
	}
	if err := codec.NewDataIfNeeded(service.Method, message); err != nil {
		return err
	}
	arg, ok := message.Data().(iface.Message)
	if !ok {
		return fmt.Errorf("invalid data: not hessian2.MessageReader")
	}
	if err := arg.Decode(decoder); err != nil {
		return err
	}
	if err := codec.SetOrCheckMethodName(service.Method, message); err != nil {
		return err
	}

	if err := processAttachments(decoder, message); err != nil {
		return err
	}

	return nil
}

func (m *Hessian2Codec) decodeResponseBody(ctx context.Context, header *dubbo.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	decoder := hessian.NewDecoder(body)
	payloadType, err := dubbo.DecodePayloadType(decoder)
	if err != nil {
		return err
	}
	switch payloadType {
	case dubbo.RESPONSE_VALUE, dubbo.RESPONSE_VALUE_WITH_ATTACHMENTS:
		msg, ok := message.Data().(iface.Message)
		if !ok {
			return fmt.Errorf("invalid data %v: not hessian2.MessageReader", msg)
		}
		if err := msg.Decode(decoder); err != nil {
			return err
		}
		if dubbo.IsAttachmentsPayloadType(payloadType) {
			if err := processAttachments(decoder, message); err != nil {
				return err
			}
		}
	// business logic exception
	case dubbo.RESPONSE_WITH_EXCEPTION, dubbo.RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS:
		exception, err := decoder.Decode()
		if err != nil {
			return err
		}
		if dubbo.IsAttachmentsPayloadType(payloadType) {
			if err := processAttachments(decoder, message); err != nil {
				return err
			}
		}
		if exceptionErr, ok := exception.(error); ok {
			return exceptionErr
		}
		return fmt.Errorf("dubbo side exception: %v", exception)
	case dubbo.RESPONSE_NULL_VALUE, dubbo.RESPONSE_NULL_VALUE_WITH_ATTACHMENTS:
		if dubbo.IsAttachmentsPayloadType(payloadType) {
			if err := processAttachments(decoder, message); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported payloadType: %v", payloadType)
	}
	return nil
}

func processAttachments(decoder iface.Decoder, message remote.Message) error {
	// decode attachments
	attachmentsRaw, err := decoder.Decode()
	if err != nil {
		return err
	}

	if attachments, ok := attachmentsRaw.(map[interface{}]interface{}); ok {
		for keyRaw, val := range attachments {
			if key, ok := keyRaw.(string); ok {
				message.Tags()[key] = val
			}
		}
		return nil
	}

	return fmt.Errorf("unsupported attachments: %v", attachmentsRaw)
}

func readBody(header *dubbo.DubboHeader, in remote.ByteBuffer) ([]byte, error) {
	length := int(header.DataLength)
	if in.ReadableLen() < length {
		return nil, errors.New("invalid dubbo package with body length being less than header specified")
	}
	return in.Next(length)
}
