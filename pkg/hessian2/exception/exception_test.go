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

package exception

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromError(t *testing.T) {
	tests := []struct {
		desc     string
		inputErr error
		expected func(t *testing.T, exception Throwabler, ok bool)
	}{
		{
			desc:     "nil err",
			inputErr: nil,
			expected: func(t *testing.T, exception Throwabler, ok bool) {
				assert.Nil(t, exception)
				assert.False(t, ok)
			},
		},
		{
			desc:     "Throwabler err",
			inputErr: NewException("FromError test"),
			expected: func(t *testing.T, exception Throwabler, ok bool) {
				assert.Equal(t, "java.lang.Exception", exception.JavaClassName())
				assert.Equal(t, "FromError test", exception.Error())
				assert.True(t, ok)
			},
		},
		{
			desc:     "DetailedError wraps Throwabler",
			inputErr: kerrors.ErrRemoteOrNetwork.WithCause(NewException("FromError test")),
			expected: func(t *testing.T, exception Throwabler, ok bool) {
				assert.Equal(t, "java.lang.Exception", exception.JavaClassName())
				assert.Equal(t, "FromError test", exception.Error())
				assert.True(t, ok)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			exception, ok := FromError(test.inputErr)
			test.expected(t, exception, ok)
		})
	}
}
