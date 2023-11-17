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

package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractGroupVersion(t *testing.T) {
	tests := []struct {
		desc     string
		expected func(t *testing.T, remaining, group, version string)
	}{
		{
			desc: "/dubbo/interface:g1:v1",
			expected: func(t *testing.T, remaining, group, version string) {
				assert.Equal(t, "/dubbo/interface", remaining)
				assert.Equal(t, "g1", group)
				assert.Equal(t, "v1", version)
			},
		},
		{
			desc: "/dubbo/interface:g1:",
			expected: func(t *testing.T, remaining, group, version string) {
				assert.Equal(t, "/dubbo/interface", remaining)
				assert.Equal(t, "g1", group)
				assert.Empty(t, version)
			},
		},
		{
			desc: "/dubbo/interface::v1",
			expected: func(t *testing.T, remaining, group, version string) {
				assert.Equal(t, "/dubbo/interface", remaining)
				assert.Empty(t, group)
				assert.Equal(t, "v1", version)
			},
		},
		{
			desc: "/dubbo/interface::",
			expected: func(t *testing.T, remaining, group, version string) {
				assert.Equal(t, "/dubbo/interface", remaining)
				assert.Empty(t, group)
				assert.Empty(t, version)
			},
		},
	}

	for _, test := range tests {
		remaining, group, version := extractGroupVersion(test.desc)
		test.expected(t, remaining, group, version)
	}
}
