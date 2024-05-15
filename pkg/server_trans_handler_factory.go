/*
 * Copyright 2024 CloudWeGo Authors
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
	"context"
	"errors"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	cnetpoll "github.com/cloudwego/netpoll"

	"github.com/kitex-contrib/codec-dubbo/pkg/dubbo_spec"
)

// NewSvrTransHandlerFactory the factory expand the implementation of DetectableServerTransHandler for
// hessian protocol to support the dubbo protocol probing.
func NewSvrTransHandlerFactory(opts ...Option) remote.ServerTransHandlerFactory {
	return &svrTransHandlerFactory{
		ServerTransHandlerFactory: netpoll.NewSvrTransHandlerFactory(),
		codec:                     NewDubboCodec(opts...),
	}
}

type svrTransHandlerFactory struct {
	remote.ServerTransHandlerFactory
	// the codec should be set with dubbo codec to keep consistent
	// with the function ProtocolMatch.
	codec remote.Codec
}

// NewTransHandler the wrapper of ServerTransHandlerFactory.NewTransHandler, replace the codec with dubbo codec when
// invoke the function NewTransHandler, and than restore it.
func (f *svrTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	sourceCodec := opt.Codec
	opt.Codec = f.codec
	defer func() {
		opt.Codec = sourceCodec
	}()

	handler, err := f.ServerTransHandlerFactory.NewTransHandler(opt)
	if err != nil {
		return nil, err
	}
	return &svrTransHandler{
		ServerTransHandler: handler,
	}, nil
}

type svrTransHandler struct {
	remote.ServerTransHandler
}

func (svr *svrTransHandler) ProtocolMatch(ctx context.Context, conn net.Conn) (err error) {
	// Check the validity of client preface.
	npReader := conn.(interface{ Reader() cnetpoll.Reader }).Reader()
	// read at most avoid block
	header, err := npReader.Peek(dubbo_spec.HEADER_SIZE)
	if err != nil {
		return err
	}
	if header[0] == dubbo_spec.MAGIC_HIGH && header[1] == dubbo_spec.MAGIC_LOW {
		return nil
	}
	return errors.New("error protocol not match dubbo")
}

func (svr *svrTransHandler) GracefulShutdown(ctx context.Context) error {
	if g, ok := svr.ServerTransHandler.(remote.GracefulShutdown); ok {
		g.GracefulShutdown(ctx)
	}
	return nil
}

func (svr *svrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	if s, ok := svr.ServerTransHandler.(remote.InvokeHandleFuncSetter); ok {
		s.SetInvokeHandleFunc(inkHdlFunc)
	}
}
