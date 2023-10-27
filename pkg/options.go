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
	"github.com/cloudwego/thriftgo/thrift_reflection"
	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
)

type Options struct {
	JavaClassName   string
	TypeAnnotations map[string]*hessian2.TypeAnnotation
}

func (o *Options) Apply(opts []Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

func newOptions(opts []Option) *Options {
	o := &Options{}

	o.Apply(opts)
	if o.JavaClassName == "" {
		panic("DubboCodec must be initialized with JavaClassName. Please use dubbo.WithJavaClassName().")
	}
	return o
}

type Option struct {
	F func(o *Options)
}

// WithJavaClassName configures InterfaceName for server-side service and specifies target InterfaceName for client.
// Each client and service must have its own corresponding DubboCodec.
func WithJavaClassName(name string) Option {
	return Option{F: func(o *Options) {
		o.JavaClassName = name
	}}
}

// WithFileDescriptor provides method annotations for DubboCodec. Adding method
// annotations allows specifying the Java types for DubboCodec encoding.
func WithFileDescriptor(fd *thrift_reflection.FileDescriptor) Option {
	return Option{F: func(o *Options) {
		o.TypeAnnotations = extractAnnotations(fd)
	}}
}

// extractAnnotations extracts method annotations from the given FileDescriptor
// and returns them as a map. These annotations allow specifying Java types for DubboCodec encoding.
// The annotation format is (hessian.argsType="arg1_type,arg2_type,arg3_type,..."),
// use an empty string or "-" as arg_type to use the default parsing method.
func extractAnnotations(fd *thrift_reflection.FileDescriptor) map[string]*hessian2.TypeAnnotation {
	if fd == nil {
		return nil
	}

	annotations := make(map[string]*hessian2.TypeAnnotation)

	for _, svc := range fd.GetServices() {
		prefix := svc.GetName() + "."

		for _, m := range svc.GetMethods() {
			annos := m.GetAnnotations()
			if v, ok := annos[hessian2.HESSIAN_ARGS_TYPE_TAG]; ok && len(v) > 0 {
				annotations[prefix+m.GetName()] = hessian2.NewTypeAnnotation(v[0])
			}
		}
	}
	return annotations
}
