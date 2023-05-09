// Copyright 2021 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec_hessian2

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

// Hessian2Encoder is a Hessian 2 encoder.
type Hessian2Encoder struct {
	writer io.Writer
}

// NewHessian2Encoder returns a new Hessian 2 encoder.
func NewHessian2Encoder(writer io.Writer) *Hessian2Encoder {
	return &Hessian2Encoder{writer: writer}
}

// Encode encodes the given value in Hessian 2 format and writes it to the underlying writer.
func (e *Hessian2Encoder) Encode(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return e.writeNil()
	case bool:
		return e.writeBool(v)
	case int:
		return e.writeInt(int64(v))
	case int8:
		return e.writeInt(int64(v))
	case int16:
		return e.writeInt(int64(v))
	case int32:
		return e.writeInt(int64(v))
	case int64:
		return e.writeInt(v)
	case uint:
		return e.writeInt(int64(v))
	case uint8:
		return e.writeInt(int64(v))
	case uint16:
		return e.writeInt(int64(v))
	case uint32:
		return e.writeInt(int64(v))
	case uint64:
		return e.writeInt(int64(v))
	case float32:
		return e.writeFloat(float64(v))
	case float64:
		return e.writeFloat(v)
	case string:
		return e.writeString(v)
	case []byte:
		return e.writeBytes(v)
	case []interface{}:
		return e.writeList(v)
	case map[string]interface{}:
		return e.writeMap(v)
	default:
		return nil
	}
}

// writeNil writes a null value to the underlying writer.
func (e *Hessian2Encoder) writeNil() error {
	_, err := e.writer.Write([]byte{0x4e})
	return err
}

// writeBool writes a boolean value to the underlying writer.
func (e *Hessian2Encoder) writeBool(value bool) error {
	var b byte
	if value {
		b = 0x54
	} else {
		b = 0x46
	}
	_, err := e.writer.Write([]byte{b})
	return err
}

// writeInt writes an integer value to the underlying writer.
func (e *Hessian2Encoder) writeInt(value int64) error {
	if value >= -0x08 && value <= 0x0f {
		b := byte(value + 0x90)
		_, err := e.writer.Write([]byte{b})
		return err
	} else if value >= -0x800 && value <= 0x7ff {
		b1 := byte((value >> 8) + 0xc8)
		b2 := byte(value & 0xff)
		_, err := e.writer.Write([]byte{b1, b2})
		return err
	} else if value >= -0x40000 && value <= 0x3ffff {
		b1 := byte((value >> 16) + 0xd4)
		b2 := byte((value >> 8) & 0xff)
		b3 := byte(value & 0xff)
		_, err := e.writer.Write([]byte{b1, b2, b3})
		return err
	} else if value >= -0x80000000 && value <= 0x7fffffff {
		b1 := byte((value >> 24) + 0xe0)
		b2 := byte((value >> 16) & 0xff)
		b3 := byte((value >> 8) & 0xff)
		b4 := byte(value & 0xff)
		_, err := e.writer.Write([]byte{b1, b2, b3, b4})
		return err
	} else {
		return nil
	}
}

// writeFloat writes a floating point value to the underlying writer.
func (e *Hessian2Encoder) writeFloat(value float64) error {
	if math.IsNaN(value) {
		_, err := e.writer.Write([]byte{0x46})
		return err
	} else if math.IsInf(value, 0) {
		if value > 0 {
			_, err := e.writer.Write([]byte{0x49})
			return err
		} else {
			_, err := e.writer.Write([]byte{0x4a})
			return err
		}
	} else {
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, value)
		if err != nil {
			return err
		}
		data := buf.Bytes()
		length := len(data)
		var b byte
		if length <= 0x0f {
			b = byte(0x4d + length)
			_, err = e.writer.Write([]byte{b})
		} else {
			b = 0x5b
			_, err = e.writer.Write([]byte{b, byte(length >> 8), byte(length & 0xff)})
		}
		if err != nil {
			return err
		}
		_, err = e.writer.Write(data)
		return err
	}
}

// writeString writes a string value to the underlying writer.
func (e *Hessian2Encoder) writeString(value string) error {
	data := []byte(value)
	length := len(data)
	var b byte
	if length <= 0x0f {
		b = byte(0x00 + length)
		_, err := e.writer.Write([]byte{b})
		if err != nil {
			return err
		}
	} else {
		b = 0x53
		_, err := e.writer.Write([]byte{b, byte(length >> 8), byte(length & 0xff)})
		if err != nil {
			return err
		}
	}
	_, err := e.writer.Write(data)
	return err
}

// writeBytes writes a byte slice to the underlying writer.
func (e *Hessian2Encoder) writeBytes(value []byte) error {
	length := len(value)
	var b byte
	if length <= 0x0f {
		b = byte(0x20 + length)
		_, err := e.writer.Write([]byte{b})
		if err != nil {
			return err
		}
	} else {
		b = 0x41
		_, err := e.writer.Write([]byte{b, byte(length >> 16), byte((length >> 8) & 0xff), byte(length & 0xff)})
		if err != nil {
			return err
		}
	}
	_, err := e.writer.Write(value)
	return err
}

// writeList writes a list of values to the underlying writer.
func (e *Hessian2Encoder) writeList(values []interface{}) error {
	length := len(values)
	var b byte
	if length <= 0x0f {
		b = byte(0x70 + length)
		_, err := e.writer.Write([]byte{b})
		if err != nil {
			return err
		}
	} else {
		b = 0x55
		_, err := e.writer.Write([]byte{b, byte(length >> 24), byte((length >> 16) & 0xff), byte((length >> 8) & 0xff), byte(length & 0xff)})
		if err != nil {
			return err
		}
	}
	for _, value := range values {
		err := e.Encode(value)
		if err != nil {
			return err
		}
	}
	return nil
}

// writeMap writes a map of string keys to values to the underlying writer.
func (e *Hessian2Encoder) writeMap(values map[string]interface{}) error {
	var b byte
	length := len(values)
	if length <= 0x0f {
		b = byte(0x60 + length)
		_, err := e.writer.Write([]byte{b})
		if err != nil {
			return err
		}
	} else {
		b = 0x4d
		_, err := e.writer.Write([]byte{b, byte(length >> 8), byte(length & 0xff)})
		if err != nil {
			return err
		}
	}
	for key, value := range values {
		err := e.writeString(key)
		if err != nil {
			return err
		}
		err = e.Encode(value)
		if err != nil {
			return err
		}
	}
	return nil
}
