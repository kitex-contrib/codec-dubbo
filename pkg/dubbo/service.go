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
	"time"

	"github.com/kitex-contrib/codec-hessian2/pkg/iface"
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
	protoVerRaw, err := decoder.Decode()
	if err != nil {
		return err
	}
	protoVer, ok := protoVerRaw.(string)
	if !ok {
		return nil
	}
	svc.ProtocolVersion = protoVer

	pathRaw, err := decoder.Decode()
	if err != nil {
		return err
	}
	path, ok := pathRaw.(string)
	if !ok {
		return nil
	}
	svc.Path = path

	versionRaw, err := decoder.Decode()
	if err != nil {
		return err
	}
	version, ok := versionRaw.(string)
	if !ok {
		return nil
	}
	svc.Version = version

	methodRaw, err := decoder.Decode()
	if err != nil {
		return err
	}
	method, ok := methodRaw.(string)
	if !ok {
		return err
	}
	svc.Method = method

	return nil
}
