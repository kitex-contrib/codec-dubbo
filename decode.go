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

package hessian2

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	ReadSize = 1024

	EndOfData = -2
)

func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{
		in:     in,
		buffer: make([]byte, ReadSize),
		sbuf:   &strings.Builder{},
	}
}

type Decoder struct {
	in io.Reader

	length int
	offset int

	buffer []byte

	method      *string
	chunkLength int
	isLastChunk bool

	refs      []interface{}
	types     []string
	classDefs []*ObjectDefinition

	sbuf *strings.Builder
}

func (d *Decoder) read() byte {
	if d.length <= d.offset && !d.readBuffer() {
		return 0xFF
	}
	r := d.buffer[d.offset]
	d.offset++
	return r
}

func (d *Decoder) readBuffer() bool {
	offset, length := d.offset, d.length
	if offset < length {
		copy(d.buffer, d.buffer[offset:length])
		offset = length - offset
	} else {
		offset = 0
	}
	l, err := d.in.Read(d.buffer[offset:])
	d.offset = 0
	if l <= 0 || err != nil {
		d.length = offset
		return offset > 0
	}
	d.length = offset + l
	return true
}

func (d *Decoder) readNull() {
	tag := d.read()
	if tag != 'N' {
		panic(d.expect("null", tag))
	}
}

func (d *Decoder) readTag() byte {
	if d.offset < d.length {
		b := d.buffer[d.offset]
		d.offset++
		return b
	} else {
		return d.read()
	}
}

func (d *Decoder) readBoolean() bool {
	tag := d.readTag()
	switch tag {
	case 'T':
		return true
	case 'F':
		return false
		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,
		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,
		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,
		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		return tag != BC_INT_ZERO
		// INT_BYTE = 0
	case 0xc8:
		return d.read() != 0
		// INT_BYTE != 0
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		d.read()
		return true
		// INT_SHORT = 0
	case 0xd4:
		return (256*int16(d.read()) + int16(d.read())) != 0
		// INT_SHORT != 0
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd5, 0xd6, 0xd7:
		d.read()
		d.read()
		return true
	case 'I':
		return d.parseInt() != 0
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,
		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		return tag != BC_LONG_ZERO
		// LONG_BYTE = 0
	case 0xf8:
		return d.read() != 0
		// LONG_BYTE != 0
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		d.read()
		return true
		// INT_SHORT = 0
	case 0x3c:
		return (256*int16(d.read()) + int16(d.read())) != 0
		// INT_SHORT != 0
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3d, 0x3e, 0x3f:
		d.read()
		d.read()
		return true
	case BC_LONG_INT:
		v := 0x1000000*int64(d.read()) + 0x10000*int64(d.read()) + 0x100*int64(d.read()) + int64(d.read())
		return v != 0
	case 'L':
		return d.parseLong() != 0
	case BC_DOUBLE_ZERO:
		return false

	case BC_DOUBLE_ONE:
		return true

	case BC_DOUBLE_BYTE:
		return d.read() != 0
	case BC_DOUBLE_SHORT:
		return (0x100*int16(d.read()) + int16(d.read())) != 0
	case BC_DOUBLE_MILL:
		return d.parseInt() != 0
	case 'D':
		return d.parseDouble() != 0.0
	case 'N':
		return false
	default:
		panic(d.expect("boolean", tag))
	}
}

func (d *Decoder) readShort() int16 {
	return int16(d.readInt())
}

