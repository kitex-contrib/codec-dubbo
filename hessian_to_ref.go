/*
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
	"fmt"
	"strings"
	"time"
)

type JavaBean interface {
	JavaClassPackage() string

	JavaClassName() string
}

type VirtualClass struct {
	className    string
	classPackage string
	fields       map[string]interface{}
}

func NewVirtualClass(typ string, fieldNames []string) *VirtualClass {
	idx := strings.LastIndexByte(typ, '.')
	classPackage := ""
	className := typ
	if idx > 0 {
		classPackage = typ[:idx]
		className = typ[idx:]
	}
	fields := make(map[string]interface{})
	for _, key := range fieldNames {
		fields[key] = nil
	}
	return &VirtualClass{
		classPackage: classPackage,
		className:    className,
		fields:       fields,
	}
}

func (vc *VirtualClass) JavaClassPackage() string {
	return vc.classPackage
}

func (vc *VirtualClass) JavaClassName() string {
	return vc.className
}

func (vc *VirtualClass) JavaFields() map[string]interface{} {
	return vc.fields
}

func convJavaType(v interface{}) (t string, r interface{}) {
	t = "Object"
	r = fmt.Sprintf("%v", v)
	switch v.(type) {
	case string:
		t = "String"
		r = fmt.Sprintf(`"%s"`, v)
	case int, int32, uint, uint32:
		t = "Integer"
	case int8, uint8:
		t = "Byte"
	case int16, uint16:
		t = "Short"
	case int64, uint64:
		t = "Long"
		r = fmt.Sprintf("%vL", v)
	case float32:
		t = "Float"
		r = fmt.Sprintf("%vF", v)
	case float64:
		t = "Double"
	case bool:
		t = "Bool"
	case time.Time:
		t = "Date"
		r = fmt.Sprintf(`new Date("%s")`, v)
	}
	return
}

func (vc *VirtualClass) String() string {
	classMsg := "\nclass " + vc.className + " {"
	if vc.classPackage != "" {
		classMsg = "\npackage " + vc.classPackage + "" + classMsg
	}
	for key, value := range vc.JavaFields() {
		t, r := convJavaType(value)
		classMsg += fmt.Sprintf("\n\tprivate %s %s = %s;", t, key, r)
	}
	classMsg += "\n}"
	return classMsg
}
