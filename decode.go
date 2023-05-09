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
	"encoding/binary"
	"io"
	"math"
)

type HessianDecoder struct {
	reader io.Reader
}

func NewHessianDecoder(reader io.Reader) *HessianDecoder {
	return &HessianDecoder{reader: reader}
}

func (d *HessianDecoder) ReadBoolean() (bool, error) {
	value, err := d.readByte()
	if err != nil {
		return false, err
	}
	return value == 'T', nil
}

func (d *HessianDecoder) ReadInt() (int32, error) {
	var result int32
	var shift uint
	for {
		b, err := d.readByte()
		if err != nil {
			return 0, err
		}
		result |= (int32(b) & 0x7f) << shift
		if b&0x80 == 0 {
			if shift < 31 && b&0x40 != 0 {
				result |= -1 << (shift + 7)
			}
			break
		}
		shift += 7
	}
	return result, nil
}

func (d *HessianDecoder) ReadLong() (int64, error) {
	var result int64
	var shift uint
	for {
		b, err := d.readByte()
		if err != nil {
			return 0, err
		}
		result |= (int64(b) & 0x7f) << shift
		if b&0x80 == 0 {
			if shift < 63 && b&0x40 != 0 {
				result |= -1 << (shift + 7)
			}
			break
		}
		shift += 7
	}
	return result, nil
}

func (d *HessianDecoder) ReadDouble() (float64, error) {
	bits, err := d.readUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

func (d *HessianDecoder) ReadString() (string, error) {
	length, err := d.ReadInt()
	if err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	_, err = io.ReadFull(d.reader, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func (d *HessianDecoder) readByte() (byte, error) {
	var b [1]byte
	_, err := io.ReadFull(d.reader, b[:])
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (d *HessianDecoder) readUint16() (uint16, error) {
	var b [2]byte
	_, err := io.ReadFull(d.reader, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b[:]), nil
}

func (d *HessianDecoder) readUint32() (uint32, error) {
	var b [4]byte
	_, err := io.ReadFull(d.reader, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(b[:]), nil
}

func (d *HessianDecoder) readUint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(d.reader, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[:]), nil
}

func (d *HessianDecoder) ReadList() ([]interface{}, error) {
	length, err := d.ReadInt()
	if err != nil {
		return nil, err
	}
	list := make([]interface{}, length)
	for i := int32(0); i < length; i++ {
		item, err := d.ReadObject()
		if err != nil {
			return nil, err
		}
		list[i] = item
	}
	return list, nil
}

func (d *HessianDecoder) ReadMap() (map[interface{}]interface{}, error) {
	length, err := d.ReadInt()
	if err != nil {
		return nil, err
	}
	hMap := make(map[interface{}]interface{})
	for i := int32(0); i < length; i++ {
		key, err := d.ReadObject()
		if err != nil {
			return nil, err
		}
		value, err := d.ReadObject()
		if err != nil {
			return nil, err
		}
		hMap[key] = value
	}
	return hMap, nil
}

func (d *HessianDecoder) ReadObject() (interface{}, error) {
	b, err := d.readByte()
	if err != nil {
		return nil, err
	}
	switch {
	case b == 'N':
		return nil, nil
	case b == 'T' || b == 'F':
		return d.ReadBoolean()
	case b == 'I':
		return d.ReadInt()
	case b == 'L':
		return d.ReadLong()
	case b == 'D':
		return d.ReadDouble()
	case b == 'S' || b == 'X':
		return d.ReadString()
	case b == 'V':
		return d.ReadList()
	case b == 'M':
		return d.ReadMap()
	default:
		return nil, nil
	}
}
