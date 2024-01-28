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
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReflectResponse(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		tests := []struct {
			desc      string
			testFunc  func(t *testing.T, expectedErr bool)
			expectErr bool
		}{
			{
				desc: "map[bool]bool",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]bool
					src := map[bool]bool{
						true: true,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]int8",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]int8
					src := map[bool]int8{
						true: 12,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]int16",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]int16
					src := map[bool]int16{
						true: 12,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]int32",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]int32
					src := map[bool]int32{
						true: 12,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]int64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]int64
					src := map[bool]int64{
						true: 12,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]float64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]float64
					src := map[bool]float64{
						true: 12.34,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool]string",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool]string
					src := map[bool]string{
						true: "12",
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]byte",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]byte
					src := map[bool][]byte{
						true: {
							'1',
							'2',
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]bool",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]bool
					src := map[bool][]bool{
						true: {
							true,
							true,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]int8",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]int8
					src := map[bool][]int8{
						true: {
							1,
							2,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]int16",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]int16
					src := map[bool][]int16{
						true: {
							1,
							2,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]int32",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]int32
					src := map[bool][]int32{
						true: {
							1,
							2,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]int64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]int64
					src := map[bool][]int64{
						true: {
							1,
							2,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]float64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]float64
					src := map[bool][]float64{
						true: {
							1.2,
							3.4,
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][]string",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][]string
					src := map[bool][]string{
						true: {
							"12",
							"34",
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "map[bool][][]byte",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest map[bool][][]byte
					src := map[bool][][]byte{
						true: {
							{
								'1',
								'2',
							},
							{
								'3',
								'4',
							},
						},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
		}

		for _, test := range tests {
			t.Run(test.desc, func(t *testing.T) {
				test.testFunc(t, test.expectErr)
			})
		}
	})
	t.Run("list", func(t *testing.T) {
		tests := []struct {
			desc      string
			testFunc  func(t *testing.T, expectedErr bool)
			expectErr bool
		}{
			{
				desc: "[]bool",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []bool
					src := []bool{
						true, true,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]int8",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []int8
					src := []int8{
						12, 34,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]int16",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []int16
					src := []int16{
						12, 34,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]int32",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []int32
					src := []int32{
						12, 34,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]int64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []int64
					src := []int64{
						12, 34,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]float64",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []float64
					src := []float64{
						12.34, 56.78,
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]string",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []string
					src := []string{
						"12", "34",
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[][]byte",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest [][]byte
					src := [][]byte{
						{1, 2}, {3, 4},
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]time.Time",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []time.Time
					src := []time.Time{
						time.Unix(1000, 0),
						time.Unix(1001, 0),
					}
					testReflectResponse(t, src, &dest, expectedErr)
					if !reflect.DeepEqual(src, dest) {
						t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					}
				},
			},
			{
				desc: "[]*time.Time",
				testFunc: func(t *testing.T, expectedErr bool) {
					var dest []*time.Time
					src := []interface{}{
						time.Unix(1000, 0),
						time.Unix(1001, 0),
					}
					testReflectResponse(t, src, &dest, expectedErr)
					assert.Equal(t, len(src), len(dest))
					for i, ptr := range dest {
						assert.Equal(t, src[i], *ptr)
					}
					//if !reflect.DeepEqual(src, dest) {
					//	t.Fatalf("src: %+v, dest: %+v, they are not equal", src, dest)
					//}
				},
			},
		}

		for _, test := range tests {
			t.Run(test.desc, func(t *testing.T) {
				test.testFunc(t, test.expectErr)
			})
		}
	})
}

func testReflectResponse(t *testing.T, src, dest interface{}, expectErr bool) {
	if err := ReflectResponse(src, dest); err != nil {
		if !expectErr {
			t.Fatal(err)
		}
		return
	}
}
