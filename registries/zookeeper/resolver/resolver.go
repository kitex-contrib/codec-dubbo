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
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/go-zookeeper/zk"
	"github.com/kitex-contrib/codec-dubbo/registries"
	"strings"
	"time"
)

const (
	defaultSessionTimeout = 30 * time.Second
)

type zookeeperResolver struct {
	conn *zk.Conn
	opt  *Options
}

func NewZookeeperResolver(opts ...Option) (discovery.Resolver, error) {
	o := newOptions(opts)
	conn, _, err := zk.Connect(o.Servers, o.SessionTimeout)
	if err != nil {
		return nil, err
	}
	return &zookeeperResolver{
		conn: conn,
		opt:  o,
	}, nil
}

func (z *zookeeperResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return fmt.Sprintf(registries.RegistryServicesKey, z.opt.RegistryGroup, z.opt.InterfaceName)
}

func (z *zookeeperResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	fmt.Printf("opt.Group: %s, opt.Version: %s\n", z.opt.ServiceGroup, z.opt.ServiceVersion)
	rawURLs, _, err := z.conn.Children(desc)
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
	// todo(DMwangnima): consider this Name since we do not want share a common Resolver
	return strings.Join([]string{"dubbo-zookeeper", z.opt.RegistryGroup, z.opt.InterfaceName, z.opt.ServiceGroup, z.opt.ServiceVersion}, ":")
}
