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
 *
 * This source file has been replicated from the original dubbo-go project
 * repository, and we extend our sincere appreciation to the dubbo-go
 * development team for their valuable contribution.
 */

package hessian2

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"
)

// MethodCache maintains a cache from method parameter types (reflect.Type) and method annotations to the type strings used by Hessian2.
type MethodCache struct {
	group    Group
	typesMap sync.Map
}

type methodKey struct {
	typ  reflect.Type
	anno string
}

// GetTypes returns the Types string for the given method parameter data and method annotations.
// It reads embedded sync.Map firstly. If cache misses, using singleFlight to process reflection and getParamsTypeList.
func (mc *MethodCache) GetTypes(data interface{}, ta *TypeAnnotation) (string, error) {
	val := reflect.ValueOf(data)
	key := methodKey{typ: val.Type()}
	if ta != nil {
		key.anno = ta.anno
	}
	typesRaw, ok := mc.typesMap.Load(key)
	if ok {
		return typesRaw.(string), nil
	}

	typesRaw, err, _ := mc.group.Do(key, func() (interface{}, error) {
		elem := val.Elem()
		numField := elem.NumField()
		fields := make([]*parameter, numField)
		for i := 0; i < numField; i++ {
			fields[i] = &parameter{
				value: elem.Field(i).Interface(),
			}
			if ta != nil {
				fields[i].typeAnno = ta.GetFieldType(i)
			}
		}

		types, err := getParamsTypeList(fields)
		if err != nil {
			return "", err
		}
		mc.typesMap.Store(key, types)

		return types, nil
	})
	if err != nil {
		return "", err
	}

	return typesRaw.(string), nil
}

// get retrieves Types string from reflect.Type directly.
// For test.
func (mc *MethodCache) get(key methodKey) (string, bool) {
	typesRaw, ok := mc.typesMap.Load(key)
	if !ok {
		return "", false
	}

	return typesRaw.(string), true
}

// len returns the length of embedded sync.Map.
// For test.
func (mc *MethodCache) len() int {
	var length int
	mc.typesMap.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

// GetParamsTypeList is copied from dubbo-go, it should be rewritten
func getParamsTypeList(params []*parameter) (string, error) {
	var (
		typ   string
		types string
	)

	for i := range params {
		typ = params[i].getType()
		if typ == "" {
			return types, fmt.Errorf("cat not get arg %#v type", params[i])
		}
		if !strings.Contains(typ, ".") {
			types += typ
		} else if strings.Index(typ, "[") == 0 {
			types += strings.Replace(typ, ".", "/", -1)
		} else {
			// java.util.List -> Ljava/util/List;
			types += "L" + strings.Replace(typ, ".", "/", -1) + ";"
		}
	}

	return types, nil
}

// parameter is used to store information about parameters.
// value stores the actual value of the parameter, and typeAnno records the type annotation added by IDL to this parameter.
type parameter struct {
	value    interface{}
	typeAnno string
}

// getType retrieves the parameter's type either through type annotation or by reflecting on the value.
func (p *parameter) getType() string {
	if p == nil {
		return "V"
	}

	// Preferentially use the type specified in the type annotation.
	if ta := p.getTypeByAnno(); len(ta) > 0 {
		return ta
	}

	return p.getTypeByValue()
}

func (p *parameter) getTypeByAnno() string {
	switch p.typeAnno {
	// When the annotation is "-", it will be skipped,
	// use the default parsing method without annotations.
	case "-":
		return ""
	case "byte":
		return "B"
	case "byte[]":
		return "[B"
	case "Byte":
		return "java.lang.Byte"
	case "Byte[]":
		return "[Ljava.lang.Byte;"
	case "short":
		return "S"
	case "short[]":
		return "[S"
	case "Short":
		return "java.lang.Short"
	case "Short[]":
		return "[Ljava.lang.Short;"
	case "int":
		return "I"
	case "int[]":
		return "[I"
	case "Integer":
		return "java.lang.Integer"
	case "Integer[]":
		return "[Ljava.lang.Integer;"
	case "long":
		return "J"
	case "long[]":
		return "[J"
	case "Long":
		return "java.lang.Long"
	case "Long[]":
		return "[Ljava.lang.Long;"
	case "float":
		return "F"
	case "float[]":
		return "[F"
	case "Float":
		return "java.lang.Float"
	case "Float[]":
		return "[Ljava.lang.Float;"
	case "double":
		return "D"
	case "double[]":
		return "[D"
	case "Double":
		return "java.lang.Double"
	case "Double[]":
		return "[Ljava.lang.Double;"
	case "boolean":
		return "Z"
	case "boolean[]":
		return "[Z"
	case "Boolean":
		return "java.lang.Boolean"
	case "Boolean[]":
		return "[Ljava.lang.Boolean;"
	case "char":
		return "C"
	case "char[]":
		return "[C"
	case "Character":
		return "java.lang.Character"
	case "Character[]":
		return "[Ljava.lang.Character;"
	case "String":
		return "java.lang.String"
	case "String[]":
		return "[Ljava.lang.String;"
	case "Object":
		return "java.lang.Object"
	case "Object[]":
		return "[Ljava.lang.Object;"
	default:
		if strings.HasSuffix(p.typeAnno, "[]") {
			return "[L" + p.typeAnno[:len(p.typeAnno)-2] + ";"
		}
		return p.typeAnno
	}
}

func (p *parameter) getTypeByValue() string {
	if p.value == nil {
		return "V"
	}

	switch typ := p.value.(type) {
	// Serialized tags for base types
	case nil:
		return "V"
	case bool:
		return "java.lang.Boolean"
	case int8:
		return "java.lang.Byte"
	case int16:
		return "java.lang.Short"
	case int32:
		return "java.lang.Integer"
	case int64:
		return "java.lang.Long"
	case float64:
		return "java.lang.Double"
	case []byte:
		return "[B"
	case time.Time:
		return "java.util.Date"
	case []time.Time:
		return "[Ljava.util.Date"
	case string:
		return "java.lang.String"
	case []hessian.Object:
		return "[Ljava.lang.Object;"
	case map[interface{}]interface{}:
		// return  "java.util.HashMap"
		return "java.util.Map"
	case hessian.POJOEnum:
		return typ.JavaClassName()
	//  Serialized tags for complex types
	default:
		reflectTyp := reflect.TypeOf(typ)
		if reflect.Ptr == reflectTyp.Kind() {
			reflectTyp = reflect.TypeOf(reflect.ValueOf(typ).Elem())
		}
		switch reflectTyp.Kind() {
		case reflect.Struct:
			hessianParam, ok := typ.(hessian.Param)
			if ok {
				return hessianParam.JavaParamName()
			}
			hessianPojo, ok := typ.(hessian.POJO)
			if ok {
				return hessianPojo.JavaClassName()
			}
			return "java.lang.Object"
		case reflect.Slice, reflect.Array:
			if reflectTyp.Elem().Kind() == reflect.Struct {
				return "[Ljava.lang.Object;"
			}
			// return "java.util.ArrayList"
			return "java.util.List"
		case reflect.Map: // Enter here, map may be map[string]int
			return "java.util.Map"
		default:
			return ""
		}
	}
}
