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

package encode

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

// Class表示Java对象的类型信息
type Class struct {
	Name   string   // 类名
	Fields []*Field // 字段信息
}

// Field表示Java对象的字段信息
type Field struct {
	Name    string // 字段名
	Type    *Class // 字段类型
	Default Object // 字段默认值
}

// Object表示Java对象的值信息
type Object interface{}

type Encoder struct {
	writer io.Writer     // 输出流
	buffer *bytes.Buffer // buffer缓存

	writeType      bool     // 是否写入类型信息
	typeRefWritten bool     // 是否已经将类型引用写入
	typeRefs       []*Class // 类型引用数组
	objRefs        []Object // 对象引用数组
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
func (e *Encoder) WriteObject(obj Object) error {
	// 判断是否为nil
	if obj == nil {
		return e.WriteNull()
	}

	// 获取对象类型信息
	clazz, err := getClass(obj)
	if err != nil {
		return err
	}

	// 写入类型信息
	if e.writeType {
		if err := e.writeClass(clazz); err != nil {
			return err
		}
	}

	// 写入对象引用
	if ref, ok := e.findObjRef(obj); ok {
		return e.writeObjectRef(ref)
	}

	// 将对象添加到引用数组
	e.objRefs = append(e.objRefs, obj)

	// 判断类型并编码
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

	// 编码复杂类型
	switch v := obj.(type) {
	case []byte:
		return e.writeBytes(v)
	case []interface{}:
		return e.writeList(v)
	case map[string]interface{}:
		return e.writeMap(v)
	}

	// 构造并编码Java对象
	return e.writeClass(clazz)
}

// 获取Java对象的类型信息
func getClass(obj Object) (*Class, error) {
	if obj == nil {
		return nil, fmt.Errorf("getClass: nil object")
	}

	// 获取对象的类型
	t := reflect.TypeOf(obj)

	// 获取对象的类名
	name := getClassName(t)
	if name == "" {
		return nil, fmt.Errorf("getClass: invalid object type: %v", t)
	}

	// 如果是struct类型，则根据结构体字段构造Class对象
	if t.Kind() == reflect.Struct {
		fields, err := getClassFields(t)
		if err != nil {
			return nil, err
		}
		return &Class{Name: name, Fields: fields}, nil
	}

	// 其他基本类型直接构造Class对象
	return &Class{Name: name}, nil
}

// 获取对象的类名
func getClassName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Struct:
		if t.NumField() == 0 {
			return ""
		}
		// 结构体类型的类名为结构体名的小写形式
		return strings.ToLower(t.Name())
	case reflect.Slice:
		// 判断是否为字节数组
		if t.Elem().Kind() == reflect.Uint8 {
			return "byte[]"
		}
		// 其他slice类型的类名为List
		return "list"
	case reflect.Map:
		// map类型的类名为Map
		return "map"
	case reflect.Ptr:
		// 指针类型递归调用获取类名
		return getClassName(t.Elem())
	default:
		// 其他基本类型的类名为类型名的小写形式
		return strings.ToLower(t.Name())
	}
}

// 将Java对象类型信息编码为commons协议格式
func (e *Encoder) writeClass(clazz *Class) error {
	// 查找类型引用
	if ref, ok := e.findTypeRef(clazz); ok {
		return e.writeTypeRef(ref)
	}

	// 将类型添加到引用数组
	e.typeRefs = append(e.typeRefs, clazz)

	// 写入类型标识
	if err := e.WriteByte(byte('C')); err != nil {
		return err
	}

	// 写入类名
	if err := e.WriteString(clazz.Name); err != nil {
		return err
	}

	// 写入类型信息
	if err := e.writeFields(clazz.Fields); err != nil {
		return err
	}

	// 标记已经写入类型信息
	e.typeRefWritten = true

	return nil
}

// 查找类型引用
func (e *Encoder) findTypeRef(clazz *Class) (int64, bool) {
	for i, ref := range e.typeRefs {
		if reflect.DeepEqual(clazz, ref) {
			return int64(i), true
		}
	}
	return -1, false
}

// 写入类型引用
func (e *Encoder) writeTypeRef(ref int64) error {
	if err := e.WriteByte(byte('T')); err != nil {
		return err
	}
	if err := e.WriteInt(ref); err != nil {
		return err
	}
	return nil
}

// 写入字段信息
func (e *Encoder) writeFields(fields []*Field) error {
	for _, field := range fields {
		if err := e.writeField(field); err != nil {
			return err
		}
	}
	return nil
}

// 写入单个字段信息
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

// 写入类型引用或类名
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

// 获取结构体字段信息
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
	return nil
}

func (e *Encoder) findObjRef(obj interface{}) (int, bool) {
	for i, ref := range e.objRefs {
		if obj == ref {
			return i, true
		}
	}
	return -1, false
}

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
		// 一字节表示
		return e.WriteByte(byte(i) + commons.BC_INT_ZERO)
	} else if i >= commons.INT_BYTE_MIN && i <= commons.INT_BYTE_MAX {
		// 两字节表示
		if err := e.WriteByte(byte(i>>8) + commons.BC_INT_ZERO); err != nil {
			return err
		}
		return e.WriteByte(byte(i))
	} else if i >= commons.INT_SHORT_MIN && i <= commons.INT_SHORT_MAX {
		// 三字节表示
		if err := e.WriteByte(byte(i>>16) + commons.BC_INT_SHORT_ZERO); err != nil {
			return err
		}
		if err := e.WriteByte(byte(i >> 8)); err != nil {
			return err
		}
		return e.WriteByte(byte(i))
	} else {
		// 八字节表示
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

	b := []byte(s)
	// 获取字符串长度
	n := len(b)
	// 写入字符串类型标记
	if n < int(commons.BC_BINARY_DIRECT) {
		e.buffer.WriteByte(byte(n + int(commons.BC_BINARY_DIRECT)))
	} else if n <= commons.PACKET_SHORT_MAX {
		e.buffer.WriteByte(commons.BC_STRING)
		e.buffer.WriteByte(byte(n))
	} else {
		e.buffer.WriteByte(commons.BC_STRING_CHUNK)
		binary.Write(e.buffer, binary.BigEndian, uint32(n))
	}
	// 写入字符串的字节序列
	e.buffer.Write(b)
	return nil
}

func (e *Encoder) writeBytes(b []byte) error {
	if len(b) <= int(commons.BINARY_DIRECT_MAX) {
		// 一字节长度表示
		if err := e.WriteByte(byte(len(b)) + commons.BC_BINARY_DIRECT); err != nil {
			return err
		}
	} else {
		// 两字节长度表示
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
