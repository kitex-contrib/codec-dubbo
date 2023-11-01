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

import "strings"

// TypeAnnotation is used to store and parse a type annotation.
type TypeAnnotation struct {
	anno       string
	fieldTypes []string
}

// NewTypeAnnotation is used to create a type annotation object.
func NewTypeAnnotation(anno string) *TypeAnnotation {
	ta := &TypeAnnotation{anno: anno}
	ta.fieldTypes = strings.Split(ta.anno, ",")
	return ta
}

// GetFieldType retrieves the type annotation for a field by its index.
func (ta *TypeAnnotation) GetFieldType(i int) string {
	if ta != nil && len(ta.fieldTypes) > i {
		return ta.fieldTypes[i]
	}
	return ""
}
