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
	"io"
	"reflect"
	"testing"
)

func TestEncoder_WriteBool(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		b bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    byte
	}{
		{
			name: "TRUE",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args:    args{b: false},
			wantErr: false,
			want:    'F',
		}, {
			name: "FALSE",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args:    args{b: true},
			wantErr: false,
			want:    'T',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.WriteBool(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("WriteBool() error = %v, wantErr %v", err, tt.wantErr)
			}
			buffer := e.buffer
			readByte, err := buffer.ReadByte()
			if readByte != tt.want {
				t.Errorf("WriteBool() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestEncoder_WriteByte(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		b byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    byte
	}{
		{
			name: "a",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args:    args{b: 'a'},
			wantErr: false,
			want:    'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}

			if err := e.WriteByte(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("WriteByte() error = %v, wantErr %v", err, tt.wantErr)
			}
			readByte, _ := e.buffer.ReadByte()
			if readByte != tt.want {
				t.Errorf("WriteBool() want = %v, res %v", tt.want, readByte)
			}
		})
	}
}

func TestEncoder_WriteInt(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		i int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    byte
	}{
		{
			name: "int",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args:    args{i: int64(1)},
			wantErr: false,
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}

			if err := e.WriteInt(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("WriteInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			readByte, _ := e.buffer.ReadByte()
			if readByte != tt.want {
				t.Errorf("WriteBool() want = %v, res %v", tt.want, readByte)
			}
		})
	}
}

func TestEncoder_WriteObject(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "string",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args: args{
				interface{}("唯德"),
			},
			wantErr: false,
			want:    "唯德",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.WriteObject(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("WriteInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			str, _ := e.buffer.ReadString(1)
			if str != tt.want {
				t.Errorf("WriteBool() want = %v, res = %v", tt.want, str)
			}
		})
	}
}

func TestEncoder_findObjRef(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			got, got1 := e.findObjRef(tt.args.obj)
			if got != tt.want {
				t.Errorf("findObjRef() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findObjRef() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEncoder_findTypeRef(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		clazz *Class
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			got, got1 := e.findTypeRef(tt.args.clazz)
			if got != tt.want {
				t.Errorf("findTypeRef() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTypeRef() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEncoder_writeBytes(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeBytes(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("writeBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeClass(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		clazz *Class
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeClass(tt.args.clazz); (err != nil) != tt.wantErr {
				t.Errorf("writeClass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeDouble(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		f float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeDouble(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("writeDouble() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeField(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeField(tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("writeField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeFields(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		fields []*Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeFields(tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("writeFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeList(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		l []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeList(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("writeList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeMap(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeMap(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("writeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeObjectRef(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeObjectRef(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("writeinterface{}Ref() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeString(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "normal_string",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args: args{
				"hello",
			},
			wantErr: false,
			want:    "%hello",
		}, {
			name: "blank_string",
			fields: fields{
				writer: io.MultiWriter(),
				buffer: bytes.NewBuffer([]byte{}),
			},
			args: args{
				"",
			},
			wantErr: false,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.WriteString(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("writeString() error = %v, wantErr %v", err, tt.wantErr)
			}
			str, _ := e.buffer.ReadString(1)
			if str != tt.want {
				t.Errorf("writeString() want = %v, res = %v", tt.want, str)
			}

		})
	}
}

func TestEncoder_writeTypeRef(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		ref int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeTypeRef(tt.args.ref); (err != nil) != tt.wantErr {
				t.Errorf("writeTypeRef() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_writeTypeRefOrClass(t *testing.T) {
	type fields struct {
		writer         io.Writer
		buffer         *bytes.Buffer
		writeType      bool
		typeRefWritten bool
		typeRefs       []*Class
		objRefs        []interface{}
	}
	type args struct {
		clazz *Class
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				writer:         tt.fields.writer,
				buffer:         tt.fields.buffer,
				writeType:      tt.fields.writeType,
				typeRefWritten: tt.fields.typeRefWritten,
				typeRefs:       tt.fields.typeRefs,
				objRefs:        tt.fields.objRefs,
			}
			if err := e.writeTypeRefOrClass(tt.args.clazz); (err != nil) != tt.wantErr {
				t.Errorf("writeTypeRefOrClass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHessian2Output(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
		want  *Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got := NewHessian2Output(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewHessian2Output() gotW = %v, want %v", gotW, tt.wantW)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHessian2Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getClass(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Class
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClass(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getClassFields(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    []*Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClassFields(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClassFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClassFields() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getClassName(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClassName(tt.args.t); got != tt.want {
				t.Errorf("getClassName() = %v, want %v", got, tt.want)
			}
		})
	}
}
