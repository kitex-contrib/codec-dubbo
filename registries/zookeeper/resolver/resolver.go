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
	defaultRegistryGroup  = "dubbo"
	defaultSessionTimeout = 30 * time.Second
)

type zookeeperResolver struct {
	conn *zk.Conn
	opt  *Options
	// registryServicePath is the path used to retrieve instances information from zookeeper.
	// format: /<RegistryGroup>/<InterfaceName>/providers
	registryServicePath string
	// uniqueName is the pre-calculated result returned by Name() and Target().
	// format: /<registryServicePath>/<ServiceGroup>/<ServiceVersion>
	// since each group or each version of the service belongs to a different BalancerFactory-Resolver,
	// we should add group and version information to it for caching.
	uniqueName string
}

func NewZookeeperResolver(opts ...Option) (discovery.Resolver, error) {
	o := newOptions(opts)
	conn, _, err := zk.Connect(o.Servers, o.SessionTimeout)
	if err != nil {
		return nil, err
	}
	regSvcPath := fmt.Sprintf(registries.RegistryServicesKey, o.RegistryGroup, o.InterfaceName)
	uniName := strings.Join([]string{"dubbo-zookeeper", regSvcPath, o.ServiceGroup, o.ServiceVersion}, "/")
	return &zookeeperResolver{
		conn:                conn,
		opt:                 o,
		registryServicePath: regSvcPath,
		uniqueName:          uniName,
	}, nil
}

func (z *zookeeperResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return z.uniqueName
}

func (z *zookeeperResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	rawURLs, _, err := z.conn.Children(z.registryServicePath)
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
		if group, _ := tmpInstance.Tag(registries.DubboServiceGroupKey); group != z.opt.ServiceGroup {
			continue
		}
		if ver, _ := tmpInstance.Tag(registries.DubboServiceVersionKey); ver != z.opt.ServiceVersion {
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
