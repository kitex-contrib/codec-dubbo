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
 */

package hessian2

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Data interface{}
	Anno *TypeAnnotation
}

type testInternalStruct struct {
	Field int8
}

type testStructA struct {
	Field int8
}

type testStructB struct {
	Internal *testInternalStruct
}

type testStructC struct {
	FieldA int8
	FieldB int64
	FieldC float64
	FieldD string
	FieldE []int32
	FieldF []string
}

func TestTypesCache_getByData(t *testing.T) {
	tests := []struct {
		desc     string
		datum    []testStruct
		expected func(t *testing.T, c *methodCache)
	}{
		{
			desc: "same structs with basic Type",
			datum: []testStruct{
				{Data: &testStructA{Field: 1}},
				{Data: &testStructA{Field: 2}},
			},
			expected: func(t *testing.T, c *methodCache) {
				assert.Equal(t, 1, c.len())
				data := &testStructA{Field: 3}
				key := methodKey{typ: reflect.ValueOf(data).Type()}
				types, ok := c.get(key)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Byte;", types)
			},
		},
		{
			desc: "same structs with embedded Type",
			datum: []testStruct{
				{Data: &testStructB{
					Internal: &testInternalStruct{
						Field: 1,
					},
				}},
				{Data: &testStructB{
					Internal: &testInternalStruct{
						Field: 2,
					},
				}},
			},
			expected: func(t *testing.T, c *methodCache) {
				assert.Equal(t, 1, c.len())
				data := &testStructB{
					Internal: &testInternalStruct{
						Field: 3,
					},
				}
				key := methodKey{typ: reflect.ValueOf(data).Type()}
				types, ok := c.get(key)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Object;", types)
			},
		},
		{
			desc: "different structs",
			datum: []testStruct{
				{Data: &testStructA{Field: 1}},
				{Data: &testStructB{
					Internal: &testInternalStruct{
						Field: 2,
					},
				}},
			},
			expected: func(t *testing.T, c *methodCache) {
				assert.Equal(t, 2, c.len())
				dataA := &testStructA{Field: 3}
				dataB := &testStructB{
					Internal: &testInternalStruct{
						Field: 3,
					},
				}
				keyA := methodKey{typ: reflect.ValueOf(dataA).Type()}
				keyB := methodKey{typ: reflect.ValueOf(dataB).Type()}
				types, ok := c.get(keyA)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Byte;", types)
				types, ok = c.get(keyB)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Object;", types)
			},
		},
		{
			desc: "use annotations to specify types",
			datum: []testStruct{
				{
					Data: &testStructA{Field: 1},
					Anno: NewTypeAnnotation("byte"),
				},
				{
					Data: &testStructB{
						Internal: &testInternalStruct{
							Field: 2,
						},
					},
					Anno: NewTypeAnnotation("java.lang.Object"),
				},
				{
					Data: &testStructC{
						FieldA: 3,
						FieldB: 4,
						FieldC: 5.0,
						FieldD: "6",
						FieldE: []int32{7, 8},
						FieldF: []string{"9", "10"},
					},
					Anno: NewTypeAnnotation("byte,long,double,java.lang.String,int[],java.lang.String[]"),
				},
				{
					Data: &testStructC{
						FieldA: 3,
						FieldB: 4,
						FieldC: 5.0,
						FieldD: "6",
						FieldE: []int32{7, 8},
						FieldF: []string{"9", "10"},
					},
					Anno: NewTypeAnnotation("-,-,-,-,-,-"),
				},
			},
			expected: func(t *testing.T, c *methodCache) {
				assert.Equal(t, 4, c.len())

				keyA := methodKey{typ: reflect.ValueOf(&testStructA{}).Type(), anno: "byte"}
				types, ok := c.get(keyA)
				assert.Equal(t, true, ok)
				assert.Equal(t, "B", types)

				keyB := methodKey{typ: reflect.ValueOf(&testStructB{}).Type(), anno: "java.lang.Object"}
				types, ok = c.get(keyB)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Object;", types)

				keyC := methodKey{typ: reflect.ValueOf(&testStructC{}).Type(), anno: "byte,long,double,java.lang.String,int[],java.lang.String[]"}
				types, ok = c.get(keyC)
				assert.Equal(t, true, ok)
				assert.Equal(t, "BJDLjava/lang/String;[I[Ljava/lang/String;", types)

				keyD := methodKey{typ: reflect.ValueOf(&testStructC{}).Type(), anno: "-,-,-,-,-,-"}
				types, ok = c.get(keyD)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Byte;Ljava/lang/Long;Ljava/lang/Double;Ljava/lang/String;Ljava/util/List;Ljava/util/List;", types)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			tc := new(methodCache)
			// run getByData concurrently
			for i, data := range test.datum {
				testData := data
				t.Run(fmt.Sprintf("struct%d", i), func(t *testing.T) {
					t.Parallel()
					_, err := tc.getTypes(testData.Data, testData.Anno)
					if err != nil {
						t.Fatal(err)
					}
				})
			}
			t.Cleanup(func() {
				test.expected(t, tc)
			})
		})
	}
}
