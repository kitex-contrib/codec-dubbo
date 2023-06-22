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
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cloudwego/kitex/pkg/utils"
	"io"
	"math"
	"reflect"
	"strings"

	commons "github.com/kitex-contrib/codec-hessian2/pkg/common"
)

func NewHessian2Output(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
		buffer: bytes.NewBuffer(nil),
	}
}

// Class is java object info
type Class struct {
	Name   string   // class name
	Fields []*Field // filed info
}

// Field is java object field info
type Field struct {
	Name    string      // filed name
	Type    *Class      // filed type
	Default interface{} // default value
}

type Encoder struct {
	writer io.Writer     // output stream
	buffer *bytes.Buffer // buffer cache

	writeType      bool          // write type info
	typeRefWritten bool          // is written type info
	typeRefs       []*Class      // class references
	objRefs        []interface{} // object references
}

func (e *Encoder) writeObjectRef(i int) error {
	if err := e.WriteByte('R'); err != nil {
		return err
	}
	if err := binary.Write(e.buffer, binary.BigEndian, i); err != nil {
		return err
	}
	return nil
}
func (e *Encoder) WriteObject(obj interface{}) error {
	// nil or nil pointer
	if obj == nil || reflect.ValueOf(obj).IsNil() {
		return e.WriteNull()
	}

	clazz, err := getClass(obj)
	if err != nil {
		return err
	}

	// write class info
	if e.writeType {
		if err := e.writeClass(clazz); err != nil {
			return err
		}
	}

	// write object reference info
	if ref, ok := e.findObjRef(obj); ok {
		return e.writeObjectRef(ref)
	}

	// write objects to references
	e.objRefs = append(e.objRefs, obj)

	// judge class type
	switch v := obj.(type) {
	case bool:
		return e.WriteBool(v)
	case int:
		return e.WriteInt(int64(v))
	case int8:
		return e.WriteInt(int64(v))
	case int16:
		return e.WriteInt(int64(v))
	case int32:
		return e.WriteInt(int64(v))
	case int64:
		return e.WriteInt(v)
	case uint:
		return e.WriteInt(int64(v))
	case uint8:
		return e.WriteInt(int64(v))
	case uint16:
		return e.WriteInt(int64(v))
	case uint32:
		return e.WriteInt(int64(v))
	case uint64:
		return e.WriteInt(int64(v))
	case float32:
		return e.writeDouble(float64(v))
	case float64:
		return e.writeDouble(v)
	case string:
		return e.WriteString(v)
	}

	// encode complex types
	switch v := obj.(type) {
	case []byte:
		return e.writeBytes(v)
	case []interface{}:
		return e.writeList(v)
	case map[string]interface{}:
		return e.writeMap(v)
	}

	// encode class types
	return e.writeClass(clazz)
}

// get object type info
func getClass(obj interface{}) (*Class, error) {
	if obj == nil {
		return nil, fmt.Errorf("getClass: nil object")
	}

	// get an object type
	t := reflect.TypeOf(obj)

	name := getClassName(t)
	if name == "" {
		return nil, fmt.Errorf("getClass: invalid object type: %v", t)
	}

	// if struct type,contract object
	if t.Kind() == reflect.Struct {
		fields, err := getClassFields(t)
		if err != nil {
			return nil, err
		}
		return &Class{Name: name, Fields: fields}, nil
	}

	return &Class{Name: name}, nil
}

// get class name
func getClassName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Struct:
		if t.NumField() == 0 {
			return ""
		}
		// struct name to lower
		return strings.ToLower(t.Name())
	case reflect.Slice:
		// is byte[]
		if t.Elem().Kind() == reflect.Uint8 {
			return "byte[]"
		}
		return "list"
	case reflect.Map:
		return "map"
	case reflect.Ptr:
		// point name
		return getClassName(t.Elem())
	default:
		return strings.ToLower(t.Name())
	}
}

// encode class info to commons protocol
func (e *Encoder) writeClass(clazz *Class) error {
	// get type reference
	if ref, ok := e.findTypeRef(clazz); ok {
		return e.writeTypeRef(ref)
	}

	// add types to references
	e.typeRefs = append(e.typeRefs, clazz)

	// write class tag
	if err := e.WriteByte(commons.BC_OBJECT_DEF); err != nil {
		return err
	}

	// write class name
	if err := e.WriteString(clazz.Name); err != nil {
		return err
	}

	// write class info
	if err := e.writeFields(clazz.Fields); err != nil {
		return err
	}

	// is written type info
	e.typeRefWritten = true

	return nil
}

// get type reference
func (e *Encoder) findTypeRef(clazz *Class) (int64, bool) {
	for i, ref := range e.typeRefs {
		if reflect.DeepEqual(clazz, ref) {
			return int64(i), true
		}
	}
	return -1, false
}

// write type reference
func (e *Encoder) writeTypeRef(ref int64) error {
	if err := e.WriteByte(byte('T')); err != nil {
		return err
	}
	if err := e.WriteInt(ref); err != nil {
		return err
	}
	return nil
}

// write fields info
func (e *Encoder) writeFields(fields []*Field) error {
	for _, field := range fields {
		if err := e.writeField(field); err != nil {
			return err
		}
	}
	return nil
}

