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
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
)

const (
	// these keys prefixed with "DubboService" are used for user configuring dubbo specific information
	// and are passed within the kitex.
	DubboServiceGroupKey       = "dubbo-service-group"
	DubboServiceVersionKey     = "dubbo-service-version"
	DubboServiceInterfaceKey   = "dubbo-service-interface"
	DubboServiceWeightKey      = "dubbo-service-weight"
	DubboServiceApplicationKey = "dubbo-service-application"

	// these keys prefixed with "dubboInternal" are used for interacting with dubbo-ecosystem
	// and are transferred out of bounds.
	dubboInternalGroupKey       = "group"
	dubboInternalVersionKey     = "version"
	dubboInternalInterfaceKey   = "interface"
	dubboInternalWeightKey      = "weight"
	dubboInternalApplicationKey = "application"

	DefaultRegistryGroup      = "dubbo"
	DefaultProtocol           = "dubbo"
	DefaultDubboServiceWeight = 100

	RegistryServicesKeyTemplate = "/%s/%s/providers"
)

var (
	outboundDubboRegistryKeysMapping = map[string]string{
		DubboServiceGroupKey:       dubboInternalGroupKey,
		DubboServiceVersionKey:     dubboInternalVersionKey,
		DubboServiceInterfaceKey:   dubboInternalInterfaceKey,
		DubboServiceWeightKey:      dubboInternalWeightKey,
		DubboServiceApplicationKey: dubboInternalApplicationKey,
	}

	errMissingInterface = errors.New("tags must contain DubboServiceInterfaceKey:<interfaceName> pair")
)

type URL struct {
	protocol      string
	host          string
	interfaceName string
	params        url.Values
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
	u.interfaceName = strings.TrimPrefix(rawURL.Path, "/")
	u.params = rawURL.Query()
	return nil
}

func (u *URL) ToString() string {
	paramsPart, _ := url.QueryUnescape(u.params.Encode())
	raw := fmt.Sprintf("%s://%s/%s?%s", u.protocol, u.host, u.interfaceName, paramsPart)
	return url.QueryEscape(raw)
}

func (u *URL) ToInstance() discovery.Instance {
	weight := DefaultDubboServiceWeight
	if weightStr := u.params.Get(dubboInternalWeightKey); weightStr != "" {
		if weightParam, err := strconv.Atoi(weightStr); err == nil {
			weight = weightParam
		}
	}
	params := map[string]string{
		DubboServiceGroupKey:   u.params.Get(dubboInternalGroupKey),
		DubboServiceVersionKey: u.params.Get(dubboInternalVersionKey),
	}
	return discovery.NewInstance("tcp", u.host, weight, params)
}

func (u *URL) FromInfo(info *registry.Info) error {
	u.protocol = DefaultProtocol
	if err := u.checkAndSetHost(info.Addr.String()); err != nil {
		return err
	}
	if err := u.filterAndSetParams(info.Tags); err != nil {
		return err
	}
	return nil
}

func (u *URL) GetRegistryServiceKey(registryGroup string) string {
	return fmt.Sprintf(RegistryServicesKeyTemplate, registryGroup, u.interfaceName)
}

func (u *URL) checkAndSetHost(addr string) error {
	sepInd := strings.LastIndex(addr, ":")
	// there is no port part
	if sepInd < 0 {
		return fmt.Errorf("addr %s missing port", addr)
	}
	host := addr[:sepInd]
	port := addr[sepInd+1:]
	finalHost := host
	if host == "" || host == "[::]" {
		ipv4, err := getLocalIPV4Address()
		if err != nil {
			return fmt.Errorf("get local ipv4 error, cause %s", err)
		}
		finalHost = ipv4
		u.host = ipv4 + ":" + port
	}

	u.host = finalHost + ":" + port

	return nil
}

func (u *URL) filterAndSetParams(params map[string]string) error {
	missingInterfaceFlag := true
	if len(params) <= 0 {
		return errMissingInterface
	}
	finalParams := make(url.Values)
	for key, val := range params {
		if dubboKey, ok := outboundDubboRegistryKeysMapping[key]; ok {
			finalParams.Set(dubboKey, val)
		}
		if key == DubboServiceInterfaceKey {
			u.interfaceName = val
			missingInterfaceFlag = false
		}
	}
	if missingInterfaceFlag {
		return errMissingInterface
	}
	u.params = finalParams
	return nil
}

func getLocalIPV4Address() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New("there is not valid ip address found")
}
