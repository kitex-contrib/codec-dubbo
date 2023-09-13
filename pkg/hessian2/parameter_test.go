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

type testInternalStruct struct {
	Field int
}

type testStructA struct {
	Field int
}

type testStructB struct {
	Internal *testInternalStruct
}

func TestTypesCache_getByData(t *testing.T) {
	tests := []struct {
		desc     string
		datum    []interface{}
		expected func(t *testing.T, c *typesCache)
	}{
		{
			desc: "same structs with basic Type",
			datum: []interface{}{
				&testStructA{Field: 1},
				&testStructA{Field: 2},
			},
			expected: func(t *testing.T, c *typesCache) {
				assert.Equal(t, 1, c.len())
				data := &testStructA{Field: 3}
				typ := reflect.ValueOf(data).Type()
				types, ok := c.get(typ)
				assert.Equal(t, true, ok)
				assert.Equal(t, "J", types)
			},
		},
		{
			desc: "same structs with embedded Type",
			datum: []interface{}{
				&testStructB{
					Internal: &testInternalStruct{
						Field: 1,
					},
				},
				&testStructB{
					Internal: &testInternalStruct{
						Field: 2,
					},
				},
			},
			expected: func(t *testing.T, c *typesCache) {
				assert.Equal(t, 1, c.len())
				data := &testStructB{
					Internal: &testInternalStruct{
						Field: 3,
					},
				}
				typ := reflect.ValueOf(data).Type()
				types, ok := c.get(typ)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Object;", types)
			},
		},
		{
			desc: "different structs",
			datum: []interface{}{
				&testStructA{Field: 1},
				&testStructB{
					Internal: &testInternalStruct{
						Field: 2,
					},
				},
			},
			expected: func(t *testing.T, c *typesCache) {
				assert.Equal(t, 2, c.len())
				dataA := &testStructA{Field: 3}
				dataB := &testStructB{
					Internal: &testInternalStruct{
						Field: 3,
					},
				}
				typA := reflect.ValueOf(dataA).Type()
				typB := reflect.ValueOf(dataB).Type()
				types, ok := c.get(typA)
				assert.Equal(t, true, ok)
				assert.Equal(t, "J", types)
				types, ok = c.get(typB)
				assert.Equal(t, true, ok)
				assert.Equal(t, "Ljava/lang/Object;", types)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			tc := new(typesCache)
			// run getByData concurrently
			for i, data := range test.datum {
				testData := data
				t.Run(fmt.Sprintf("struct%d", i), func(t *testing.T) {
					t.Parallel()
					_, err := tc.getByData(testData)
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
