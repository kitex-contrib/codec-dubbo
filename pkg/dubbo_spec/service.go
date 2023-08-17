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

package dubbo_spec

import (
	"fmt"
	"time"

	"github.com/kitex-contrib/codec-dubbo/pkg/iface"
)

const DEFAULT_DUBBO_PROTOCOL_VERSION = "2.0.2"

type Service struct {
	ProtocolVersion string
	Path            string
	Version         string
	Method          string
	Timeout         time.Duration
	Group           string
}

func (svc *Service) Decode(decoder iface.Decoder) error {
	if err := decodeString(decoder, &svc.ProtocolVersion, "ProtocolVersion"); err != nil {
		return err
	}
	if err := decodeString(decoder, &svc.Path, "Path"); err != nil {
		return err
	}
	if err := decodeString(decoder, &svc.Version, "Version"); err != nil {
		return err
	}
	if err := decodeString(decoder, &svc.Method, "Method"); err != nil {
		return err
	}

	return nil
}

// decodeString decodes dubbo Service string field
func decodeString(decoder iface.Decoder, target *string, targetName string) error {
	strRaw, err := decoder.Decode()
	if err != nil {
		return err
	}
	str, ok := strRaw.(string)
	if !ok {
		return fmt.Errorf("decode dubbo Service field %s failed, got %v", targetName, strRaw)
	}
	*target = str
	return nil
}
