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
	"fmt"
	"math"
	"testing"
)

// TestDecodeJavaSimpleClass test decode java simple class
func TestDecodeJavaSimpleClass(t *testing.T) {
	byteArray := []byte{
		67, 10, 97, 46, 98, 46, 99, 46, 84, 101, 115, 116, 153, 4, 110, 97, 109, 101, 3, 97, 103, 101, 4, 97, 103, 101,
		50, 5, 119, 105, 103, 104, 116, 6, 119, 105, 103, 104, 116, 50, 4, 97, 103, 101, 51, 5, 98, 121, 116, 101, 115,
		4, 108, 105, 115, 116, 3, 109, 97, 112, 96, 9, 67, 104, 97, 110, 103, 101, 100, 101, 110, 154, 95, 0, 0, 44,
		236, 236, 84, 74, 0, 0, 1, 136, 35, 86, 112, 163, 35, 1, 0, 6, 113, 6, 77, 97, 105, 110, 36, 49, 96, 1, 116,
		73, 100, 99, 40, 143, 68, 65, 164, 19, 212, 224, 0, 0, 0, 76, 0, 0, 1, 136, 35, 86, 112, 163, 70, 74, 0, 0, 1,
		136, 35, 86, 112, 163, 35, 1, 0, 6, 112, 12, 97, 46, 98, 46, 99, 46, 84, 101, 115, 116, 36, 49, 77, 12, 97, 46,
		98, 46, 99, 46, 84, 101, 115, 116, 36, 50, 90, 77, 6, 77, 97, 105, 110, 36, 50, 1, 84, 96, 1, 116, 73, 100, 99,
		40, 143, 68, 65, 164, 19, 212, 224, 0, 0, 0, 76, 0, 0, 1, 136, 35, 86, 112, 163, 70, 74, 0, 0, 1, 136, 35, 86,
		112, 163, 35, 1, 0, 6, 112, 145, 77, 146, 90, 90,
	}

	decoder := NewDecoderWithByteArray(byteArray)
	obj := decoder.ReadObject()
	switch vc := obj.(type) {
	case *VirtualClass:
		if vc.JavaClassPackage() != "a.b.c" {
			t.Errorf("Java Package name decode error, expect: %s\n", "a.b.c")
			t.Fail()
			return
		}
		if vc.JavaClassName() != "Test" {
			t.Errorf("Java Class name decode error, expect: %s\n", "Test")
			t.Fail()
			return
		}
		t.Logf("%s\n", vc.String())
	default:
		t.Error("Decode error\n")
		t.Fail()
		return
	}
}

// TestDecodeString test decode string value
func TestDecodeString(t *testing.T) {
	byteArray := []byte{9, 67, 104, 97, 110, 103, 101, 100, 101, 110} // Changeden

	decoder := NewDecoderWithByteArray(byteArray)
	str := decoder.readString()
	if nil == str {
		t.Errorf("String value decode error, expect: not nil")
		t.Fail()
		return
	}
	if *str != "Changeden" {
		t.Errorf("String value decode error, expect: %s\n", "Changeden")
		t.Fail()
		return
	}
	t.Logf("String value: %s\n", *str)
}

// TestDecodeString test decode int64(Java Long) value
func TestDecodeInt64(t *testing.T) {
	byteArray := []byte{76, 127, 255, 255, 255, 255, 255, 255, 255} // math.MaxInt64

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readLong()
	if math.MaxInt64 != i {
		t.Errorf("int64 value decode error, expect: %d\n", math.MaxInt64)
		t.Fail()
		return
	}
	t.Logf("int64 value: %d\n", i)
}

// TestDecodeInt32 test decode int32(Java Integer) value
func TestDecodeInt32(t *testing.T) {
	byteArray := []byte{73, 127, 255, 255, 255} // math.MaxInt32

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readInt()
	if math.MaxInt32 != i {
		t.Errorf("int32 value decode error, expect: %d\n", math.MaxInt32)
		t.Fail()
		return
	}
	t.Logf("int32 value: %d\n", i)
}

// TestDecodeInt16 test decode int16(Java Short) value
func TestDecodeInt16(t *testing.T) {
	byteArray := []byte{212, 127, 255} // math.MaxInt16

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readShort()
	if math.MaxInt16 != i {
		t.Errorf("int16 value decode error, expect: %d\n", math.MaxInt16)
		t.Fail()
		return
	}
	t.Logf("int16 value: %d\n", i)
}

