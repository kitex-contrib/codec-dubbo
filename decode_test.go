package codec_hessian2

import (
	"io"
	"reflect"
	"testing"
)

func TestHessianDecoder_ReadBoolean(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadBoolean()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBoolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadBoolean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadDouble(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadDouble()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadDouble() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadInt(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadList(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadList()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadLong(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadLong()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadLong() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadMap(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[interface{}]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadObject(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadObject()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_ReadString(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.ReadString()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_readByte(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.readByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("readByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_readUint16(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.readUint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("readUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_readUint32(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.readUint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("readUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHessianDecoder_readUint64(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &HessianDecoder{
				reader: tt.fields.reader,
			}
			got, err := d.readUint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("readUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHessianDecoder(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want *HessianDecoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHessianDecoder(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHessianDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
