/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
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

// NewDecoder Get Hessian2 decoder instance with a reader
func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{
		in:     in,
		buffer: make([]byte, ReadSize),
		sbuf:   &strings.Builder{},
	}
}

// NewDecoderWithByteArray Get Hessian2 decoder instance with a byte array buffer
func NewDecoderWithByteArray(buf []byte) *Decoder {
	return NewDecoder(bytes.NewBuffer(buf))
}

type Decoder struct {
	in io.Reader // input stream

	length int // buffer length
	offset int // read index

	buffer []byte // data buffer

	chunkLength int  // read String value size
	isLastChunk bool // read String end mark

	refs      []interface{}       // temporary data reference
	types     []string            // temporary data type/ data class
	classDefs []*ObjectDefinition // temporary data model

	sbuf *strings.Builder // String value cache buffer
}

// read Get next byte
func (d *Decoder) read() byte {
	if d.length <= d.offset && !d.readBuffer() {
		return 0xFF
	}
	r := d.buffer[d.offset]
	d.offset++
	return r
}

// readBuffer Read data to buffer from input stream
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

// readTag Read next tag mark
func (d *Decoder) readTag() byte {
	if d.offset < d.length {
		b := d.buffer[d.offset]
		d.offset++
		return b
	} else {
		return d.read()
	}
}

// readNull Try to read a null value
func (d *Decoder) readNull() {
	tag := d.read()
	if tag != 'N' {
		panic(d.expect("null", tag))
	}
}

// readBoolean Try to read a boolean value
func (d *Decoder) readBoolean() bool {
	tag := d.readTag()
	switch {
	case tag == 'T':
		return true

	case tag == 'F':
		return false

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		return tag != BC_INT_ZERO

	// INT_BYTE = 0
	case tag == 0xc8:
		return d.read() != 0

	// INT_BYTE != 0
	case (tag >= 0xc0 && tag <= 0xc7) ||
		(tag >= 0xc9 && tag <= 0xcf):
		d.read()
		return true

	// INT_SHORT = 0
	case tag == 0xd4:
		return (256*int16(d.read()) + int16(d.read())) != 0

	// INT_SHORT != 0
	case (tag >= 0xd0 && tag <= 0xd3) ||
		(tag >= 0xd5 && tag <= 0xd7):
		d.read()
		d.read()
		return true

	case tag == 'I':
		return d.parseInt() != 0

	case tag >= 0xd8 && tag <= 0xef:
		return tag != BC_LONG_ZERO

	// LONG_BYTE = 0
	case tag == 0xf8:
		return d.read() != 0

	// LONG_BYTE != 0
	case (tag >= 0xf0 && tag <= 0xf7) ||
		(tag >= 0xf9):
		d.read()
		return true

	// INT_SHORT = 0
	case tag == 0x3c:
		return (256*int16(d.read()) + int16(d.read())) != 0

	// INT_SHORT != 0
	case (tag >= 0x38 && tag <= 0x3b) ||
		(tag >= 0x3d && tag <= 0x3f):
		d.read()
		d.read()
		return true

	case tag == BC_LONG_INT:
		v := 0x1000000*int64(d.read()) + 0x10000*int64(d.read()) + 0x100*int64(d.read()) + int64(d.read())
		return v != 0

	case tag == 'L':
		return d.parseLong() != 0

	case tag == BC_DOUBLE_ZERO:
		return false

	case tag == BC_DOUBLE_ONE:
		return true

	case tag == BC_DOUBLE_BYTE:
		return d.read() != 0

	case tag == BC_DOUBLE_SHORT:
		return (0x100*int16(d.read()) + int16(d.read())) != 0

	case tag == BC_DOUBLE_MILL:
		return d.parseInt() != 0

	case tag == 'D':
		return d.parseDouble() != 0.0

	case tag == 'N':
		return false

	default:
		panic(d.expect("boolean", tag))
	}
}

// readShort Try to read a int16 value
func (d *Decoder) readShort() int16 {
	return int16(d.readInt())
}