// TestDecodeInt8ByShort test decode int8(Java Byte) value
func TestDecodeInt8ByShort(t *testing.T) {
	byteArray := []byte{200, 127} // math.MaxInt8

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readShort()
	if math.MaxInt8 != i {
		t.Errorf("int8 value decode error, expect: %d\n", math.MaxInt8)
		t.Fail()
		return
	}
	t.Logf("int8 value: %d\n", i)
}

// TestDecodeInt8ByInt test decode int8(Java Byte) value
func TestDecodeInt8ByInt(t *testing.T) {
	byteArray := []byte{200, 127} // math.MaxInt8

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readInt()
	if math.MaxInt8 != i {
		t.Errorf("int8 value decode error, expect: %d\n", math.MaxInt8)
		t.Fail()
		return
	}
	t.Logf("int8 value: %d\n", i)
}

// TestDecodeFloat64 test decode float64(Java Double) value
func TestDecodeFloat64(t *testing.T) {
	byteArray := []byte{68, 127, 239, 255, 255, 255, 255, 255, 255} // math.MaxFloat64

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readDouble()
	if math.MaxFloat64 != i {
		t.Errorf("float64 value decode error, expect: %v\n", math.MaxFloat64)
		t.Fail()
		return
	}
	t.Logf("float64 value: %v\n", i)
}

// TestDecodeFloat32 test decode float32(Java Float) value
func TestDecodeFloat32(t *testing.T) {
	byteArray := []byte{68, 71, 239, 255, 255, 224, 0, 0, 0} // math.MaxFloat32

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readFloat()
	if math.MaxFloat32 != i {
		t.Errorf("float32 value decode error, expect: %v\n", math.MaxFloat64)
		t.Fail()
		return
	}
	t.Logf("float32 value: %v\n", i)
}

// TestDecodeBooleanTrue test decode bool(Java Boolean) value
func TestDecodeBooleanTrue(t *testing.T) {
	byteArray := []byte{84} // true

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readBoolean()
	if !i {
		t.Errorf("boolean value decode error, expect: %v\n", true)
		t.Fail()
		return
	}
	t.Logf("boolean value: %v\n", i)
}

// TestDecodeBooleanFalse test decode bool(Java Boolean) value
func TestDecodeBooleanFalse(t *testing.T) {
	byteArray := []byte{70} // false

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readBoolean()
	if i {
		t.Errorf("boolean value decode error, expect: %v\n", false)
		t.Fail()
		return
	}
	t.Logf("boolean value: %v\n", i)
}

// TestDecodeNull test decode nil(Java null) value
func TestDecodeNull(t *testing.T) {
	byteArray := []byte{78} // nil

	decoder := NewDecoderWithByteArray(byteArray)
	defer func() {
		if err := recover(); nil != err {
			t.Errorf("boolean value decode error, expect: %v\n", false)
			t.Fail()
			return
		}
	}()
	decoder.readNull()
}

// TestDecodeArray test decode array/slice(Java Array/List) value
func TestDecodeArray(t *testing.T) {
	byteArray := []byte{123, 1, 51, 1, 50, 1, 49} // [3, 2, 1]
	expectValue := "[3 2 1]"

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.ReadObject()
	if nil == i {
		t.Errorf("string array decode error, expect: %s\n", expectValue)
		t.Fail()
		return
	}
	iv := fmt.Sprintf("%v", i)
	if expectValue != iv {
		t.Errorf("string array decode error, expect: %s\n", expectValue)
		t.Fail()
		return
	}
	t.Logf("string array: %v\n", i)
}

// TestDecodeMap test decode map(Java Map) value
func TestDecodeMap(t *testing.T) {
	byteArray := []byte{72, 145, 1, 52, 146, 1, 51, 147, 1, 50, 148, 1, 49, 90} // {1=4, 2=3, 3=2, 4=1}
	expectValue := "map[1:4 2:3 3:2 4:1]"

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.ReadObject()
	if nil == i {
		t.Errorf("map decode error, expect: %s\n", expectValue)
		t.Fail()
		return
	}
	iv := fmt.Sprintf("%v", i)
	if expectValue != iv {
		t.Errorf("map decode error, expect: %s\n", expectValue)
		t.Fail()
		return
	}
	t.Logf("map: %v\n", i)
}

// TestDecodeUTCDate test decode UTC Date value
func TestDecodeUTCDate(t *testing.T) {
	byteArray := []byte{74, 0, 0, 1, 136, 77, 42, 50, 164} // 1684921791140
	var expectValue int64 = 1684921791140

	decoder := NewDecoderWithByteArray(byteArray)
	i := decoder.readUTCDate()
	if expectValue != i {
		t.Errorf("UTC date value decode error, expect: %d\n", expectValue)
		t.Fail()
		return
	}
	t.Logf("UTC date value: %v\n", i)
}
