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
	"fmt"
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
)

type Options struct {
	JavaClassName     string
	MethodAnnotations map[string]*hessian2.MethodAnnotation
	// store method name mapping of java -> go.
	// use the annotation method name + parameter types as the unique identifier.
	MethodNames map[string]string
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

// WithFileDescriptor provides method annotations for DubboCodec.
// Adding method annotations allows you to specify the method parameters and method name on the Java side.
func WithFileDescriptor(fd *thrift_reflection.FileDescriptor) Option {
	if fd == nil {
		panic("Please pass in a valid FileDescriptor.")
	}

	return Option{F: func(o *Options) {
		parseAnnotations(o, fd)
	}}
}

// parseAnnotations parse method annotations and store them in options.
func parseAnnotations(o *Options, fd *thrift_reflection.FileDescriptor) {
	o.MethodAnnotations = make(map[string]*hessian2.MethodAnnotation)
	o.MethodNames = make(map[string]string)

	for _, svc := range fd.GetServices() {
		prefix := svc.GetName() + "."

		for _, m := range svc.GetMethods() {
			ma := hessian2.NewMethodAnnotation(m.GetAnnotations())
			o.MethodAnnotations[prefix+m.GetName()] = ma
			params := getMethodParams(m, ma)

			if method, exists := ma.GetMethodName(); exists {
				types, err := hessian2.GetParamsTypeList(params)
				if err != nil {
					panic(fmt.Sprintf("Get method %s parameter types failed: %s", m.GetName(), err.Error()))
				}
				o.MethodNames[method+types] = m.GetName()
			}
		}
	}
}

// getMethodParams get the parameter list of a method.
func getMethodParams(m *thrift_reflection.MethodDescriptor, ma *hessian2.MethodAnnotation) []*hessian2.Parameter {
	params := make([]*hessian2.Parameter, len(m.GetArgs()))
	for i, a := range m.GetArgs() {
		typ, err := a.GetGoType()
		if err != nil {
			panic(fmt.Sprintf("obtain the type of parameter %s in method %s failed: %s", a.GetName(), m.GetName(), err.Error()))
		}
		val := reflect.New(typ).Elem().Interface()
		params[i] = hessian2.NewParameter(val, ma.GetFieldType(i))
	}
	return params
}
