/*
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
 */

package commons

// constants
const (
	BC_NULL = byte('N') // x4e

	BC_TRUE  = byte('T') // boolean true
	BC_FALSE = byte('F') // boolean false

	BC_BINARY       = byte('B') // final chunk
	BC_BINARY_CHUNK = byte('A') // non-final chunk

	BC_BINARY_DIRECT  = byte(0x20) // 1-byte length binary
	BINARY_DIRECT_MAX = byte(0x0f)
	BC_DOUBLE         = byte('D') // IEEE 64-bit double

	INT_DIRECT_MIN = -0x10
	INT_DIRECT_MAX = byte(0x2f)
	BC_INT_ZERO    = byte(0x90)

	INT_BYTE_MIN      = -0x800
	INT_BYTE_MAX      = 0x7ff
	INT_SHORT_MIN     = -0x40000
	INT_SHORT_MAX     = 0x3ffff
	BC_INT_SHORT_ZERO = byte(0xd4)
	BC_LIST_FIXED     = byte('V')

	BC_LONG = byte('L') // 64-bit signed integer

	BC_MAP = byte('M')

	BC_STRING       = byte('S') // final string
	BC_STRING_CHUNK = byte('R') // non-final string
	BC_OBJECT       = byte('O')
	BC_OBJECT_DEF   = byte('C')
	BC_INT          = byte('I')
	BC_LIST         = byte('Z')
	BC_MAP_NON_TYPE = byte('H') //non-type key map

	PACKET_SHORT_MAX = 0xfff
)