// write field info
func (e *Encoder) writeField(field *Field) error {
	if err := e.WriteString(field.Name); err != nil {
		return err
	}
	if err := e.writeTypeRefOrClass(field.Type); err != nil {
		return err
	}
	if field.Default != nil {
		if err := e.WriteObject(field.Default); err != nil {
			return err
		}
	}
	return nil
}

// write type reference or class info
func (e *Encoder) writeTypeRefOrClass(clazz *Class) error {
	if ref, ok := e.findTypeRef(clazz); ok {
		if err := e.writeTypeRef(ref); err != nil {
			return err
		}
	} else {
		if err := e.writeClass(clazz); err != nil {
			return err
		}
	}
	return nil
}

// get struct fields
func getClassFields(t reflect.Type) ([]*Field, error) {
	var fields []*Field

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := strings.ToLower(field.Name)
		typ, err := getClass(field.Type)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &Field{Name: name, Type: typ})
	}

	return fields, nil
}

func (e *Encoder) WriteNull() error {
	return e.WriteByte(commons.BC_NULL)
}

func (e *Encoder) findObjRef(obj interface{}) (int, bool) {
	for i, ref := range e.objRefs {
		if obj == ref {
			return i, true
		}
	}
	return -1, false
}

// # boolean true/false
//
//	::= 'T'
//	::= 'F'
func (e *Encoder) WriteBool(b bool) error {
	if b {
		return e.WriteByte(commons.BC_TRUE)
	} else {
		return e.WriteByte(commons.BC_FALSE)
	}
}

func (e *Encoder) WriteByte(b byte) error {
	e.buffer.WriteByte(b)
	return nil
}
func (e *Encoder) writeDouble(f float64) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	if err := e.WriteByte(commons.BC_DOUBLE); err != nil {
		return err
	}
	_, err := e.buffer.Write(buf[:])
	return err
}

func (e *Encoder) WriteInt(i int64) error {
	if i >= commons.INT_DIRECT_MIN && i <= int64(commons.INT_DIRECT_MAX) {
		// 1 byte
		return e.WriteByte(byte(i) + commons.BC_INT_ZERO)
	} else if i >= commons.INT_BYTE_MIN && i <= commons.INT_BYTE_MAX {
		// 2 bytes
		if err := e.WriteByte(byte(i>>8) + commons.BC_INT_ZERO); err != nil {
			return err
		}
		return e.WriteByte(byte(i))
	} else if i >= commons.INT_SHORT_MIN && i <= commons.INT_SHORT_MAX {
		// 3 bytes
		if err := e.WriteByte(byte(i>>16) + commons.BC_INT_SHORT_ZERO); err != nil {
			return err
		}
		if err := e.WriteByte(byte(i >> 8)); err != nil {
			return err
		}
		return e.WriteByte(byte(i))
	} else {
		// 8 bytes
		if err := e.WriteByte(commons.BC_LONG); err != nil {
			return err
		}
		if err := binary.Write(e.buffer, binary.BigEndian, i); err != nil {
			return err
		}
		return nil
	}
}

func (e *Encoder) WriteString(s string) error {
	// 将字符串转换为UTF-8编码的字节序列
	b := utils.StringToSliceByte(s)
	n := len(b)
	// write string tag
	if n < int(commons.BC_BINARY_DIRECT) {
		e.buffer.WriteByte(byte(n + int(commons.BC_BINARY_DIRECT)))
	} else if n <= commons.PACKET_SHORT_MAX {
		e.buffer.WriteByte(commons.BC_STRING)
		e.buffer.WriteByte(byte(n))
	} else {
		e.buffer.WriteByte(commons.BC_STRING_CHUNK)
		err := binary.Write(e.buffer, binary.BigEndian, uint32(n))
		if err != nil {
			return err
		}
	}
	e.buffer.Write(b)
	return nil
}

func (e *Encoder) writeBytes(b []byte) error {
	if len(b) <= int(commons.BINARY_DIRECT_MAX) {
		// 1 byte
		if err := e.WriteByte(byte(len(b)) + commons.BC_BINARY_DIRECT); err != nil {
			return err
		}
	} else {
		// 2 bytes
		if err := e.WriteByte(commons.BC_BINARY); err != nil {
			return err
		}
		if err := binary.Write(e.buffer, binary.BigEndian, int16(len(b))); err != nil {
			return err
		}
	}
	_, err := e.buffer.Write(b)
	return err
}

func (e *Encoder) writeList(l []interface{}) error {
	if err := e.WriteByte(commons.BC_LIST_FIXED); err != nil {
		return err
	}
	if err := binary.Write(e.buffer, binary.BigEndian, int32(len(l))); err != nil {
		return err
	}
	for _, v := range l {
		if err := e.WriteObject(v); err != nil {
			return err
		}
	}
	return nil
}

func (e *Encoder) writeMap(m map[string]interface{}) error {
	if err := e.WriteByte(commons.BC_MAP); err != nil {
		return err
	}
	if err := binary.Write(e.buffer, binary.BigEndian, int32(len(m))); err != nil {
		return err
	}
	for k, v := range m {
		if err := e.WriteString(k); err != nil {
			return err
		}
		if err := e.WriteObject(v); err != nil {
			return err
		}
	}
	return nil
}
