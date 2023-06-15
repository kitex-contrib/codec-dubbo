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
	ReadSize = 1024 // reader buffer size
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
		return CHAR_END_MARK
	}
	r := d.buffer[d.offset]
	d.offset++
	return r
}

// readBuffer Read data to buffer from input stream
func (d *Decoder) readBuffer() bool {
	l, err := d.in.Read(d.buffer)
	d.offset = 0
	if l <= 0 || nil != err {
		d.length = 0
		return false
	}
	d.length = l
	return true
}

// readNext Read next tag mark
func (d *Decoder) readNext() byte {
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
	if tag != BC_NULL { // null tag
		panic(d.expect("null", tag))
	}
}

// readBoolean Try to read a boolean value
func (d *Decoder) readBoolean() bool {
	tag := d.readNext()
	switch {
	case tag == BC_TRUE: // boolean true tag
		return true

	case tag == BC_FALSE: // boolean false tag
		return false

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		return tag != BC_INT_ZERO

	// INT_BYTE = 0
	case tag == BC_INT_BYTE_ZERO:
		return d.read() != 0

	// INT_BYTE != 0
	case (tag >= BYTE_INT && tag <= BYTE_INT_MAX) ||
		(tag >= BYTE_INT_LIMIT && tag <= BYTE_INT_LIMIT_MAX):
		d.read()
		return true

	// INT_SHORT = 0
	case tag == BC_INT_SHORT_ZERO:
		return (int16(d.read())<<8 + int16(d.read())) != 0

	// INT_SHORT != 0
	case (tag >= SHORT_INT && tag <= SHORT_INT_MAX) ||
		(tag >= SHORT_INT_LIMIT && tag <= SHORT_INT_LIMIT_MAX):
		d.read()
		d.read()
		return true

	case tag == BC_INT:
		return d.parseInt() != 0

	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		return tag != BC_LONG_ZERO

	// LONG_BYTE = 0
	case tag == BC_LONG_BYTE_ZERO:
		return d.read() != 0

	// LONG_BYTE != 0
	case (tag >= BYTE_LONG && tag <= BYTE_LONG_MAX) ||
		(tag >= BYTE_LONG_LIMIT):
		d.read()
		return true

	// INT_SHORT = 0
	case tag == BC_LONG_SHORT_ZERO:
		return (int16(d.read())<<8 + int16(d.read())) != 0

	// INT_SHORT != 0
	case (tag >= SHORT_LONG && tag <= SHORT_LONG_MAX) ||
		(tag >= SHORT_LONG_LIMIT && tag <= SHORT_LONG_LIMIT_MAX):
		d.read()
		d.read()
		return true

	case tag == BC_LONG_INT:
		v := 0x1000000*int64(d.read()) + 0x10000*int64(d.read()) + 0x100*int64(d.read()) + int64(d.read())
		return v != 0

	case tag == BC_LONG:
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

	case tag == BC_DOUBLE:
		return d.parseDouble() != 0.0

	case tag == BC_NULL:
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
	case tag == BC_NULL:
		return 0

	case tag == BC_FALSE:
		return 0

	case tag == BC_TRUE:
		return 1

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		return int(tag) - BC_INT_ZERO

	// byte int
	case tag >= BYTE_INT && tag <= BYTE_INT_LIMIT_MAX:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

	// short int
	case tag >= SHORT_INT && tag <= SHORT_INT_LIMIT_MAX:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + int(d.read())<<8 + int(d.read())

	// 32bits int
	case tag == BC_INT || tag == BC_LONG_INT:
		return (int(d.read()) << 24) + (int(d.read()) << 16) + (int(d.read()) << 8) + int(d.read())

	// direct long
	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		return int(tag) - BC_LONG_ZERO

	// byte long
	case tag >= BYTE_LONG:
		return ((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read())

	// short long
	case tag >= SHORT_LONG && tag <= SHORT_LONG_LIMIT_MAX:
		return ((int(tag) - BC_LONG_SHORT_ZERO) << 16) + int(d.read())<<8 + int(d.read())

	// long
	case tag == BC_LONG:
		return int(d.parseLong())

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	// case LONG_BYTE:
	case tag == BC_DOUBLE_BYTE:
		return int(d.readNext())

	// case INT_SHORT:
	// case LONG_SHORT:
	case tag == BC_DOUBLE_SHORT:
		return int(int16(d.read())<<8 + int16(d.read()))

	case tag == BC_DOUBLE_MILL:
		return d.parseInt() / 1000

	case tag == BC_DOUBLE:
		return int(d.parseDouble())

	default:
	}
	panic(d.expect("integer", tag))
}

// readLong Try to read a int64 value
func (d *Decoder) readLong() int64 {
	tag := d.read()
	switch {
	case tag == BC_NULL:
		return 0

	case tag == BC_FALSE:
		return 0

	case tag == BC_TRUE:
		return 1

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		return int64(tag) - BC_INT_ZERO

	// byte int
	case tag >= BYTE_INT && tag <= BYTE_INT_LIMIT_MAX:
		return ((int64(tag) - BC_INT_BYTE_ZERO) << 8) + int64(d.read())

	// short int
	case tag >= SHORT_INT && tag <= SHORT_INT_LIMIT_MAX:
		return ((int64(tag) - BC_INT_SHORT_ZERO) << 16) + int64(d.read())<<8 + int64(d.read())

	// case LONG_BYTE:
	case tag == BC_DOUBLE_BYTE:
		return int64(d.readNext())

	// case INT_SHORT:
	// case LONG_SHORT:
	case tag == BC_DOUBLE_SHORT:
		return int64(d.read())<<8 + int64(d.read())

	// 32bits int
	case tag == BC_INT || tag == BC_LONG_INT:
		return int64(d.parseInt())

	// direct long
	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		return int64(tag) - BC_LONG_ZERO

	// byte long
	case tag >= BYTE_LONG:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

	// short long
	case tag >= SHORT_LONG && tag <= SHORT_LONG_LIMIT_MAX:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + int64(d.read())<<8 + int64(d.read())

	// long
	case tag == BC_LONG:
		return d.parseLong()

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	case tag == BC_DOUBLE_MILL:
		return int64(d.parseInt() / 1000)

	case tag == BC_DOUBLE:
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
	case tag == BC_NULL:
		return 0

	case tag == BC_FALSE:
		return 0

	case tag == BC_TRUE:
		return 1

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		return float64(tag) - BC_INT_ZERO

	// byte int
	case tag >= BYTE_INT && tag <= BYTE_INT_LIMIT_MAX:
		return float64((int64(tag)-BC_INT_BYTE_ZERO)<<8) + float64(d.read())

	// short int
	case tag >= SHORT_INT && tag <= SHORT_INT_LIMIT_MAX:
		return float64((int64(tag)-BC_INT_SHORT_ZERO)<<16) + float64(int64(d.read())<<8) + float64(d.read())

	// 32bits int
	case tag == BC_INT || tag == BC_LONG_INT:
		return float64(d.parseInt())

	// direct long
	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		return float64(tag) - BC_LONG_ZERO

	// byte long
	case tag >= BYTE_LONG:
		return float64((int64(tag)-BC_LONG_BYTE_ZERO)<<8) + float64(d.read())

	// short long
	case tag >= SHORT_LONG && tag <= SHORT_LONG_LIMIT_MAX:
		return float64((int64(tag)-BC_LONG_SHORT_ZERO)<<16) + float64(int64(d.read())<<8) + float64(d.read())

	// long
	case tag == BC_LONG:
		return float64(d.parseLong())

	case tag == BC_DOUBLE_ZERO:
		return 0

	case tag == BC_DOUBLE_ONE:
		return 1

	case tag == BC_DOUBLE_BYTE:
		return float64(d.readNext())

	case tag == BC_DOUBLE_SHORT:
		return float64(int64(d.read())<<8) + float64(d.read())

	case tag == BC_DOUBLE_MILL:
		return float64(d.parseInt() / 1000)

	case tag == BC_DOUBLE:
		return d.parseDouble()

	default:
	}
	panic(d.expect("double", tag))
}

// readUTCDate Try to read a timestamp value
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

// readString Try to read a String value
func (d *Decoder) readString() *string {
	tag := d.read()
	var val string
	switch {
	case tag == BC_NULL:
		return nil

	case tag == BC_TRUE:
		val = "true"
		return &val

	case tag == BC_FALSE:
		val = "false"
		return &val

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		val = strconv.Itoa(int(tag) - BC_INT_ZERO)
		return &val

	// byte int
	case tag >= BYTE_INT && tag <= BYTE_INT_LIMIT_MAX:
		val = strconv.Itoa(((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read()))
		return &val

	// short int
	case tag >= SHORT_INT && tag <= SHORT_INT_LIMIT_MAX:
		val = strconv.Itoa(((int(tag) - BC_INT_SHORT_ZERO) << 16) + int(d.read())<<8 + int(d.read()))
		return &val

	// 32bits int
	case tag == BC_INT || tag == BC_LONG_INT:
		val = strconv.Itoa(d.parseInt())
		return &val

	// direct long
	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		val = strconv.Itoa(int(tag) - BC_LONG_ZERO)
		return &val

	// byte long
	case tag >= BYTE_LONG:
		val = strconv.Itoa(((int(tag) - BC_LONG_BYTE_ZERO) << 8) + int(d.read()))
		return &val

	// short long
	case tag >= SHORT_LONG && tag <= SHORT_LONG_LIMIT_MAX:
		val = strconv.Itoa(((int(tag) - BC_LONG_SHORT_ZERO) << 16) + int(d.read())<<8 + int(d.read()))
		return &val

	// long
	case tag == BC_LONG:
		val = fmt.Sprintf("%d", d.parseLong())
		return &val

	case tag == BC_DOUBLE_ZERO:
		val = "0.0"
		return &val

	case tag == BC_DOUBLE_ONE:
		val = "1.0"
		return &val

	case tag == BC_DOUBLE_BYTE:
		val = strconv.Itoa(int(d.readNext()))
		return &val

	case tag == BC_DOUBLE_SHORT:
		val = strconv.Itoa(int(int16(int(d.read())<<8 + int(d.read()))))
		return &val

	case tag == BC_DOUBLE_MILL:
		{
			mills := d.parseInt()
			val = strconv.Itoa(mills / 1000)
			return &val
		}

	case tag == BC_DOUBLE:
		val = fmt.Sprintf("%v", d.parseDouble())
		return &val

	case tag == BC_STRING || tag == BC_STRING_CHUNK:
		d.isLastChunk = tag == BC_STRING
		d.chunkLength = (int(d.read()) << 8) + int(d.read())
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == CHAR_END_MARK {
				break
			}
			d.sbuf.WriteRune(ch)
		}
		val = d.sbuf.String()
		return &val

	// 0-byte string
	case tag <= STRING_DIRECT_MAX:
		d.isLastChunk = true
		d.chunkLength = int(tag)
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == CHAR_END_MARK {
				break
			}
			d.sbuf.WriteRune(ch)
		}

		val = d.sbuf.String()
		return &val

	case tag >= BC_STRING_SHORT && tag <= BC_STRING_SHORT_MAX:
		d.isLastChunk = true
		d.chunkLength = (int(tag)-BC_STRING_SHORT)<<8 + int(d.read())
		d.sbuf.Reset()
		var ch rune
		for {
			ch = d.parseChar()
			if ch == CHAR_END_MARK {
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

// ReadObject Decode Hessian2 data
func (d *Decoder) ReadObject() interface{} {
	tag := d.readNext()
	switch {
	case tag == BC_NULL:
		return nil

	case tag == BC_TRUE:
		return true

	case tag == BC_FALSE:
		return false

	// direct integer
	case tag >= DIRECT_INT && tag <= DIRECT_INT_MAX:
		return int(tag) - BC_INT_ZERO

	// byte int
	case tag >= BYTE_INT && tag <= BYTE_INT_LIMIT_MAX:
		return ((int(tag) - BC_INT_BYTE_ZERO) << 8) + int(d.read())

	// short int
	case tag >= SHORT_INT && tag <= SHORT_INT_LIMIT_MAX:
		return ((int(tag) - BC_INT_SHORT_ZERO) << 16) + int(d.read())<<8 + int(d.read())

	// 32bits int
	case tag == BC_INT:
		return d.parseInt()

	// direct long
	case tag >= DIRECT_LONG && tag <= DIRECT_LONG_MAX:
		return int64(tag) - BC_LONG_ZERO

	// byte long
	case tag >= BYTE_LONG:
		return ((int64(tag) - BC_LONG_BYTE_ZERO) << 8) + int64(d.read())

	// short long
	case tag >= SHORT_LONG && tag <= SHORT_LONG_LIMIT_MAX:
		return ((int64(tag) - BC_LONG_SHORT_ZERO) << 16) + int64(d.read())<<8 + int64(d.read())

	// 32bits int
	case tag == BC_LONG_INT:
		return int64(d.parseInt())

	// long
	case tag == BC_LONG:
		return d.parseLong()

	case tag == BC_DOUBLE_ZERO:
		return 0.0

	case tag == BC_DOUBLE_ONE:
		return 1.0

	case tag == BC_DOUBLE_BYTE:
		return float64(d.read())

	case tag == BC_DOUBLE_SHORT:
		return float64(int16(d.read())<<8 + int16(d.read()))

	case tag == BC_DOUBLE_MILL:
		{
			mills := float64(d.parseInt())

			return 0.001 * mills
		}

	case tag == BC_DOUBLE:
		return d.parseDouble()

	case tag == BC_DATE:
		l64 := d.parseLong()
		return time.Unix(l64/1000, l64%1000*10e5)

	case tag == BC_DATE_MINUTE:
		return time.Unix(int64(d.parseInt())*60000, 0)

	case tag == BC_STRING_CHUNK || tag == BC_STRING:
		{
			d.isLastChunk = tag == BC_STRING
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case tag <= STRING_DIRECT_MAX:
		{
			d.isLastChunk = true
			d.chunkLength = int(tag)

			d.sbuf.Reset()

			builder := d.sbuf
			d.parseString(d.sbuf)

			return builder.String()
		}

	case tag >= BC_STRING_SHORT && tag <= BC_STRING_SHORT_MAX:
		{
			d.isLastChunk = true
			d.chunkLength = (int(tag)-BC_STRING_SHORT)<<8 + int(d.read())

			d.sbuf.Reset()

			d.parseString(d.sbuf)

			return d.sbuf.String()
		}

	case tag == BC_BINARY_CHUNK || tag == BC_BINARY:
		{
			d.isLastChunk = tag == BC_BINARY
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

			bos := bytes.NewBuffer([]byte{})

			var data byte
			for {
				data = d.parseByte()
				if data == CHAR_END_MARK {
					break
				}
				bos.WriteByte(data)
			}

			return bos.Bytes()
		}

	case tag >= BC_BINARY_DIRECT && tag <= INT_DIRECT_MAX:
		{
			d.isLastChunk = true
			length := int(tag) - BC_BINARY_DIRECT
			d.chunkLength = 0

			data := make([]byte, length)

			for i := 0; i < length; i++ {
				data[i] = d.read()
			}

			return data
		}

	case tag >= BC_BINARY_SHORT && tag <= BC_BINARY_SHORT_MAX:
		{
			d.isLastChunk = true
			length := (int(tag)-BC_BINARY_SHORT)<<8 + int(d.read())
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
	case tag >= P_PACKET_SHORT && tag <= P_PACKET_SHORT_MAX:
		{
			typ := d.readType()
			length := int(tag) - P_PACKET_SHORT
			return d.readList(length, &typ)
		}

	// compact fixed untyped list
	case tag >= BC_LIST_DIRECT_UNTYPED && tag <= PACKET_DIRECT_MAX:
		{
			length := int(tag) - BC_LIST_DIRECT_UNTYPED
			return d.readList(length, nil)
		}

	case tag == BC_MAP_UNTYPED:
		{
			return d.readMap(nil)
		}

	case tag == BC_MAP:
		{
			typ := d.readType()
			return d.readMap(&typ)
		}

	case tag == BC_OBJECT_DEF:
		{
			d.readObjectDefinition(nil)

			return d.ReadObject()
		}

	case tag >= BC_OBJECT_DIRECT && tag <= BC_OBJECT_DIRECT_MAX:
		{
			ref := int(tag) - BC_OBJECT_DIRECT

			if len(d.classDefs) <= ref {
				panic(d.error(fmt.Sprintf("No classes defined at reference '%s'", strings.ToUpper(hex.EncodeToString([]byte{tag})))))
			}

			def := d.classDefs[ref]

			return d.readObjectInstance(nil, def)
		}

	case tag == BC_OBJECT:
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
		if tag == CHAR_END_MARK {
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
		if code != CHAR_END_MARK {
			d.offset--
		}
	}
	return code == CHAR_END_MARK || code == BC_END
}

// readEnd Check data read is end
func (d *Decoder) readEnd() {
	code := d.readNext()
	if code == BC_END {
		return
	}
	if code == CHAR_END_MARK {
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
	code := d.readNext()
	d.offset--

	switch {
	case (code <= STRING_DIRECT_MAX) ||
		(code >= BC_STRING_SHORT && code <= BC_STRING_SHORT_MAX) ||
		code == BC_STRING_CHUNK ||
		code == BC_STRING:
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
			return CHAR_END_MARK
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

	code := d.readNext()
	switch {
	case code == BC_STRING_CHUNK:
		d.isLastChunk = false
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case code == BC_STRING:
		d.isLastChunk = true
		d.chunkLength = (int(d.read()) << 8) + int(d.read())

	case code <= STRING_DIRECT_MAX:
		d.isLastChunk = true
		d.chunkLength = int(code)

	case code >= BC_STRING_SHORT && code <= BC_STRING_SHORT_MAX:
		d.isLastChunk = true
		d.chunkLength = (int(code)-BC_STRING_SHORT)<<8 + int(d.read())

	default:
		panic(d.expect("string", code))
	}

	return true
}

// parseUTF8Char Convert buffer to rune value
func (d *Decoder) parseUTF8Char() rune {
	ch := d.readNext()
	if ch <= ASCII_CODE_MAX {
		return rune(ch)
	} else if (ch & BC_LONG_ZERO) == BYTE_INT {
		ch1 := d.read()
		v := (int(ch&STRING_DIRECT_MAX) << 6) + int(ch1&SHORT_LONG_LIMIT_MAX)

		return rune(v)
	} else if (ch & BYTE_LONG) == BC_LONG_ZERO {
		ch1 := d.read()
		ch2 := d.read()
		v := (int(ch&BINARY_DIRECT_MAX) << 12) + (int(ch1&SHORT_LONG_LIMIT_MAX) << 6) + int(ch2&SHORT_LONG_LIMIT_MAX)

		return rune(v)
	} else {
		panic(d.error("bad utf-8 encoding at " + d.codeName(ch)))
	}
}

// parseByte Convert buffer to byte value
func (d *Decoder) parseByte() byte {
	for d.chunkLength <= 0 {
		if d.isLastChunk {
			return CHAR_END_MARK
		}

		code := d.read()

		switch {
		case code == BC_BINARY_CHUNK: // 8-bit binary data non-final chunk ('A')
			d.isLastChunk = false
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case code == BC_BINARY: // 8-bit binary data final chunk ('B')
			d.isLastChunk = true
			d.chunkLength = (int(d.read()) << 8) + int(d.read())

		case code >= BC_BINARY_DIRECT && code <= INT_DIRECT_MAX: // binary data length 0-16
			d.isLastChunk = true
			d.chunkLength = int(code) - BC_BINARY_DIRECT

		case code >= BC_BINARY_SHORT && code <= BC_BINARY_SHORT_MAX: // binary data length 0-1023
			d.isLastChunk = true
			d.chunkLength = (int(code)-BC_BINARY_SHORT)<<8 + int(d.read())

		default:
			panic(d.expect("byte[]", code))
		}
	}

	d.chunkLength--

	return d.read()
}

// codeName Convert byte to hexadecimal format
func (d *Decoder) codeName(ch byte) string {
	if ch == CHAR_END_MARK {
		return "end of file"
	}
	return fmt.Sprintf("0x%s (%v)", strings.ToUpper(hex.EncodeToString([]byte{ch})), rune(ch))
}

// expect Get a expect information error message
func (d *Decoder) expect(expect string, ch byte) error {
	if ch == CHAR_END_MARK {
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
