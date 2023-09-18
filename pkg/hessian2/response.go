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
 *
 * This source file has been replicated from the original dubbo-go-hessian2
 * project repository, and we extend our sincere appreciation to the
 * dubbo-go development team for their valuable contribution.
 */

package hessian2

import (
	"errors"
	"fmt"
	"reflect"
)

// _Rune is an alias for rune, so that to get the correct runtime type of rune.
// The runtime type of rune is int32, which is not expected.
type _Rune rune

var (
	_varRune       = _Rune(0)
	_typeOfRunePtr = reflect.TypeOf(&_varRune)
)

var (
	// reflect.PointerTo is supported only from go1.20.
	// So we use a dummy variable to get the pointer type.
	_varInt     = 0
	_varInt8    = int8(0)
	_varInt16   = int16(0)
	_varInt32   = int32(0)
	_varInt64   = int64(0)
	_varUint    = uint(0)
	_varUint8   = uint8(0)
	_varUint16  = uint16(0)
	_varUint32  = uint32(0)
	_varUint64  = uint64(0)
	_varFloat32 = float32(0)
	_varFloat64 = float64(0)
)

var (
	_typeOfIntPtr     = reflect.TypeOf(&_varInt)
	_typeOfInt8Ptr    = reflect.TypeOf(&_varInt8)
	_typeOfInt16Ptr   = reflect.TypeOf(&_varInt16)
	_typeOfInt32Ptr   = reflect.TypeOf(&_varInt32)
	_typeOfInt64Ptr   = reflect.TypeOf(&_varInt64)
	_typeOfUintPtr    = reflect.TypeOf(&_varUint)
	_typeOfUint8Ptr   = reflect.TypeOf(&_varUint8)
	_typeOfUint16Ptr  = reflect.TypeOf(&_varUint16)
	_typeOfUint32Ptr  = reflect.TypeOf(&_varUint32)
	_typeOfUint64Ptr  = reflect.TypeOf(&_varUint64)
	_typeOfFloat32Ptr = reflect.TypeOf(&_varFloat32)
	_typeOfFloat64Ptr = reflect.TypeOf(&_varFloat64)
)

// _refHolder is used to record decode list, the address of which may change when appending more element.
type _refHolder struct {
	// destinations
	destinations []reflect.Value
}

// add destination
func (h *_refHolder) add(dest reflect.Value) {
	h.destinations = append(h.destinations, dest)
}

// the ref holder pointer type.
var _refHolderPtrType = reflect.TypeOf(&_refHolder{})

// ReflectResponse reflect return value
func ReflectResponse(in, out interface{}) error {
	if in == nil {
		return fmt.Errorf("@in is nil")
	}

	if out == nil {
		return fmt.Errorf("@out is nil")
	}
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("@out should be a pointer")
	}

	inValue := ensurePackValue(in)
	outValue := ensurePackValue(out)

	outType := outValue.Type().String()
	if outType == "interface {}" || outType == "*interface {}" {
		setValue(outValue, inValue)
		return nil
	}

	switch inValue.Type().Kind() {
	case reflect.Slice, reflect.Array:
		return copySlice(inValue, outValue)
	case reflect.Map:
		return copyMap(inValue, outValue)
	default:
		setValue(outValue, inValue)
	}

	return nil
}

// setValue set the value to dest.
// It will auto check the Ptr pack level and unpack/pack to the right level.
// It makes sure success to set value
func setValue(dest, v reflect.Value) {
	// zero value not need to set
	if !v.IsValid() {
		return
	}

	vType := v.Type()
	destType := dest.Type()

	// for most cases, the types are the same and can set the value directly.
	if dest.CanSet() && destType == vType {
		dest.Set(v)
		return
	}

	// check whether the v is a ref holder
	if vType == _refHolderPtrType {
		h := v.Interface().(*_refHolder)
		h.add(dest)
		return
	}

	vRawType, vPtrDepth := unpackType(vType)

	// unpack to the root addressable value, so that to set the value.
	dest = unpackToRootAddressableValue(dest)
	destType = dest.Type()
	destRawType, destPtrDepth := unpackType(destType)

	// it can set the value directly if the raw types are of the same type.
	if destRawType == vRawType {
		if destPtrDepth > vPtrDepth {
			// pack to the same level of dest
			for i := 0; i < destPtrDepth-vPtrDepth; i++ {
				v = packPtr(v)
			}
		} else if destPtrDepth < vPtrDepth {
			// unpack to the same level of dest
			for i := 0; i < vPtrDepth-destPtrDepth; i++ {
				v = v.Elem()
			}
		}

		dest.Set(v)

		return
	}

	if vRawType.String() == "interface {}" {
		v = v.Elem()
	}
	switch destType.Kind() {
	case reflect.Float32, reflect.Float64:
		dest.SetFloat(v.Float())
		return
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dest.SetInt(v.Int())
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// hessian only support 64-bit signed long integer.
		dest.SetUint(uint64(v.Int()))
		return
	case reflect.Ptr:
		setValueToPtrDest(dest, v)
		return
	case reflect.Bool:
		dest.SetBool(v.Bool())
	default:
		// It's ok when the dest is an interface{}, while the v is a pointer.
		dest.Set(v)
	}
}

