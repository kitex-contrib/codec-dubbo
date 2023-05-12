package codec_hessian2

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestHessian2Encoder_Encode(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value interface{}
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.Encode(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeBool(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value bool
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeBool(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("writeBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeBytes(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value []byte
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeBytes(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("writeBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeFloat(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value float64
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeFloat(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("writeFloat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeInt(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value int64
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeInt(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("writeInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeList(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		values []interface{}
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeList(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("writeList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeMap(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		values map[string]interface{}
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeMap(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("writeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeNil(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeNil(); (err != nil) != tt.wantErr {
				t.Errorf("writeNil() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHessian2Encoder_writeString(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		value string
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
			e := &Hessian2Encoder{
				writer: tt.fields.writer,
			}
			if err := e.writeString(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("writeString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHessian2Encoder(t *testing.T) {
	tests := []struct {
		name       string
		wantWriter string
		want       *Hessian2Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			got := NewHessian2Encoder(writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("NewHessian2Encoder() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHessian2Encoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
