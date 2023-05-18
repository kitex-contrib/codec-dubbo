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

const (
	BC_BINARY           = 'B'  // final chunk
	BC_BINARY_CHUNK     = 'A'  // non-final chunk
	BC_BINARY_DIRECT    = 0x20 // 1-byte length binary
	BINARY_DIRECT_MAX   = 0x0f
	BC_BINARY_SHORT     = 0x34  // 2-byte length binary
	BC_BINARY_SHORT_MAX = 0x37  // 2-byte length binary max
	BINARY_SHORT_MAX    = 0x3ff // 0-1023 binary

	BC_CLASS_DEF = 'C' // object/class definition

	BC_DATE        = 0x4a // 64-bit millisecond UTC date
	BC_DATE_MINUTE = 0x4b // 32-bit minute UTC date

	BC_DOUBLE = 'D' // IEEE 64-bit double

	BC_DOUBLE_ZERO  = 0x5b
	BC_DOUBLE_ONE   = 0x5c
	BC_DOUBLE_BYTE  = 0x5d
	BC_DOUBLE_SHORT = 0x5e
	BC_DOUBLE_MILL  = 0x5f
	BC_FALSE        = 'F' // boolean false

	BC_INT = 'I' // 32-bit int

	INT_DIRECT_MIN           = -0x10
	INT_DIRECT_MAX           = 0x2f
	BC_INT_ZERO              = 0x90
	INT_BYTE_MIN             = -0x800
	INT_BYTE_MAX             = 0x7ff
	BC_INT_BYTE_ZERO         = 0xc8
	BC_END                   = 'Z'
	INT_SHORT_MIN            = -0x40000
	INT_SHORT_MAX            = 0x3ffff
	BC_INT_SHORT_ZERO        = 0xd4
	BC_LIST_VARIABLE         = 0x55
	BC_LIST_FIXED            = 'V'
	BC_LIST_VARIABLE_UNTYPED = 0x57
	BC_LIST_FIXED_UNTYPED    = 0x58
	BC_LIST_DIRECT           = 0x70
	BC_LIST_DIRECT_UNTYPED   = 0x78
	LIST_DIRECT_MAX          = 0x7
	BC_LONG                  = 'L' // 64-bit signed integer

	LONG_DIRECT_MIN      int64 = -0x08
	LONG_DIRECT_MAX      int64 = 0x0f
	BC_LONG_ZERO               = 0xe0
	LONG_BYTE_MIN        int64 = -0x800
	LONG_BYTE_MAX        int64 = 0x7ff
	BC_LONG_BYTE_ZERO          = 0xf8
	LONG_SHORT_MIN             = -0x40000
	LONG_SHORT_MAX             = 0x3ffff
	BC_LONG_SHORT_ZERO         = 0x3c
	BC_LONG_INT                = 0x59
	BC_MAP                     = 'M'
	BC_MAP_UNTYPED             = 'H'
	BC_NULL                    = 'N'
	BC_OBJECT                  = 'O'
	BC_OBJECT_DEF              = 'C'
	BC_OBJECT_DIRECT           = 0x60
	BC_OBJECT_DIRECT_MAX       = 0x6f
	OBJECT_DIRECT_MAX          = 0x0f
	BC_REF                     = 0x51
	BC_STRING                  = 'S' // final string
	BC_STRING_CHUNK            = 'R' // non-final string

	BC_STRING_DIRECT    = 0x00
	STRING_DIRECT_MAX   = 0x1f
	BC_STRING_SHORT     = 0x30
	BC_STRING_SHORT_MAX = 0x33
	STRING_SHORT_MAX    = 0x3ff
	BC_TRUE             = 'T'
	P_PACKET_CHUNK      = 0x4f
	P_PACKET            = 'P'
	P_PACKET_DIRECT     = 0x80
	PACKET_DIRECT_MAX   = 0x7f
	P_PACKET_SHORT      = 0x70
	P_PACKET_SHORT_MAX  = 0x77
	PACKET_SHORT_MAX    = 0xfff

	ASCII_CODE_MAX = 0x7f

	CHAR_END_MARK = 0xff

	DIRECT_INT     = 0x80
	DIRECT_INT_MAX = 0xbf

	BYTE_INT     = 0xc0
	BYTE_INT_MAX = 0xc7

	BYTE_INT_LIMIT     = 0xc9
	BYTE_INT_LIMIT_MAX = 0xcf

	SHORT_INT           = 0xd0
	SHORT_INT_MAX       = 0xd7
	SHORT_INT_LIMIT     = 0xd5
	SHORT_INT_LIMIT_MAX = 0xd7

	DIRECT_LONG     = 0xd8
	DIRECT_LONG_MAX = 0xef

	BYTE_LONG       = 0xf0
	BYTE_LONG_MAX   = 0xf7
	BYTE_LONG_LIMIT = 0xf9

	SHORT_LONG           = 0x38
	SHORT_LONG_MAX       = 0x3b
	SHORT_LONG_LIMIT     = 0x3d
	SHORT_LONG_LIMIT_MAX = 0x3f
)
