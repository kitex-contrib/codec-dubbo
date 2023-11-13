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
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/kitex-contrib/codec-dubbo/pkg/dubbo_spec"
	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	"github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

var _ remote.Codec = (*DubboCodec)(nil)

// DubboCodec NewDubboCodec creates the dubbo codec.
type DubboCodec struct {
	opt         *Options
	methodCache hessian2.MethodCache
}

// NewDubboCodec creates a new codec instance.
func NewDubboCodec(opts ...Option) *DubboCodec {
	o := newOptions(opts)
	return &DubboCodec{opt: o}
}

// Name codec name
func (m *DubboCodec) Name() string {
	return "dubbo"
}

// Marshal encode method
func (m *DubboCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	var payload []byte
	var err error
	var status dubbo_spec.StatusCode
	// indicate whether this pkg is event
	var eventFlag bool
	msgType := message.MessageType()
	switch msgType {
	case remote.Call, remote.Oneway:
		payload, err = m.encodeRequestPayload(ctx, message)
	case remote.Exception:
		payload, err = m.encodeExceptionPayload(ctx, message)
		// todo(DMwangnima): refer to exception processing logic of dubbo-java, use status to determine if this exception
		// is in outside layer.(eg. non-exist InterfaceName)
		// for now, use StatusOK by default, regardless of whether it is in outside layer.
		status = dubbo_spec.StatusOK
	case remote.Reply:
		payload, err = m.encodeResponsePayload(ctx, message)
		status = dubbo_spec.StatusOK
	case remote.Heartbeat:
		payload, err = m.encodeHeartbeatPayload(ctx, message)
		status = dubbo_spec.StatusOK
		eventFlag = true
	default:
		return fmt.Errorf("unsupported MessageType: %v", msgType)
	}

	if err != nil {
		return err
	}

	header := m.buildDubboHeader(message, status, len(payload), eventFlag)

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

func (m *DubboCodec) encodeRequestPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian2.NewEncoder()

	service := &dubbo_spec.Service{
		ProtocolVersion: dubbo_spec.DEFAULT_DUBBO_PROTOCOL_VERSION,
		Path:            m.opt.JavaClassName,
		// todo: kitex mapping
		Version: "",
		Method:  message.RPCInfo().Invocation().MethodName(),
		Timeout: message.RPCInfo().Config().RPCTimeout(),
		// todo: kitex mapping
		Group: "",
	}
	methodAnno := m.getMethodAnnotation(message)

	// if a method name annotation exists, update the method name to the annotation value.
	if methodName, exists := methodAnno.GetMethodName(); exists {
		service.Method = methodName
	}

	if err = m.messageServiceInfo(ctx, service, encoder); err != nil {
		return nil, err
	}

	if err = m.messageData(message, methodAnno, encoder); err != nil {
		return nil, err
	}

	if err = m.messageAttachment(ctx, service, encoder); err != nil {
		return nil, err
	}

	return encoder.Buffer(), nil
}

func (m *DubboCodec) encodeResponsePayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian2.NewEncoder()
	var payloadType dubbo_spec.PayloadType
	if len(message.Tags()) != 0 {
		payloadType = dubbo_spec.RESPONSE_VALUE_WITH_ATTACHMENTS
	} else {
		payloadType = dubbo_spec.RESPONSE_VALUE
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
	if dubbo_spec.IsAttachmentsPayloadType(payloadType) {
		if err := encoder.Encode(message.Tags()); err != nil {
			return nil, err
		}
	}

	return encoder.Buffer(), nil
}

func (m *DubboCodec) encodeExceptionPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian2.NewEncoder()
	var payloadType dubbo_spec.PayloadType
	if len(message.Tags()) != 0 {
		payloadType = dubbo_spec.RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS
	} else {
		payloadType = dubbo_spec.RESPONSE_WITH_EXCEPTION
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
	if exception, ok := data.(hessian2.Throwabler); ok {
		if err := encoder.Encode(exception); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(hessian2.NewException(errRaw.Error())); err != nil {
			return nil, err
		}
	}

	if dubbo_spec.IsAttachmentsPayloadType(payloadType) {
		if err := encoder.Encode(message.Tags()); err != nil {
			return nil, err
		}
	}

	return encoder.Buffer(), nil
}

// Event Flag set in dubbo header and 'N' body determines that this pkg is heartbeat.
// For dubbo-go, it does not decode the body of the pkg when Event Flag is set in dubbo header.
// For dubbo-java, it reads the body of the pkg and use this statement to judge when Event Flag is set in dubbo header.
// Arrays.equals(payload, getNullBytesOf(getSerializationById(proto)))
// For hessian2, NullByte is 'N'.
// As a result, we need to encode nil in heartbeat response body for both dubbo-go side and dubbo-java side.
func (m *DubboCodec) encodeHeartbeatPayload(ctx context.Context, message remote.Message) (buf []byte, err error) {
	encoder := hessian2.NewEncoder()

	if err := encoder.Encode(nil); err != nil {
		return nil, err
	}

	return encoder.Buffer(), nil
}

func (m *DubboCodec) buildDubboHeader(message remote.Message, status dubbo_spec.StatusCode, size int, eventFlag bool) *dubbo_spec.DubboHeader {
	msgType := message.MessageType()
	return &dubbo_spec.DubboHeader{
		IsRequest:       msgType == remote.Call || msgType == remote.Oneway,
		IsEvent:         eventFlag,
		IsOneWay:        msgType == remote.Oneway,
		SerializationID: dubbo_spec.SERIALIZATION_ID_HESSIAN,
		Status:          status,
		RequestID:       uint64(message.RPCInfo().Invocation().SeqID()),
		DataLength:      uint32(size),
	}
}

