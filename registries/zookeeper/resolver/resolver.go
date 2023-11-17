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
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/go-zookeeper/zk"
	"github.com/kitex-contrib/codec-dubbo/registries"
)

const (
	defaultSessionTimeout = 30 * time.Second
	groupVersionSeparator = ":"
)

type zookeeperResolver struct {
	conn       *zk.Conn
	opt        *Options
	uniqueName string
}

func NewZookeeperResolver(opts ...Option) (discovery.Resolver, error) {
	o := newOptions(opts)
	conn, _, err := zk.Connect(o.Servers, o.SessionTimeout)
	if err != nil {
		return nil, err
	}
	if o.Username != "" && o.Password != "" {
		if err := conn.AddAuth("digest", []byte(fmt.Sprintf("%s:%s", o.Username, o.Password))); err != nil {
			return nil, err
		}
	}
	uniName := "dubbo-zookeeper" + "/" + o.RegistryGroup
	return &zookeeperResolver{
		conn:       conn,
		opt:        o,
		uniqueName: uniName,
	}, nil
}

func (z *zookeeperResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	interfaceName, ok := target.Tag(registries.DubboServiceInterfaceKey)
	if !ok {
		panic("please specify target dubbo interface with \"client.WithTag(registries.DubboServiceInterfaceKey, <interfaceName>)")
	}
	group := target.DefaultTag(registries.DubboServiceGroupKey, "")
	version := target.DefaultTag(registries.DubboServiceVersionKey, "")
	regSvcKey := fmt.Sprintf(registries.RegistryServicesKeyTemplate, z.opt.RegistryGroup, interfaceName)
	return regSvcKey + groupVersionSeparator + group + groupVersionSeparator + version
}

func (z *zookeeperResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	regSvcKey, svcGroup, svcVersion := extractGroupVersion(desc)
	rawURLs, _, err := z.conn.Children(regSvcKey)
	if err != nil {
		return discovery.Result{}, err
	}
	instances := make([]discovery.Instance, 0, len(rawURLs))
	for _, rawURL := range rawURLs {
		u := new(registries.URL)
		if err := u.FromString(rawURL); err != nil {
			klog.Errorf("invalid dubbo URL from zookeeper: %s, err :%s", rawURL, err)
			continue
		}
		tmpInstance := u.ToInstance()
		if group, _ := tmpInstance.Tag(registries.DubboServiceGroupKey); group != svcGroup {
			continue
		}
		if ver, _ := tmpInstance.Tag(registries.DubboServiceVersionKey); ver != svcVersion {
			continue
		}
		instances = append(instances, tmpInstance)
	}
	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}, nil
}

func (z *zookeeperResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (z *zookeeperResolver) Name() string {
	return z.uniqueName
}

// extractGroupVersion extract group and version from desc returned by Target()
// e.g.
// input: desc /dubbo/interfaceName:g1:v1
//
// output: remaining /dubbo/interfaceName
//
//	group g1
//	version v1
func extractGroupVersion(desc string) (remaining, group, version string) {
	// retrieve version
	verSepIdx := strings.LastIndex(desc, groupVersionSeparator)
	version = desc[verSepIdx+1:]
	remaining = desc[:verSepIdx]

	// retrieve group
	groSepIdx := strings.LastIndex(remaining, groupVersionSeparator)
	group = remaining[groSepIdx+1:]
	remaining = remaining[:groSepIdx]

	return
}