// readInt Try to read a int value
func (d *Decoder) readInt() int {
	tag := d.read()
	switch {
	case tag == 'N':
		return 0

	case tag == 'F':
		return 0

	case tag == 'T':
		return 1

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		return int(tag) - BC_INT_ZERO

	/* byte int */
	case tag >= 0xc0 && tag <= 0xcf:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

	/* short int */
	case tag >= 0xd0 && tag <= 0xd7:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case tag == 'I' || tag == BC_LONG_INT:
		return (int(d.read()) << 24) + (int(d.read()) << 16) + (int(d.read()) << 8) + int(d.read())

	// direct long
	case tag >= 0xd8 && tag <= 0xef:
		return int(tag) - BC_LONG_ZERO

	/* byte long */
	case tag >= 0xf0:
		return ((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read())

	/* short long */
	case tag >= 0x38 && tag <= 0x3f:
		return ((int(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case tag == 'L':
		return int(d.parseLong())

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	//case LONG_BYTE:
	case tag == BC_DOUBLE_BYTE:
		if d.offset < d.length {
			v := int(d.buffer[d.offset])
			d.offset++
			return v
		} else {
			return int(d.read())
		}

	//case INT_SHORT:
	//case LONG_SHORT:
	case tag == BC_DOUBLE_SHORT:
		return int(256*int16(d.read()) + int16(d.read()))

	case tag == BC_DOUBLE_MILL:
		return d.parseInt() / 1000

	case tag == 'D':
		return int(d.parseDouble())

	default:
	}
	panic(d.expect("integer", tag))
}

// readLong Try to read a int64 value
func (d *Decoder) readLong() int64 {
	tag := d.read()
	switch {
	case tag == 'N':
		return 0

	case tag == 'F':
		return 0

	case tag == 'T':
		return 1

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		return int64(tag) - BC_INT_ZERO

	/* byte int */
	case tag >= 0xc0 && tag <= 0xcf:
		return ((int64(tag) - BC_INT_BYTE_ZERO) << 8) + int64(d.read())

	/* short int */
	case tag >= 0xd0 && tag <= 0xd7:
		return ((int64(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

	//case LONG_BYTE:
	case tag == BC_DOUBLE_BYTE:
		if d.offset < d.length {
			v := int64(d.buffer[d.offset])
			d.offset++
			return v
		} else {
			return int64(d.read())
		}

	//case INT_SHORT:
	//case LONG_SHORT:
	case tag == BC_DOUBLE_SHORT:
		return 256*int64(d.read()) + int64(d.read())

	case tag == 'I' || tag == BC_LONG_INT:
		return int64(d.parseInt())

	// direct long
	case tag >= 0xd8 && tag <= 0xef:
		return int64(tag) - BC_LONG_ZERO

	/* byte long */
	case tag >= 0xf0:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

	/* short long */
	case tag >= 0x38 && tag <= 0x3f:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

	case tag == 'L':
		return d.parseLong()

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	case tag == BC_DOUBLE_MILL:
		return int64(d.parseInt() / 1000)

	case tag == 'D':
		return int64(d.parseDouble())

	default:
	}
	panic(d.expect("long", tag))
}

// readFloat Try to read a float32 value
func (d *Decoder) readFloat() float32 {
	return float32(d.readDouble())
}

// readDouble Try to read a float64 value
func (d *Decoder) readDouble() float64 {
	tag := d.read()
	switch {
	case tag == 'N':
		return 0

	case tag == 'F':
		return 0

	case tag == 'T':
		return 1

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		return float64(tag) - 0x90

	/* byte int */
	case tag >= 0xc0 && tag <= 0xcf:
		return float64((int64(tag)-BC_INT_BYTE_ZERO)<<8) + float64(d.read())

	/* short int */
	case tag >= 0xd0 && tag <= 0xd7:
		return float64((int64(tag)-BC_INT_SHORT_ZERO)<<16) + 256*float64(d.read()) + float64(d.read())

	case tag == 'I' || tag == BC_LONG_INT:
		return float64(d.parseInt())

	// direct long
	case tag >= 0xd8 && tag <= 0xef:
		return float64(tag) - BC_LONG_ZERO

	/* byte long */
	case tag >= 0xf0:
		return float64((int64(tag)-BC_LONG_BYTE_ZERO)<<8) + float64(d.read())

	/* short long */
	case tag >= 0x38 && tag <= 0x3f:
		return float64((int64(tag)-BC_LONG_SHORT_ZERO)<<16) + 256*float64(d.read()) + float64(d.read())

	case tag == 'L':
		return float64(d.parseLong())

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	case tag == BC_DOUBLE_BYTE:
		if d.offset < d.length {
			v := float64(d.buffer[d.offset])
			d.offset++
			return v
		} else {
			return float64(d.read())
		}

	case tag == BC_DOUBLE_SHORT:
		return 256*float64(d.read()) + float64(d.read())

	case tag == BC_DOUBLE_MILL:
		return float64(d.parseInt() / 1000)

	case tag == 'D':
		return d.parseDouble()

	default:
	}
	panic(d.expect("double", tag))
}

// readUTCDate Try to read a time.Time value
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

// readChar Try to read a rune value
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

// readStringToBuffer Try to read a String value to the buffer
func (d *Decoder) readStringToBuffer(buffer []rune, offset int, length int) int {
	if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return 0xff
	} else if d.chunkLength == 0 {
		tag := d.read()

		switch {
		case tag == 'N':
			return 0xff

		case tag == 'S' || tag == BC_STRING_CHUNK:
			d.isLastChunk = tag == 'S'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case tag <= 0x1f:
			d.isLastChunk = true
			d.chunkLength = int(tag)

		case tag >= 0x30 && tag <= 0x33:
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

			switch {
			case tag == 'S' || tag == BC_STRING_CHUNK:
				d.isLastChunk = tag == 'S'
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case tag <= 0x1f:
				d.isLastChunk = true
				d.chunkLength = int(tag)

			case tag >= 0x30 && tag <= 0x33:
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

// readString Try to read a String value
func (d *Decoder) readString() *string {
	tag := d.read()
	var val string
	switch {
	case tag == 'N':
		return nil

	case tag == 'T':
		val = "true"
		return &val

	case tag == 'F':
		val = "false"
		return &val

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		val = strconv.Itoa(int(tag) - 0x90)
		return &val

	/* byte int */
	case tag >= 0xc0 && tag <= 0xcf:
		val = strconv.Itoa(((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read()))
		return &val

	/* short int */
	case tag >= 0xd0 && tag <= 0xd7:
		val = strconv.Itoa(((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read()))
		return &val

	case tag == 'I' || tag == BC_LONG_INT:
		val = strconv.Itoa(d.parseInt())
		return &val

	// direct long
	case tag >= 0xd8 && tag <= 0xef:
		val = strconv.Itoa(int(tag) - BC_LONG_ZERO)
		return &val

	/* byte long */
	case tag >= 0xf0:
		val = strconv.Itoa(((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read()))
		return &val

	/* short long */
	case tag >= 0x38 && tag <= 0x3f:
		val = strconv.Itoa(((int(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read()))
		return &val

	case tag == 'L':
		val = fmt.Sprintf("%d", d.parseLong())
		return &val

	case tag == BC_DOUBLE_ZERO:
		val = "0.0"
		return &val

	case tag == BC_DOUBLE_ONE:
		val = "1.0"
		return &val

	case tag == BC_DOUBLE_BYTE:
		val = strconv.Itoa(int(d.readTag()))
		return &val

	case tag == BC_DOUBLE_SHORT:
		val = strconv.Itoa(int(int16(256*int(d.read()) + int(d.read()))))
		return &val

	case tag == BC_DOUBLE_MILL:
		{
			mills := d.parseInt()
			val = strconv.Itoa(mills / 1000)
			return &val
		}

	case tag == 'D':
		val = fmt.Sprintf("%v", d.parseDouble())
		return &val

	case tag == 'S' || tag == BC_STRING_CHUNK:
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
	case tag <= 0x1f:
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

	case tag >= 0x30 && tag <= 0x33:
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

// readBytes Try to read a byte array
func (d *Decoder) readBytes() []byte {
	tag := d.read()

	switch {
	case tag == 'N':
		return nil

	case tag == BC_BINARY || tag == BC_BINARY_CHUNK:
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

	case tag >= 0x20 && tag <= 0x2f:
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

	case tag >= 0x34 && tag <= 0x37:
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

// readByte Try to read a byte value
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

	switch {
	case tag == 'N':
		return 0xFF

	case tag == 'B' || tag == BC_BINARY_CHUNK:
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

	case tag >= 0x20 && tag <= 0x2f:
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

	case tag >= 0x34 && tag <= 0x37:
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

// readBytes2 Try to read a byte array to buffer
func (d *Decoder) readBytes2(buffer []byte, offset int, length int) int {
	if d.chunkLength == EndOfData {
		d.chunkLength = 0
		return -1
	} else if d.chunkLength == 0 {
		tag := d.read()

		switch {
		case tag == 'N':
			return -1

		case tag == 'B' || tag == BC_BINARY_CHUNK:
			d.isLastChunk = tag == 'B'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case tag >= 0x20 && tag <= 0x2f:

			d.isLastChunk = true
			d.chunkLength = int(tag) - 0x20

		case tag >= 0x34 && tag <= 0x37:
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

// ReadObject Decode Hessian2 data
func (d *Decoder) ReadObject() interface{} {
	tag := d.readTag()
	switch {
	case tag == 'N':
		return nil

	case tag == 'T':
		return true

	case tag == 'F':
		return false

	// direct integer
	case tag >= 0x80 && tag <= 0xbf:
		return int(tag) - BC_INT_ZERO

	/* byte int */
	case tag >= 0xc0 && tag <= 0xcf:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

	/* short int */
	case tag >= 0xd0 && tag <= 0xd7:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + 256*int(d.read()) + int(d.read())

	case tag == 'I':
		return d.parseInt()

	// direct long
	case tag >= 0xd8 && tag <= 0xef:
		return int64(tag) - BC_LONG_ZERO

	/* byte long */
	case tag >= 0xf0:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

	/* short long */
	case tag >= 0x38 && tag <= 0x3f:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + 256*int64(d.read()) + int64(d.read())

	case tag == BC_LONG_INT:
		return int64(d.parseInt())

	case tag == 'L':
		return d.parseLong()

	case tag == BC_DOUBLE_ZERO:
		return 0.0

	case tag == BC_DOUBLE_ONE:
		return 1.0

	case tag == BC_DOUBLE_BYTE:
		return float64(d.read())

	case tag == BC_DOUBLE_SHORT:
		return float64(256*int16(d.read()) + int16(d.read()))

	case tag == BC_DOUBLE_MILL:
		{
			mills := float64(d.parseInt())

			return 0.001 * mills
		}

	case tag == 'D':
		return d.parseDouble()

	case tag == BC_DATE:
		l64 := d.parseLong()
		return time.Unix(l64/1000, l64%1000*10e5)

	case tag == BC_DATE_MINUTE:
		return time.Unix(int64(d.parseInt())*60000, 0)

	case tag == BC_STRING_CHUNK || tag == 'S':
		{
			d.isLastChunk = tag == 'S'
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case tag <= 0x1f:
		{
			d.isLastChunk = true
			d.chunkLength = int(tag)

			d.sbuf.Reset()

			builder := d.sbuf
			d.parseString(d.sbuf)

			return builder.String()
		}

	case tag >= 0x30 && tag <= 0x33:
		{
			d.isLastChunk = true
			d.chunkLength = (int(tag)-0x30)*256 + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case tag == BC_BINARY_CHUNK || tag == 'B':
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

	case tag >= 0x20 && tag <= 0x2f:
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

	case tag >= 0x34 && tag <= 0x37:
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

	case tag == BC_LIST_VARIABLE:
		{
			typ := d.readType()
			return d.readList(-1, &typ)
		}

	case tag == BC_LIST_VARIABLE_UNTYPED:
		{
			return d.readList(-1, nil)
		}

	case tag == BC_LIST_FIXED:
		{
			typ := d.readType()
			length := d.readInt()
			return d.readList(length, &typ)
		}

	case tag == BC_LIST_FIXED_UNTYPED:
		{
			length := d.readInt()
			return d.readList(length, nil)
		}

	// compact fixed list
	case tag >= 0x70 && tag <= 0x77:
		{
			typ := d.readType()
			length := int(tag) - 0x70
			return d.readList(length, &typ)
		}

	// compact fixed untyped list
	case tag >= 0x78 && tag <= 0x7f:
		{
			length := int(tag) - 0x78
			return d.readList(length, nil)
		}

	case tag == 'H':
		{
			return d.readMap(nil)
		}

	case tag == 'M':
		{
			typ := d.readType()
			return d.readMap(&typ)
		}

	case tag == 'C':
		{
			d.readObjectDefinition(nil)

			return d.ReadObject()
		}

	case tag >= 0x60 && tag <= 0x6f:
		{
			ref := int(tag) - 0x60

			if len(d.classDefs) <= ref {
				panic(d.error(fmt.Sprintf("No classes defined at reference '%s'", strings.ToUpper(hex.EncodeToString([]byte{tag})))))
			}

			def := d.classDefs[ref]

			return d.readObjectInstance(nil, def)
		}

	case tag == 'O':
		{
			ref := d.readInt()

			if len(d.classDefs) <= ref {

				panic(d.error(fmt.Sprintf("Illegal object reference #%d", ref)))
			}

			def := d.classDefs[ref]

			return d.readObjectInstance(nil, def)
		}

	case tag == BC_REF:
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

// readList Try to read a array data
func (d *Decoder) readList(length int, typ *string) []interface{} {
	list := make([]interface{}, length)
	d.addRef(list)
	for i := range list {
		list[i] = d.ReadObject()
	}
	return list
}

// readMap Try to read a map data
func (d *Decoder) readMap(typ *string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	d.addRef(m)
	for !d.isEnd() {
		m[d.ReadObject()] = d.ReadObject()
	}
	d.readEnd()
	return m
}

// readObjectDefinition Get a temporary data model
func (d *Decoder) readObjectDefinition(i interface{}) {
	typ := d.readString()
	length := d.readInt()

	fieldNames := make([]string, length)

	for i := 0; i < length; i++ {
		name := d.readString()

		fieldNames[i] = *name
	}

	def := newObjectDefinition(*typ, fieldNames)
	d.classDefs = append(d.classDefs, def)
}

// readObjectInstance Get a virtual Java class
func (d *Decoder) readObjectInstance(cl interface{}, def *ObjectDefinition) interface{} {
	vc := NewVirtualClass(def.typ, def.fieldNames)
	for _, key := range def.fieldNames {
		v := d.ReadObject()
		vc.fields[key] = v
	}
	return vc
}

// isEnd Check array or map data read is end
func (d *Decoder) isEnd() bool {
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

// readEnd Check data read is end
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

// addRef Add temporary data reference
func (d *Decoder) addRef(ref interface{}) int {
	d.refs = append(d.refs, ref)
	return len(d.refs) - 1
}

// resetRef Remove all temporary data reference
func (d *Decoder) resetRef() {
	d.refs = d.refs[0:0]
}

// Reset Remove all cache data
func (d *Decoder) Reset() {
	d.resetRef()
	d.classDefs = d.classDefs[0:0]
	d.types = d.types[0:0]
}

// readType Get the data type
func (d *Decoder) readType() string {
	code := d.readTag()
	d.offset--

	switch {
	case (code <= 0x1f) ||
		(code >= 0x30 && code <= 0x33) ||
		code == BC_STRING_CHUNK ||
		code == 'S':
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

// parseInt Convert buffer to int value
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

// parseLong Convert buffer to int64 value
func (d *Decoder) parseLong() int64 {
	return int64(binary.BigEndian.Uint64([]byte{d.read(), d.read(), d.read(), d.read(), d.read(), d.read(), d.read(), d.read()}))
}

// parseDouble Convert buffer to float64 value
func (d *Decoder) parseDouble() float64 {
	bits := d.parseLong()
	return math.Float64frombits(uint64(bits))
}

// parseString Convert String buffer to String value
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

// parseChar Convert buffer to rune value
func (d *Decoder) parseChar() rune {
	for d.chunkLength <= 0 {
		if !d.parseChunkLength() {
			return 0xff
		}
	}
	d.chunkLength--
	return d.parseUTF8Char()
}

// parseChunkLength Check String value size
func (d *Decoder) parseChunkLength() bool {
	if d.isLastChunk {
		return false
	}

	code := d.readTag()
	switch {
	case code == BC_STRING_CHUNK:
		d.isLastChunk = false
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case code == 'S':
		d.isLastChunk = true
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case code <= 0x1f:
		d.isLastChunk = true
		d.chunkLength = int(code)

	case code >= 0x30 && code <= 0x33:
		d.isLastChunk = true
		d.chunkLength = (int(code)-0x30)*256 + int(d.read())

	default:
		panic(d.expect("string", code))
	}

	return true
}

// parseUTF8Char Convert buffer to rune value
func (d *Decoder) parseUTF8Char() rune {
	ch := d.readTag()
	if ch < 0x80 {
		return rune(ch)
	} else if (ch & 0xe0) == 0xc0 {
		ch1 := d.read()
		v := (int(ch&0x1f) << 6) + int(ch1&0x3f)

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

// parseByte Convert buffer to byte value
func (d *Decoder) parseByte() byte {
	for d.chunkLength <= 0 {
		if d.isLastChunk {
			return 0xff
		}

		code := d.read()

		switch {
		case code == BC_BINARY_CHUNK:
			d.isLastChunk = false
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case code == 'B':
			d.isLastChunk = true
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case code >= 0x20 && code <= 0x2f:
			d.isLastChunk = true
			d.chunkLength = int(code) - 0x20

		case code >= 0x34 && code <= 0x37:
			d.isLastChunk = true
			d.chunkLength = (int(code)-0x34)*256 + int(d.read())

		default:
			panic(d.expect("byte[]", code))
		}
	}

	d.chunkLength--

	return d.read()
}

// read2 Read data to buffer
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

			switch {
			case code == BC_BINARY_CHUNK:
				d.isLastChunk = false
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case code == BC_BINARY:
				d.isLastChunk = true
				d.chunkLength = (int(d.read()) << 8) + int(d.read())

			case code >= 0x20 && code <= 0x2f:
				d.isLastChunk = true
				d.chunkLength = int(code) - 0x20

			case code >= 0x34 && code <= 0x37:
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

// codeName Convert byte to hexadecimal format
func (d *Decoder) codeName(ch byte) string {
	if ch == 0xff {
		return "end of file"
	}
	return fmt.Sprintf("0x%s (%v)", strings.ToUpper(hex.EncodeToString([]byte{ch})), rune(ch))
}

// expect Get a expect information error message
func (d *Decoder) expect(expect string, ch byte) error {
	if ch == 0xff {
		return d.error(fmt.Sprintf("expected %s at end of file", expect))
	}
	return d.error(fmt.Sprintf("expected %s(0x%s) at end of file", expect, strings.ToUpper(hex.EncodeToString([]byte{ch}))))
}

// error Get a error message
func (d *Decoder) error(message string) error {
	return fmt.Errorf(message)
}

// ObjectDefinition temporary data model
type ObjectDefinition struct {
	typ        string
	vc         *VirtualClass
	fieldNames []string
}

// newObjectDefinition Get a temporary data model instance
func newObjectDefinition(typ string, fieldNames []string) *ObjectDefinition {
	return &ObjectDefinition{
		typ:        typ,
		vc:         NewVirtualClass(typ, fieldNames),
		fieldNames: fieldNames,
	}
}