func (d *Decoder) readInt() int {
	tag := d.read()
	switch tag {
	case 'N':
		return 0

	case 'F':
		return 0

	case 'T':
		return 1

		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,

		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,

		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,

		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		return int(tag) - BC_INT_ZERO

		/* byte int */
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc8, 0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

		/* short int */
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd4, 0xd5, 0xd6, 0xd7:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case 'I',
		BC_LONG_INT:
		return (int(d.read()) << 24) + (int(d.read()) << 16) + (int(d.read()) << 8) + int(d.read())

		// direct long
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,

		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		return int(tag) - BC_LONG_ZERO

		/* byte long */
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf8, 0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		return ((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read())

		/* short long */
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3c, 0x3d, 0x3e, 0x3f:
		return ((int(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case 'L':
		return int(d.parseLong())

	case BC_DOUBLE_ZERO:
		return 0

	case BC_DOUBLE_ONE:
		return 1

		//case LONG_BYTE:
	case BC_DOUBLE_BYTE:
		if d.offset < d.length {
			return int(d.buffer[d.offset])
			d.offset++
		} else {
			return int(d.read())
		}

		//case INT_SHORT:
		//case LONG_SHORT:
	case BC_DOUBLE_SHORT:
		return int(256*int16(d.read()) + int16(d.read()))

	case BC_DOUBLE_MILL:
		return d.parseInt() / 1000

	case 'D':
		return int(d.parseDouble())

	default:
	}
	panic(d.expect("integer", tag))
}

func (d *Decoder) readLong() int64 {
	tag := d.read()
	switch tag {
	case 'N':
		return 0

	case 'F':
		return 0

	case 'T':
		return 1

		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,

		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,

		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,

		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		return int64(tag) - BC_INT_ZERO

		/* byte int */
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc8, 0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		return ((int64(tag) - BC_INT_BYTE_ZERO) << 8) + int64(d.read())

		/* short int */
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd4, 0xd5, 0xd6, 0xd7:
		return ((int64(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

		//case LONG_BYTE:
	case BC_DOUBLE_BYTE:
		if d.offset < d.length {
			return int64(d.buffer[d.offset])
			d.offset++
		} else {
			return int64(d.read())
		}

		//case INT_SHORT:
		//case LONG_SHORT:
	case BC_DOUBLE_SHORT:
		return 256*int64(d.read()) + int64(d.read())

	case 'I', BC_LONG_INT:
		return int64(d.parseInt())

		// direct long
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,

		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		return int64(tag) - BC_LONG_ZERO

		/* byte long */
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf8, 0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

		/* short long */
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3c, 0x3d, 0x3e, 0x3f:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

	case 'L':
		return d.parseLong()

	case BC_DOUBLE_ZERO:
		return 0

	case BC_DOUBLE_ONE:
		return 1

	case BC_DOUBLE_MILL:
		return int64(d.parseInt() / 1000)

	case 'D':
		return int64(d.parseDouble())

	default:
	}
	panic(d.expect("long", tag))
}

func (d *Decoder) readFloat() float32 {
	return float32(d.readDouble())
}

func (d *Decoder) readDouble() float64 {
	tag := d.read()
	switch tag {
	case 'N':
		return 0

	case 'F':
		return 0

	case 'T':
		return 1

		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,

		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,

		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,

		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		return float64(tag) - 0x90

		/* byte int */
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc8, 0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		return float64((int64(tag)-BC_INT_BYTE_ZERO)<<8) + float64(d.read())

		/* short int */
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd4, 0xd5, 0xd6, 0xd7:
		return float64((int64(tag)-BC_INT_SHORT_ZERO)<<16) + 256*float64(d.read()) + float64(d.read())

	case 'I', BC_LONG_INT:
		return float64(d.parseInt())

		// direct long
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,

		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		return float64(tag) - BC_LONG_ZERO

		/* byte long */
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf8, 0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		return float64((int64(tag)-BC_LONG_BYTE_ZERO)<<8) + float64(d.read())

		/* short long */
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3c, 0x3d, 0x3e, 0x3f:
		return float64((int64(tag)-BC_LONG_SHORT_ZERO)<<16) + 256*float64(d.read()) + float64(d.read())

	case 'L':
		return float64(d.parseLong())

	case BC_DOUBLE_ZERO:
		return 0

	case BC_DOUBLE_ONE:
		return 1

	case BC_DOUBLE_BYTE:
		if d.offset < d.length {
			return float64(d.buffer[d.offset])
			d.offset++
		} else {
			return float64(d.read())
		}

	case BC_DOUBLE_SHORT:
		return 256*float64(d.read()) + float64(d.read())

	case BC_DOUBLE_MILL:
		return float64(d.parseInt() / 1000)

	case 'D':
		return d.parseDouble()

	default:
	}
	panic(d.expect("double", tag))
}

func (d *Decoder) readUTCDate() int64 {
	tag := d.read()
	if tag == BC_DATE {
		return d.parseLong()
	} else if tag == BC_DATE_MINUTE {
		return int64(d.parseInt())
	} else {
		panic(d.expect("date", tag))
	}
}

func (d *Decoder) readChar() rune {
	if d.chunkLength > 0 {
		d.chunkLength--
		if d.chunkLength == 0 && d.isLastChunk {
			d.chunkLength = EndOfData
		}

		return d.parseUTF8Char()
	} else if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return rune(0xff)
	}

	tag := d.read()

	switch tag {
	case 'N':
		return rune(0xff)

	case 'S', BC_STRING_CHUNK:
		d.isLastChunk = tag == 'S'
		d.chunkLength = (int(d.read()) << 8) + int(d.read())
		d.chunkLength--
		value := d.parseUTF8Char()
		if d.chunkLength == 0 && d.isLastChunk {
			d.chunkLength = EndOfData
		}

		return value

	default:
		panic(d.expect("char", tag))
	}
}

func (d *Decoder) readStringToBuffer(buffer []rune, offset int, length int) int {
	if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return 0xff
	} else if d.chunkLength == 0 {
		tag := d.read()

		switch tag {
		case 'N':
			return 0xff

		case 'S':
		case BC_STRING_CHUNK:
			d.isLastChunk = tag == 'S'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case 0x00, 0x01, 0x02, 0x03,
			0x04, 0x05, 0x06, 0x07,
			0x08, 0x09, 0x0a, 0x0b,
			0x0c, 0x0d, 0x0e, 0x0f,

			0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17,
			0x18, 0x19, 0x1a, 0x1b,
			0x1c, 0x1d, 0x1e, 0x1f:
			d.isLastChunk = true
			d.chunkLength = int(tag)

		case 0x30, 0x31, 0x32, 0x33:
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x30)*256 + int(d.read())

		default:
			panic(d.expect("string", tag))
		}
	}

	readLength := 0

	for length > 0 {
		if d.chunkLength > 0 {
			buffer[offset] = d.parseUTF8Char()
			offset++
			d.chunkLength--
			length--
			readLength++
		} else if d.isLastChunk {
			if readLength == 0 {
				return -1
			} else {
				d.chunkLength = EndOfData
				return readLength
			}
		} else {
			tag := d.read()

			switch tag {
			case 'S', BC_STRING_CHUNK:
				d.isLastChunk = tag == 'S'
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case 0x00, 0x01, 0x02, 0x03,
				0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0a, 0x0b,
				0x0c, 0x0d, 0x0e, 0x0f,

				0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17,
				0x18, 0x19, 0x1a, 0x1b,
				0x1c, 0x1d, 0x1e, 0x1f:
				d.isLastChunk = true
				d.chunkLength = int(tag)

			case 0x30, 0x31, 0x32, 0x33:
				d.isLastChunk = true
				d.chunkLength = (int(tag)-0x30)*256 + int(d.read())

			default:
				panic(d.expect("string", tag))
			}
		}
	}
	if readLength == 0 {
		return -1
	} else if d.chunkLength > 0 || !d.isLastChunk {
		return readLength
	} else {
		d.chunkLength = EndOfData
		return readLength
	}
}

func (d *Decoder) readString() *string {
	tag := d.read()
	var val string
	switch tag {
	case 'N':
		return nil
	case 'T':
		val = "true"
		return &val
	case 'F':
		val = "false"
		return &val

		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,

		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,

		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,

		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		val = strconv.Itoa(int(tag) - 0x90)
		return &val

		/* byte int */
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc8, 0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		val = strconv.Itoa(((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read()))
		return &val

		/* short int */
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd4, 0xd5, 0xd6, 0xd7:
		val = strconv.Itoa(((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read()))
		return &val

	case 'I', BC_LONG_INT:
		val = strconv.Itoa(d.parseInt())
		return &val

		// direct long
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,

		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		val = strconv.Itoa(int(tag) - BC_LONG_ZERO)
		return &val

		/* byte long */
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf8, 0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		val = strconv.Itoa(((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read()))
		return &val

		/* short long */
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3c, 0x3d, 0x3e, 0x3f:
		val = strconv.Itoa(((int(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read()))
		return &val

	case 'L':
		val = fmt.Sprintf("%d", d.parseLong())
		return &val

	case BC_DOUBLE_ZERO:
		val = "0.0"
		return &val

	case BC_DOUBLE_ONE:
		val = "1.0"
		return &val

	case BC_DOUBLE_BYTE:
		val = strconv.Itoa(int(d.readTag()))
		return &val

	case BC_DOUBLE_SHORT:
		val = strconv.Itoa(int(int16(256*int(d.read()) + int(d.read()))))
		return &val

	case BC_DOUBLE_MILL:
		{
			mills := d.parseInt()
			val = strconv.Itoa(mills / 1000)
			return &val
		}

	case 'D':
		val = fmt.Sprintf("%v", d.parseDouble())
		return &val

	case 'S', BC_STRING_CHUNK:
		d.isLastChunk = tag == 'S'
		d.chunkLength = (int(d.read()) << 8) + int(d.read())
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == 0xff {
				break
			}
			d.sbuf.WriteRune(ch)
		}
		val = d.sbuf.String()
		return &val

		// 0-byte string
	case 0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,

		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f:
		d.isLastChunk = true
		d.chunkLength = int(tag)
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == 0xff {
				break
			}
			d.sbuf.WriteRune(ch)
		}

		val = d.sbuf.String()
		return &val

	case 0x30, 0x31, 0x32, 0x33:
		d.isLastChunk = true
		d.chunkLength = (int(tag)-0x30)*256 + int(d.read())
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == 0xff {
				break
			}
			d.sbuf.WriteRune(ch)
		}

		val = d.sbuf.String()
		return &val

	default:
		panic(d.expect("string", tag))
	}
}

func (d *Decoder) readBytes() []byte {
	tag := d.read()

	switch tag {
	case 'N':
		return nil

	case BC_BINARY,
		BC_BINARY_CHUNK:
		d.isLastChunk = tag == BC_BINARY
		d.chunkLength = (int(d.read()) << 8) + int(d.read())
		bos := bytes.NewBuffer([]byte{})

		var data byte
		for {
			data = d.parseByte()
			if data == 0xFF {
				break
			}
			bos.WriteByte(data)
		}

		return bos.Bytes()

	case 0x20,
		0x21,
		0x22,
		0x23,
		0x24,
		0x25,
		0x26,
		0x27,
		0x28,
		0x29,
		0x2a,
		0x2b,
		0x2c,
		0x2d,
		0x2e,
		0x2f:
		{
			d.isLastChunk = true
			d.chunkLength = int(tag) - 0x20

			buffer := make([]byte, 0, d.chunkLength)

			offset := 0
			for offset < d.chunkLength {
				sublen := d.read2(buffer, 0, d.chunkLength-offset)

				if sublen <= 0 {
					break
				}

				offset += sublen
			}

			return buffer
		}

	case 0x34,
		0x35,
		0x36,
		0x37:
		{
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x34)*256 + int(d.read())

			buffer := make([]byte, 0, d.chunkLength)

			offset := 0
			for offset < d.chunkLength {
				sublen := d.read2(buffer, 0, d.chunkLength-offset)

				if sublen <= 0 {
					break
				}

				offset += sublen
			}

			return buffer
		}

	default:
		panic(d.expect("bytes", tag))
	}
}

func (d *Decoder) readByte() byte {
	if d.chunkLength > 0 {
		d.chunkLength--
		if d.chunkLength == 0 && d.isLastChunk {
			d.chunkLength = EndOfData
		}

		return d.read()
	} else if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return 0xFF
	}

	tag := d.read()

	switch tag {
	case 'N':
		return 0xFF

	case 'B', BC_BINARY_CHUNK:
		{
			d.isLastChunk = tag == 'B'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			value := d.parseByte()

			// special code so successive read byte won't
			// be read as a single object.
			if d.chunkLength == 0 && d.isLastChunk {
				d.chunkLength = EndOfData
			}

			return value
		}

	case 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27,
		0x28, 0x29, 0x2a, 0x2b,
		0x2c, 0x2d, 0x2e, 0x2f:
		{
			d.isLastChunk = true
			d.chunkLength = int(tag) - 0x20

			value := d.parseByte()

			// special code so successive read byte won't
			// be read as a single object.
			if d.chunkLength == 0 {
				d.chunkLength = EndOfData
			}

			return value
		}

	case 0x34, 0x35, 0x36, 0x37:
		{
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x34)*256 + int(d.read())

			value := d.parseByte()

			// special code so successive read byte won't
			// be read as a single object.
			if d.chunkLength == 0 {
				d.chunkLength = EndOfData
			}

			return value
		}

	default:
		panic(d.expect("binary", tag))
	}
}

func (d *Decoder) readBytes2(buffer []byte, offset int, length int) int {
	if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return -1
	} else if d.chunkLength == 0 {
		tag := d.read()

		switch tag {
		case 'N':
			return -1

		case 'B', BC_BINARY_CHUNK:
			d.isLastChunk = tag == 'B'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case 0x20, 0x21, 0x22, 0x23,
			0x24, 0x25, 0x26, 0x27,
			0x28, 0x29, 0x2a, 0x2b,
			0x2c, 0x2d, 0x2e, 0x2f:

			d.isLastChunk = true
			d.chunkLength = int(tag) - 0x20

		case 0x34, 0x35, 0x36, 0x37:
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x34)*256 + int(d.read())

		default:
			panic(d.expect("binary", tag))
		}
	}
	readLength := 0
	for length > 0 {
		if d.chunkLength > 0 {
			buffer[offset] = d.read()
			offset++
			d.chunkLength--
			length--
			readLength++
		} else if d.isLastChunk {
			if readLength == 0 {
				return -1
			} else {
				d.chunkLength = EndOfData
				return readLength
			}
		} else {
			tag := d.read()

			switch tag {
			case 'B', BC_BINARY_CHUNK:
				d.isLastChunk = tag == 'B'
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			default:
				panic(d.expect("binary", tag))
			}
		}
	}

	if readLength == 0 {
		return -1
	} else if d.chunkLength > 0 || !d.isLastChunk {
		return readLength
	} else {
		d.chunkLength = EndOfData
		return readLength
	}
}

func (d *Decoder) ReadObject() interface{} {
	tag := d.readTag()
	switch tag {
	case 'N':
		return nil

	case 'T':
		return true

	case 'F':
		return false

		// direct integer
	case 0x80, 0x81, 0x82, 0x83,
		0x84, 0x85, 0x86, 0x87,
		0x88, 0x89, 0x8a, 0x8b,
		0x8c, 0x8d, 0x8e, 0x8f,

		0x90, 0x91, 0x92, 0x93,
		0x94, 0x95, 0x96, 0x97,
		0x98, 0x99, 0x9a, 0x9b,
		0x9c, 0x9d, 0x9e, 0x9f,

		0xa0, 0xa1, 0xa2, 0xa3,
		0xa4, 0xa5, 0xa6, 0xa7,
		0xa8, 0xa9, 0xaa, 0xab,
		0xac, 0xad, 0xae, 0xaf,

		0xb0, 0xb1, 0xb2, 0xb3,
		0xb4, 0xb5, 0xb6, 0xb7,
		0xb8, 0xb9, 0xba, 0xbb,
		0xbc, 0xbd, 0xbe, 0xbf:
		return int(tag) - BC_INT_ZERO

		/* byte int */
	case 0xc0, 0xc1, 0xc2, 0xc3,
		0xc4, 0xc5, 0xc6, 0xc7,
		0xc8, 0xc9, 0xca, 0xcb,
		0xcc, 0xcd, 0xce, 0xcf:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

		/* short int */
	case 0xd0, 0xd1, 0xd2, 0xd3,
		0xd4, 0xd5, 0xd6, 0xd7:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case 'I':
		return d.parseInt()

		// direct long
	case 0xd8, 0xd9, 0xda, 0xdb,
		0xdc, 0xdd, 0xde, 0xdf,

		0xe0, 0xe1, 0xe2, 0xe3,
		0xe4, 0xe5, 0xe6, 0xe7,
		0xe8, 0xe9, 0xea, 0xeb,
		0xec, 0xed, 0xee, 0xef:
		return int64(tag) - BC_LONG_ZERO

		/* byte long */
	case 0xf0, 0xf1, 0xf2, 0xf3,
		0xf4, 0xf5, 0xf6, 0xf7,
		0xf8, 0xf9, 0xfa, 0xfb,
		0xfc, 0xfd, 0xfe, 0xff:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

		/* short long */
	case 0x38, 0x39, 0x3a, 0x3b,
		0x3c, 0x3d, 0x3e, 0x3f:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

	case BC_LONG_INT:
		return int64(d.parseInt())

	case 'L':
		return d.parseLong()

	case BC_DOUBLE_ZERO:
		return 0.0

	case BC_DOUBLE_ONE:
		return 1.0

	case BC_DOUBLE_BYTE:
		return float64(d.read())

	case BC_DOUBLE_SHORT:
		return float64(256*int16(d.read()) + int16(d.read()))

	case BC_DOUBLE_MILL:
		{
			mills := float64(d.parseInt())

			return 0.001 * mills
		}

	case 'D':
		return d.parseDouble()

	case BC_DATE:
		l64 := d.parseLong()
		return time.Unix(l64/1000, l64%1000*10e5)

	case BC_DATE_MINUTE:
		return time.Unix(int64(d.parseInt())*60000, 0)

	case BC_STRING_CHUNK, 'S':
		{
			d.isLastChunk = tag == 'S'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case 0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,

		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f:
		{
			d.isLastChunk = true
			d.chunkLength = int(tag)

			d.sbuf.Reset()

			builder := d.sbuf
			d.parseString(d.sbuf)

			return builder.String()
		}

	case 0x30, 0x31, 0x32, 0x33:
		{
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x30)*256 + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case BC_BINARY_CHUNK, 'B':
		{
			d.isLastChunk = tag == 'B'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			bos := bytes.NewBuffer([]byte{})

			var data byte
			for {
				data = d.parseByte()
				if data == 0xFF {
					break
				}
				bos.WriteByte(data)
			}

			return bos.Bytes()
		}

	case 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27,
		0x28, 0x29, 0x2a, 0x2b,
		0x2c, 0x2d, 0x2e, 0x2f:
		{
			d.isLastChunk = true
			length := int(tag) - 0x20
			d.chunkLength = 0

			data := make([]byte, length)

			for i := 0; i < length; i++ {
				data[i] = d.read()
			}

			return data
		}

	case 0x34, 0x35, 0x36, 0x37:
		{
			d.isLastChunk = true
			length := (int(tag)-0x34)*256 + int(d.read())
			d.chunkLength = 0

			buffer := make([]byte, 0, length)

			for i := 0; i < length; i++ {
				buffer[i] = d.read()
			}

			return buffer
		}

	case BC_LIST_VARIABLE:
		{
			typ := d.readType()
			return d.readList(-1, &typ)
		}

	case BC_LIST_VARIABLE_UNTYPED:
		{
			return d.readList(-1, nil)
		}

	case BC_LIST_FIXED:
		{
			typ := d.readType()
			length := d.readInt()
			return d.readList(length, &typ)
		}

	case BC_LIST_FIXED_UNTYPED:
		{
			length := d.readInt()
			return d.readList(length, nil)
		}

		// compact fixed list
	case 0x70, 0x71, 0x72, 0x73,
		0x74, 0x75, 0x76, 0x77:
		{
			typ := d.readType()
			length := int(tag) - 0x70
			return d.readList(length, &typ)
		}

		// compact fixed untyped list
	case 0x78, 0x79, 0x7a, 0x7b,
		0x7c, 0x7d, 0x7e, 0x7f:
		{
			length := int(tag) - 0x78
			return d.readList(length, nil)
		}

	case 'H':
		{
			return d.readMap(nil)
		}

	case 'M':
		{
			typ := d.readType()
			return d.readMap(&typ)
		}

	case 'C':
		{
			d.readObjectDefinition(nil)

			return d.ReadObject()
		}

	case 0x60, 0x61, 0x62, 0x63,
		0x64, 0x65, 0x66, 0x67,
		0x68, 0x69, 0x6a, 0x6b,
		0x6c, 0x6d, 0x6e, 0x6f:
		{
			ref := int(tag) - 0x60

			if len(d.classDefs) <= ref {
				panic(d.error(fmt.Sprintf("No classes defined at reference '%s'", strings.ToUpper(hex.EncodeToString([]byte{tag})))))
			}

			def := d.classDefs[ref]

			return d.readObjectInstance(nil, def)
		}

	case 'O':
		{
			ref := d.readInt()

			if len(d.classDefs) <= ref {

				panic(d.error(fmt.Sprintf("Illegal object reference #%d", ref)))
			}

			def := d.classDefs[ref]

			return d.readObjectInstance(nil, def)
		}

	case BC_REF:
		{
			ref := d.readInt()

			return d.refs[ref]
		}

	default:
		if tag == 0xff {
			panic(fmt.Errorf("readObject: unexpected end of file"))
		} else {
			panic(d.error("readObject: unknown code " + d.codeName(tag)))
		}
	}
}

func (d *Decoder) readList(length int, typ *string) []interface{} {
	list := make([]interface{}, length)
	d.AddRef(list)
	for i := range list {
		list[i] = d.ReadObject()
	}
	return list
}

func (d *Decoder) readMap(typ *string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	d.AddRef(m)
	for !d.IsEnd() {
		m[d.ReadObject()] = d.ReadObject()
	}
	d.readEnd()
	return m
}

func (d *Decoder) readObjectDefinition(i interface{}) {
	typ := d.readString()
	length := d.readInt()

	fieldNames := make([]string, length)

	for i := 0; i < length; i++ {
		name := d.readString()

		fieldNames[i] = *name
	}

	def := NewObjectDefinition(*typ, fieldNames)
	d.classDefs = append(d.classDefs, def)
}

func (d *Decoder) readObjectInstance(cl interface{}, def *ObjectDefinition) interface{} {
	vc := NewVirtualClass(def.typ, def.fieldNames)
	for _, key := range def.fieldNames {
		v := d.ReadObject()
		vc.fields[key] = v
	}
	return vc
}

func (d *Decoder) readRef() interface{} {
	value := d.parseInt()
	return d.refs[value]
}

func (d *Decoder) readListStart() byte {
	return d.read()
}

func (d *Decoder) readMapStart() byte {
	return d.read()
}

func (d *Decoder) IsEnd() bool {
	var code byte
	if d.offset < d.length {
		code = d.buffer[d.offset]
	} else {
		code = d.read()
		if code != 0xff {
			d.offset--
		}
	}
	return code == 0xff || code == 'Z'
}

func (d *Decoder) readEnd() {
	code := d.readTag()
	if code == 'Z' {
		return
	}
	if code == 0xff {
		panic(d.error("unexpected end of file"))
	}
	panic(d.error("unknown code:" + d.codeName(code)))
}

func (d *Decoder) readMapEnd() {
	code := d.readTag()
	if code != 'Z' {
		panic(d.error("expected end of map ('Z') at '" + d.codeName(code) + "'"))
	}
}

func (d *Decoder) readListEnd() {
	code := d.readTag()
	if code != 'Z' {
		panic(d.error("expected end of list ('Z') at '" + d.codeName(code) + "'"))
	}
}

func (d *Decoder) AddRef(ref interface{}) int {
	d.refs = append(d.refs, ref)
	return len(d.refs) - 1
}

func (d *Decoder) SetRef(idx int, ref interface{}) {
	d.refs[idx] = ref
}

func (d *Decoder) ResetRef() {
	d.refs = d.refs[0:0]
}

func (d *Decoder) Reset() {
	d.ResetRef()
	d.classDefs = d.classDefs[0:0]
	d.types = d.types[0:0]
}

func (d *Decoder) ResetBuffer() {
	offset := d.offset
	d.offset = 0

	length := d.length
	d.length = 0

	if length > 0 && offset != length {
		panic(fmt.Errorf("offset=%d length=%d", offset, length))
	}
}

func (d *Decoder) readType() string {
	code := d.readTag()
	d.offset--

	switch code {
	case 0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,

		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f,

		0x30, 0x31, 0x32, 0x33, BC_STRING_CHUNK, 'S':
		{
			typ := d.readString()

			d.types = append(d.types, *typ)

			return *typ
		}

	default:
		{
			ref := d.readInt()

			if len(d.types) <= ref {
				panic(fmt.Errorf("type ref #%d is greater than the number of valid types (%d)", ref, len(d.types)))
			}

			return d.types[ref]
		}
	}
}

func (d *Decoder) parseInt() int {
	offset := d.offset

	if offset+3 < d.length {
		buffer := d.buffer

		d.offset = offset + 4

		return int(binary.BigEndian.Uint32(buffer[offset+0 : offset+4]))
	} else {

		return int(binary.BigEndian.Uint32([]byte{d.read(), d.read(), d.read(), d.read()}))
	}
}

func (d *Decoder) parseLong() int64 {
	return int64(binary.BigEndian.Uint64([]byte{d.read(), d.read(), d.read(), d.read(), d.read(), d.read(), d.read(), d.read()}))
}

func (d *Decoder) parseDouble() float64 {
	bits := d.parseLong()
	return math.Float64frombits(uint64(bits))
}

func (d *Decoder) parseString(builder *strings.Builder) {
	for {
		if d.chunkLength <= 0 {
			if !d.parseChunkLength() {
				return
			}
		}

		length := d.chunkLength
		d.chunkLength = 0
		for length > 0 {
			length--
			builder.WriteRune(d.parseUTF8Char())
		}
	}
}

func (d *Decoder) parseChar() rune {
	for d.chunkLength <= 0 {
		if !d.parseChunkLength() {
			return 0xff
		}
	}
	d.chunkLength--
	return d.parseUTF8Char()
}

func (d *Decoder) parseChunkLength() bool {
	if d.isLastChunk {
		return false
	}

	code := d.readTag()
	switch code {
	case BC_STRING_CHUNK:
		d.isLastChunk = false
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case 'S':
		d.isLastChunk = true
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case 0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,

		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f:
		d.isLastChunk = true
		d.chunkLength = int(code)

	case 0x30, 0x31, 0x32, 0x33:
		d.isLastChunk = true
		d.chunkLength = (int(code)-0x30)*256 + int(d.read())

	default:
		panic(d.expect("string", code))
	}

	return true
}

func (d *Decoder) parseUTF8Char() rune {
	ch := d.readTag()
	if ch < 0x80 {
		return rune(ch)
	} else if (ch & 0xe0) == 0xc0 {
		ch1 := d.read()
		v := ((ch & 0x1f) << 6) + (ch1 & 0x3f)

		return rune(v)
	} else if (ch & 0xf0) == 0xe0 {
		ch1 := d.read()
		ch2 := d.read()
		v := (int(ch&0x0f) << 12) + (int(ch1&0x3f) << 6) + int(ch2&0x3f)

		return rune(v)
	} else {
		panic(d.error("bad utf-8 encoding at " + d.codeName(ch)))
	}
}

func (d *Decoder) parseByte() byte {
	for d.chunkLength <= 0 {
		if d.isLastChunk {
			return 0xff
		}

		code := d.read()

		switch code {
		case BC_BINARY_CHUNK:
			d.isLastChunk = false
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case 'B':
			d.isLastChunk = true
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case 0x20, 0x21, 0x22, 0x23,
			0x24, 0x25, 0x26, 0x27,
			0x28, 0x29, 0x2a, 0x2b,
			0x2c, 0x2d, 0x2e, 0x2f:
			d.isLastChunk = true
			d.chunkLength = int(code) - 0x20

		case 0x34, 0x35, 0x36, 0x37:
			d.isLastChunk = true
			d.chunkLength = (int(code)-0x34)*256 + int(d.read())

		default:
			panic(d.expect("byte[]", code))
		}
	}

	d.chunkLength--

	return d.read()
}

func (d *Decoder) read2(buffer []byte, offset int, length int) int {
	readLength := 0
	for length > 0 {
		for d.chunkLength <= 0 {
			if d.isLastChunk {
				if readLength == 0 {
					return 0xff
				}
				return readLength
			}

			code := d.read()

			switch code {
			case BC_BINARY_CHUNK:
				d.isLastChunk = false
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case BC_BINARY:
				d.isLastChunk = true
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case 0x20, 0x21, 0x22, 0x23,
				0x24, 0x25, 0x26, 0x27,
				0x28, 0x29, 0x2a, 0x2b,
				0x2c, 0x2d, 0x2e, 0x2f:
				d.isLastChunk = true
				d.chunkLength = int(code) - 0x20

			case 0x34, 0x35, 0x36, 0x37:
				d.isLastChunk = true
				d.chunkLength = (int(code)-0x34)*256 + int(d.read())

			default:
				panic(d.expect("byte[]", code))
			}
		}

		sublen := d.chunkLength
		if length < sublen {
			sublen = length
		}

		if d.length <= d.offset && !d.readBuffer() {
			return -1
		}

		if d.length-d.offset < sublen {
			sublen = d.length - d.offset
		}

		copy(buffer[offset:], d.buffer[d.offset:d.offset+sublen])

		d.offset += sublen

		offset += sublen
		readLength += sublen
		length -= sublen
		d.chunkLength -= sublen
	}

	return readLength
}

func (d *Decoder) unread() {
	if d.offset <= 0 {
		panic(fmt.Errorf("illegal state"))
	}
	d.offset--
}

func (d *Decoder) codeName(ch byte) string {
	if ch == 0xff {
		return "end of file"
	}
	return fmt.Sprintf("0x%s (%v)", strings.ToUpper(hex.EncodeToString([]byte{ch})), rune(ch))
}

func (d *Decoder) expect(expect string, ch byte) error {
	if ch == 0xff {
		return d.error(fmt.Sprintf("expected %s at end of file", expect))
	}
	return d.error(fmt.Sprintf("expected %s(0x%s) at end of file", expect, strings.ToUpper(hex.EncodeToString([]byte{ch}))))
}

func (d *Decoder) error(message string) error {
	if d.method != nil {
		message = *d.method + ": " + message
	}
	return fmt.Errorf(message)
}

type ObjectDefinition struct {
	typ        string
	vc         *VirtualClass
	fieldNames []string
}

func NewObjectDefinition(typ string, fieldNames []string) *ObjectDefinition {
	return &ObjectDefinition{
		typ:        typ,
		vc:         NewVirtualClass(typ, fieldNames),
		fieldNames: fieldNames,
	}
}