// copySlice copy from inSlice to outSlice
func copySlice(inSlice, outSlice reflect.Value) error {
	if inSlice.IsNil() {
		return errors.New("@in is nil")
	}
	if inSlice.Kind() != reflect.Slice {
		return fmt.Errorf("@in is not slice, but %v", inSlice.Kind())
	}

	for outSlice.Kind() == reflect.Ptr {
		outSlice = outSlice.Elem()
	}

	size := inSlice.Len()
	outSlice.Set(reflect.MakeSlice(outSlice.Type(), size, size))

	for i := 0; i < size; i++ {
		inSliceValue := inSlice.Index(i)
		outSliceValue := reflect.New(outSlice.Index(i).Type()).Elem()
		setValue(outSliceValue, inSliceValue)
		outSlice.Index(i).Set(outSliceValue)
	}

	return nil
}

// copyMap copy from in map to out map
func copyMap(inMapValue, outMapValue reflect.Value) error {
	if inMapValue.IsNil() {
		return errors.New("@in is nil")
	}
	if !inMapValue.CanInterface() {
		return errors.New("@in's Interface can not be used")
	}
	if inMapValue.Kind() != reflect.Map {
		return fmt.Errorf("@in is not map, but %v", inMapValue.Kind())
	}

	outMapType := unpackPtrType(outMapValue.Type())
	setValue(outMapValue, reflect.MakeMap(outMapType))

	outKeyType := outMapType.Key()

	outMapValue = unpackPtrValue(outMapValue)
	outValueType := outMapValue.Type().Elem()

	for _, inKey := range inMapValue.MapKeys() {
		inValue := inMapValue.MapIndex(inKey)
		outKey := reflect.New(outKeyType).Elem()
		setValue(outKey, inKey)
		outValue := reflect.New(outValueType).Elem()
		setValue(outValue, inValue)

		outMapValue.SetMapIndex(outKey, outValue)
	}

	return nil
}

// unpackPtrType unpack pointer type to original type
func unpackPtrType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

// unpackPtrValue unpack pointer value to original value
// return the pointer if its elem is zero value, because lots of operations on zero value is invalid
func unpackPtrValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr && v.Elem().IsValid() {
		v = v.Elem()
	}
	return v
}

// ensurePackValue pack the interface with value
func ensurePackValue(in interface{}) reflect.Value {
	if v, ok := in.(reflect.Value); ok {
		return v
	}
	return reflect.ValueOf(in)
}

// unpackType unpack pointer type to original type and return the pointer depth.
func unpackType(typ reflect.Type) (reflect.Type, int) {
	depth := 0
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		depth++
	}
	return typ, depth
}

// unpackToRootAddressableValue unpack pointer value to the root addressable value.
func unpackToRootAddressableValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr && v.Elem().CanAddr() {
		v = v.Elem()
	}
	return v
}

// packPtr pack a Ptr value
func packPtr(v reflect.Value) reflect.Value {
	vv := reflect.New(v.Type())
	vv.Elem().Set(v)
	return vv
}

// setValueToPtrDest set the raw value to a pointer dest.
func setValueToPtrDest(dest, v reflect.Value) {
	// for number, the type of value may be different with the dest,
	// must convert it to the correct type of value then set.
	switch dest.Type() {
	case _typeOfIntPtr:
		vv := v.Int()
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfInt8Ptr:
		vv := int8(v.Int())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfInt16Ptr:
		vv := int16(v.Int())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfInt32Ptr:
		if v.Kind() == reflect.String {
			vv := rune(v.String()[0])
			dest.Set(reflect.ValueOf(&vv))
			return
		}
		vv := int32(v.Int())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfInt64Ptr:
		vv := v.Int()
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfUintPtr:
		vv := uint(v.Uint())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfUint8Ptr:
		// v is a int32 here.
		vv := uint8(v.Int())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfUint16Ptr:
		vv := uint16(v.Uint())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfUint32Ptr:
		vv := uint32(v.Uint())
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfUint64Ptr:
		vv := v.Uint()
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfFloat32Ptr:
		vv := float32(v.Float())
		dest.Set(reflect.ValueOf(&vv))
	case _typeOfFloat64Ptr:
		vv := v.Float()
		dest.Set(reflect.ValueOf(&vv))
		return
	case _typeOfRunePtr:
		if v.Kind() == reflect.String {
			vv := _Rune(v.String()[0])
			dest.Set(reflect.ValueOf(&vv))
			return
		}
		vv := _Rune(v.Int())
		dest.Set(reflect.ValueOf(&vv))
		return
	default:
		dest.Set(v)
	}
}
