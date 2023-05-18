/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
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
	"encoding/json"
	"log"
	"strings"
)

type JavaBean interface {
	JavaClassPackage() string

	JavaClassName() string
}

// VirtualClass Java virtual class model
type VirtualClass struct {
	className    string                 // Java class name
	classPackage string                 // Java class package
	fields       map[string]interface{} // Java class field map
}

// NewVirtualClass Get a Java virtual class instance
func NewVirtualClass(typ string, fieldNames []string) *VirtualClass {
	idx := strings.LastIndexByte(typ, '.')
	classPackage := ""
	className := typ
	if idx > 0 {
		classPackage = typ[:idx]
		className = strings.Trim(typ[idx:], ".")
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

// JavaClassPackage Get Java class package name
func (vc *VirtualClass) JavaClassPackage() string {
	return vc.classPackage
}

// JavaClassName Get Java class name
func (vc *VirtualClass) JavaClassName() string {
	return vc.className
}

// JavaFields Get Java class field map
func (vc *VirtualClass) JavaFields() map[string]interface{} {
	return vc.fields
}

// JavaFields Convert Java virtual class to simple String
func (vc *VirtualClass) String() string {
	m := make(map[string]interface{})
	m["class_package"] = vc.classPackage
	m["class_name"] = vc.className
	fields := make(map[string]interface{})
	for k, v := range vc.fields {
		switch v2 := v.(type) {
		case map[interface{}]interface{}:
			tm := make(map[string]interface{})
			for mk, mv := range v2 {
				tm[mk.(string)] = mv
				switch mv2 := mv.(type) {
				case *VirtualClass:
					tm[mk.(string)] = mv2.String()
				default:
					tm[mk.(string)] = mv2.(string)
				}
			}
			v = tm
		case []interface{}:
			tl := make([]string, len(v2))
			for li, lv := range v2 {
				switch lv2 := lv.(type) {
				case *VirtualClass:
					tl[li] = lv2.String()
				default:
					tl[li] = lv2.(string)
				}
			}
			v = tl
		default:
			v = v2
		}
		fields[k] = v
	}
	m["class_fields"] = fields
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}
