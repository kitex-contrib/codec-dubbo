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

import "google.golang.org/protobuf/compiler/protogen"

type Field struct {
	Name         string
	Type         string
	FunctionName string
}

type StructLike struct {
	Name        string
	Fields      []*Field
	JavaPackage string
	GoPackage   string
}

type ExtService struct {
	Name        string
	PkgName     string
	Functions   []*Function
	Annotations map[string]string
	Version     string
}

type Function struct {
	Method     string
	InputType  string
	OutputType string
}

type ExtFile struct {
	JavaPackage string
	GoPackage   string
	PkgName     protogen.GoPackageName
	IDLName     string
	StructLikes []*StructLike
	Services    []*ExtService
	Version     string
}
