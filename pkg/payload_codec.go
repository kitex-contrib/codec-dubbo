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
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
)

type Message struct {
}

type ByteBuffer struct {
}

// Hessian2Codec NewHessian2Codec creates the hessian2 codec.
type Hessian2Codec struct {
}

// NewHessian2Codec creates a new codec instance.
func NewHessian2Codec() *Hessian2Codec {
	// TOOD: WIP
	panic("not implemented")
}

// Marshal encode method
func (m *Hessian2Codec) Marshal(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	// TOOD: WIP
	panic("not implemented")
}

// Unmarshal decode method
func (m *Hessian2Codec) Unmarshal(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	// TOOD: WIP
	panic("not implemented")
}

// Name codec name
func (m *Hessian2Codec) Name() string {
	return "hessian2"
}
