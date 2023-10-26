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
	anno  string
	types []string
}

// NewTypeAnnotation is used to create a type annotation object.
// This function only accepts and stores the annotation string; it does not directly parse it.
func NewTypeAnnotation(anno string) *TypeAnnotation {
	return &TypeAnnotation{anno: anno}
}

// GetFieldType retrieves the type annotation for a field by its index.
func (ta *TypeAnnotation) GetFieldType(i int) string {
	if ta != nil && len(ta.getTypes()) > i {
		return ta.getTypes()[i]
	}
	return ""
}

// getTypes is used to retrieve the list of types from the type annotation.
// This function will first return the existing type list if it exists. If not, it will parse the annotation string,
// store the parsing result, and then return it.
func (ta *TypeAnnotation) getTypes() []string {
	if ta == nil {
		return nil
	}
	if ta.types == nil {
		ta.types = strings.Split(ta.anno, ",")
	}
	return ta.types
}
