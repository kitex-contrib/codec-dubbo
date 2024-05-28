// Copyright 2024 CloudWeGo Authors
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"html/template"

	"google.golang.org/protobuf/compiler/protogen"
)

var extTemplates []string

// RegisterTemplate add templates to the generator
func RegisterTemplate(t ...string) {
	extTemplates = append(extTemplates, t...)
}

const (
	filePrefix = "kitex_gen/"
	fileExt    = "hessian2_ext.go"
	version    = "v0.0.1"
)

type Generator interface {
	Generate(gen *protogen.Plugin) error
}

type generator struct {
	extLang string
	tpl     *template.Template
	funcs   []func(g *generator, gen *protogen.Plugin) error
}

func (g *generator) Generate(gen *protogen.Plugin) error {
	for _, f := range g.funcs {
		err := f(g, gen)
		if err != nil {
			return err
		}
	}
	return nil
}

// New implements Generator
func New(req *protogen.Plugin, lang string) (Generator, error) {
	tpl := template.New("kitex-dubbo")
	var err error
	for _, temp := range extTemplates {
		tpl, err = tpl.Parse(temp)
		if err != nil {
			return nil, err
		}
	}
	funcs := []func(g *generator, gen *protogen.Plugin) error{
		generateHessian2Ext,
	}
	return &generator{
		tpl:     tpl,
		extLang: lang,
		funcs:   funcs,
	}, nil
}