func (m *DubboCodec) messageData(message remote.Message, methodAnno *hessian2.MethodAnnotation, e iface.Encoder) error {
	data, ok := message.Data().(iface.Message)
	if !ok {
		return fmt.Errorf("invalid data: not hessian2.MessageWriter")
	}

	types, err := m.methodCache.GetTypes(data, methodAnno)
	if err != nil {
		return err
	}
	if err := e.Encode(types); err != nil {
		return err
	}
	return data.Encode(e)
}

func (m *DubboCodec) messageServiceInfo(ctx context.Context, service *dubbo_spec.Service, e iface.Encoder) error {
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

func (m *DubboCodec) messageAttachment(ctx context.Context, service *dubbo_spec.Service, e iface.Encoder) error {
	attachment := dubbo_spec.NewAttachment(
		service.Path,
		service.Group,
		service.Path,
		service.Version,
		service.Timeout,
	)
	return e.Encode(attachment)
}

func (m *DubboCodec) getMethodAnnotation(message remote.Message) *hessian2.MethodAnnotation {
	methodKey := message.ServiceInfo().ServiceName + "." + message.RPCInfo().To().Method()
	if m.opt.MethodAnnotations != nil {
		if t, ok := m.opt.MethodAnnotations[methodKey]; ok {
			return t
		}
	}
	return nil
}

// Unmarshal decode method
func (m *DubboCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	// parse header part
	header := new(dubbo_spec.DubboHeader)
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

	if header.Status != dubbo_spec.StatusOK {
		return m.decodeExceptionBody(ctx, header, message, in)
	}
	return m.decodeResponseBody(ctx, header, message, in)
}

func (m *DubboCodec) decodeEventBody(ctx context.Context, header *dubbo_spec.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	// entire body equals to BC_NULL determines that this request is a heartbeat
	if len(body) == 1 && body[0] == hessian2.NULL {
		message.SetMessageType(remote.Heartbeat)
	}
	// todo(DMwangnima): there are other events(READONLY_EVENT, WRITABLE_EVENT) in dubbo-java that we are planning to implement currently

	return nil
}

func (m *DubboCodec) decodeRequestBody(ctx context.Context, header *dubbo_spec.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	decoder := hessian2.NewDecoder(body)
	service := new(dubbo_spec.Service)
	if err := service.Decode(decoder); err != nil {
		return err
	}

	if name := m.opt.JavaClassName; service.Path != name {
		return fmt.Errorf("dubbo requested Path: %s, kitex service specified JavaClassName: %s", service.Path, name)
	}

	// decode payload
	if types, err := decoder.Decode(); err != nil {
		return err
	} else if method, exists := m.opt.MethodNames[service.Method+types.(string)]; exists {
		service.Method = method
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

// decodeExceptionBody is responsible for processing exception in the outer layer which means business logic
// in the remoting service has not been invoked. (eg. wrong request with non-exist InterfaceName)
func (m *DubboCodec) decodeExceptionBody(ctx context.Context, header *dubbo_spec.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	decoder := hessian2.NewDecoder(body)
	exception, err := decoder.Decode()
	if err != nil {
		return err
	}
	exceptionStr, ok := exception.(string)
	if !ok {
		return fmt.Errorf("exception %v is not of string", exception)
	}
	return fmt.Errorf("dubbo side exception: %s", exceptionStr)
}

func (m *DubboCodec) decodeResponseBody(ctx context.Context, header *dubbo_spec.DubboHeader, message remote.Message, in remote.ByteBuffer) error {
	body, err := readBody(header, in)
	if err != nil {
		return err
	}

	decoder := hessian2.NewDecoder(body)
	payloadType, err := dubbo_spec.DecodePayloadType(decoder)
	if err != nil {
		return err
	}
	switch payloadType {
	case dubbo_spec.RESPONSE_VALUE, dubbo_spec.RESPONSE_VALUE_WITH_ATTACHMENTS:
		msg, ok := message.Data().(iface.Message)
		if !ok {
			return fmt.Errorf("invalid data %v: not hessian2.MessageReader", msg)
		}
		if err := msg.Decode(decoder); err != nil {
			return err
		}
		if dubbo_spec.IsAttachmentsPayloadType(payloadType) {
			if err := processAttachments(decoder, message); err != nil {
				return err
			}
		}
	// business logic exception
	case dubbo_spec.RESPONSE_WITH_EXCEPTION, dubbo_spec.RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS:
		exception, err := decoder.Decode()
		if err != nil {
			return err
		}
		if dubbo_spec.IsAttachmentsPayloadType(payloadType) {
			if err := processAttachments(decoder, message); err != nil {
				return err
			}
		}
		if exceptionErr, ok := exception.(error); ok {
			return exceptionErr
		}
		return fmt.Errorf("dubbo side exception: %v", exception)
	case dubbo_spec.RESPONSE_NULL_VALUE, dubbo_spec.RESPONSE_NULL_VALUE_WITH_ATTACHMENTS:
		if dubbo_spec.IsAttachmentsPayloadType(payloadType) {
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

func readBody(header *dubbo_spec.DubboHeader, in remote.ByteBuffer) ([]byte, error) {
	length := int(header.DataLength)
	return in.Next(length)
}
