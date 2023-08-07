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

package dubbo

import (
	"fmt"

	"github.com/kitex-contrib/codec-hessian2/pkg/iface"
)

type PayloadType int32

// Response payload type enum
const (
	RESPONSE_WITH_EXCEPTION PayloadType = iota
	RESPONSE_VALUE
	RESPONSE_NULL_VALUE
	RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS
	RESPONSE_VALUE_WITH_ATTACHMENTS
	RESPONSE_NULL_VALUE_WITH_ATTACHMENTS
)

var (
	attachmentsPair = map[PayloadType]PayloadType{
		RESPONSE_WITH_EXCEPTION: RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS,
		RESPONSE_VALUE:          RESPONSE_VALUE_WITH_ATTACHMENTS,
		RESPONSE_NULL_VALUE:     RESPONSE_NULL_VALUE_WITH_ATTACHMENTS,
	}
	attachmentsSet = map[PayloadType]struct{}{
		RESPONSE_WITH_EXCEPTION_WITH_ATTACHMENTS: {},
		RESPONSE_VALUE_WITH_ATTACHMENTS:          {},
		RESPONSE_NULL_VALUE_WITH_ATTACHMENTS:     {},
	}
)

// GetAttachmentsPayloadType returns base PayloadType or base with attachments PayloadType based on expression.
// If base PayloadType does not have responding attachments PayloadType, returns itself.
func GetAttachmentsPayloadType(expression bool, base PayloadType) PayloadType {
	if expression {
		if pair, ok := attachmentsPair[base]; ok {
			return pair
		}
	}

	return base
}

// IsAttachmentsPayloadType determines whether typ is an attachments PayloadType
func IsAttachmentsPayloadType(typ PayloadType) bool {
	_, ok := attachmentsSet[typ]
	return ok
}

func DecodePayloadType(decoder iface.Decoder) (PayloadType, error) {
	payloadTypeRaw, err := decoder.Decode()
	if err != nil {
		return 0, err
	}
	payloadTypeInt32, ok := payloadTypeRaw.(int32)
	if !ok {
		return 0, fmt.Errorf("dubbo PayloadType decoded failed, got: %v", payloadTypeRaw)
	}
	return PayloadType(payloadTypeInt32), nil
}
