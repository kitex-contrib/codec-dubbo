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
	"github.com/kitex-contrib/codec-hessian2/pkg/codec"
	commons "github.com/kitex-contrib/codec-hessian2/pkg/common"
)

func getCodec(protocolType commons.ProtocolType) (c Codec) {
	switch protocolType {
	case commons.TTHeader:
	case commons.HTTP2:
	default:
		c = codec.NewDefaultCodec()
	}
	return c
}
