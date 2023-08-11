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
	"strings"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"
)

// GetParamsTypeList is copied from dubbo-go, it should be rewritten
func GetParamsTypeList(params []interface{}) (string, error) {
	var (
		typ   string
		types string
	)

	for i := range params {
		typ = getParamType(params[i])
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

func getParamType(param interface{}) string {
	if param == nil {
		return "V"
	}

	switch typ := param.(type) {
	// Serialized tags for base types
	case nil:
		return "V"
	case bool:
		return "Z"
	case []bool:
		return "[Z"
	case byte:
		return "B"
	case []byte:
		return "[B"
	case int8:
		return "B"
	case []int8:
		return "[B"
	case int16:
		return "S"
	case []int16:
		return "[S"
	case uint16: // Equivalent to Char of Java
		return "C"
	case []uint16:
		return "[C"
	// case rune:
	//	return "C"
	case int:
		return "J"
	case []int:
		return "[J"
	case int32:
		return "I"
	case []int32:
		return "[I"
	case int64:
		return "J"
	case []int64:
		return "[J"
	case time.Time:
		return "java.util.Date"
	case []time.Time:
		return "[Ljava.util.Date"
	case float32:
		return "F"
	case []float32:
		return "[F"
	case float64:
		return "D"
	case []float64:
		return "[D"
	case string:
		return "java.lang.String"
	case []string:
		return "[Ljava.lang.String;"
	case []hessian.Object:
		return "[Ljava.lang.Object;"
	case map[interface{}]interface{}:
		// return  "java.util.HashMap"
		return "java.util.Map"
	case hessian.POJOEnum:
		return typ.(hessian.POJOEnum).JavaClassName()
	case *int8:
		return "java.lang.Byte"
	case *int16:
		return "java.lang.Short"
	case *uint16:
		return "java.lang.Character"
	case *int:
		return "java.lang.Long"
	case *int32:
		return "java.lang.Integer"
	case *int64:
		return "java.lang.Long"
	case *float32:
		return "java.lang.Float"
	case *float64:
		return "java.lang.Double"
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
