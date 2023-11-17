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

package registry

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/go-zookeeper/zk"
	"github.com/kitex-contrib/codec-dubbo/registries"
)

const (
	defaultSessionTimeout = 30 * time.Second
)

type zookeeperRegistry struct {
	conn     *zk.Conn
	opt      *Options
	canceler *canceler
}

func NewZookeeperRegistry(opts ...Option) (registry.Registry, error) {
	o := newOptions(opts)
	conn, eventChan, err := zk.Connect(o.Servers, o.SessionTimeout)
	if err != nil {
		return nil, err
	}
	if o.Username != "" && o.Password != "" {
		if err := conn.AddAuth("digest", []byte(fmt.Sprintf("%s:%s", o.Username, o.Password))); err != nil {
			return nil, err
		}
	}
	// This connection timeout should not exceed sessionTimeout and should not be too small.
	// So just pick halfTimeout in the middle range.
	halfTimeout := o.SessionTimeout / 2
	ticker := time.NewTimer(halfTimeout)
	for {
		select {
		case event := <-eventChan:
			if event.State == zk.StateConnected {
				return &zookeeperRegistry{
					conn:     conn,
					opt:      o,
					canceler: newCanceler(),
				}, nil
			}
		case <-ticker.C:
			return nil, fmt.Errorf("waiting for zookeeper connected time out: elapsed %d seconds", halfTimeout/time.Second)
		}
	}
}

func (z *zookeeperRegistry) Register(info *registry.Info) error {
	if info == nil {
		return nil
	}
	u := new(registries.URL)
	if err := u.FromInfo(info); err != nil {
		return err
	}
	path := u.GetRegistryServiceKey(z.opt.RegistryGroup)
	content := u.ToString()
	finalPath := path + "/" + content
	if err := z.createNode(finalPath, nil, true); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	z.canceler.add(finalPath, cancel)
	go z.keepalive(ctx, finalPath, nil)
	return nil
}

func (z *zookeeperRegistry) createNode(path string, content []byte, ephemeral bool) error {
	exists, stat, err := z.conn.Exists(path)
	if err != nil {
		return err
	}
	// ephemeral nodes handling after restart
	// fixes a race condition if the server crashes without using CreateProtectedEphemeralSequential()
	// https://github.com/go-kratos/kratos/blob/main/contrib/registry/zookeeper/register.go
	if exists && ephemeral {
		err = z.conn.Delete(path, stat.Version)
		if err != nil && err != zk.ErrNoNode {
			return err
		}
		exists = false
	}
	if !exists {
		i := strings.LastIndex(path, "/")
		if i > 0 {
			err := z.createNode(path[0:i], nil, false)
			if err != nil && !errors.Is(err, zk.ErrNodeExists) {
				return err
			}
		}
		var flag int32
		if ephemeral {
			flag = zk.FlagEphemeral
		}
		if z.opt.Username != "" && z.opt.Password != "" {
			_, err = z.conn.Create(path, content, flag, zk.DigestACL(zk.PermAll, z.opt.Username, z.opt.Password))
		} else {
			_, err = z.conn.Create(path, content, flag, zk.WorldACL(zk.PermAll))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *zookeeperRegistry) keepalive(ctx context.Context, path string, content []byte) {
	sessionID := z.conn.SessionID()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cur := z.conn.SessionID()
			if cur != 0 && sessionID != cur {
				if err := z.createNode(path, content, true); err == nil {
					sessionID = cur
				}
			}
		}
	}
}

func (z *zookeeperRegistry) Deregister(info *registry.Info) error {
	if info == nil {
		return nil
	}
	u := new(registries.URL)
	if err := u.FromInfo(info); err != nil {
		return err
	}
	path := u.GetRegistryServiceKey(z.opt.RegistryGroup)
	content := u.ToString()
	finalPath := path + "/" + content
	cancel, ok := z.canceler.remove(finalPath)
	if ok {
		cancel()
	}
	return z.deleteNode(finalPath)
}

func (z *zookeeperRegistry) deleteNode(path string) error {
	return z.conn.Delete(path, -1)
}

type canceler struct {
	mu sync.Mutex
	// key: zookeeper path, val: CancelFunc to stop the keepalive functionality of this zookeeper path
	cancelMap map[string]context.CancelFunc
}

func newCanceler() *canceler {
	return &canceler{
		cancelMap: make(map[string]context.CancelFunc),
	}
}

func (c *canceler) add(path string, f context.CancelFunc) {
	c.mu.Lock()
	c.cancelMap[path] = f
	c.mu.Unlock()
}

// remove deletes the CancelFunc specified by path.
// If this CancelFunc exists, returns it and true.
// If not, returns nil and false.
func (c *canceler) remove(path string) (context.CancelFunc, bool) {
	c.mu.Lock()
	cancel, ok := c.cancelMap[path]
	if ok {
		delete(c.cancelMap, path)
		c.mu.Unlock()
		return cancel, true
	}
	c.mu.Unlock()
	return nil, false
}
