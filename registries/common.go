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

package registries

import (
	"net/url"
	"strconv"

	"github.com/cloudwego/kitex/pkg/discovery"
)

const (
	DubboServiceProtocolKey = "dubbo-service-protocol"
	DubboServiceGroupKey    = "dubbo-service-group"
	DubboServiceVersionKey  = "dubbo-service-version"

	DefaultDubboServiceWeight = 100
)

var RegistryServicesKey = "/%s/%s/providers"

type URL struct {
	protocol string
	host     string
	params   url.Values
}

func (u *URL) FromString(raw string) error {
	decodedRaw, err := url.PathUnescape(raw)
	if err != nil {
		return err
	}
	rawURL, err := url.Parse(decodedRaw)
	if err != nil {
		return err
	}
	u.protocol = rawURL.Scheme
	u.host = rawURL.Host
	u.params = rawURL.Query()
	return nil
}

func (u *URL) ToInstance() discovery.Instance {
	weight := DefaultDubboServiceWeight
	if weightStr := u.params.Get("weight"); weightStr != "" {
		if weightParam, err := strconv.Atoi(weightStr); err == nil {
			weight = weightParam
		}
	}
	params := map[string]string{
		DubboServiceProtocolKey: u.protocol,
		DubboServiceGroupKey:    u.params.Get("group"),
		DubboServiceVersionKey:  u.params.Get("version"),
	}
	return discovery.NewInstance("tcp", u.host, weight, params)
}
