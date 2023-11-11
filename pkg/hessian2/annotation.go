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

package hessian2

import (
	"strings"
)

// MethodAnnotation Used to store parameter types and parameter names in method annotations.
type MethodAnnotation struct {
	argsAnno   string
	methodName string
	fieldTypes []string
}

// NewMethodAnnotation is used to create a method annotation object.
func NewMethodAnnotation(annos map[string][]string) *MethodAnnotation {
	ma := new(MethodAnnotation)
	if v, ok := annos[HESSIAN_ARGS_TYPE_TAG]; ok && len(v) > 0 {
		ma.argsAnno = v[0]
		ma.fieldTypes = strings.Split(ma.argsAnno, ",")
	}
	if v, ok := annos[HESSIAN_JAVA_METHOD_NAME_TAG]; ok && len(v) > 0 {
		ma.methodName = v[0]
	}
	return ma
}

// GetFieldType retrieves the type annotation for a field by its index.
func (ma *MethodAnnotation) GetFieldType(i int) string {
	if ma != nil && len(ma.fieldTypes) > i {
		return ma.fieldTypes[i]
	}
	return ""
}

// GetMethodName get the method name specified by the method annotation.
func (ma *MethodAnnotation) GetMethodName() (string, bool) {
	if ma == nil || ma.methodName == "" {
		return "", false
	}
	return ma.methodName, true
}
