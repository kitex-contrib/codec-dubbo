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

package dubbo

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	hessian "github.com/apache/dubbo-go-hessian2"
)

var cache = new(typesCache)

// typesCache maintains a cache from type of data(reflect.Type) to Types string used by Hessian2.
type typesCache struct {
	group    Group
	typesMap sync.Map
}

// getByData returns the Types string of given data.
// It reads embedded sync.Map firstly. If cache misses, using singleFlight to process reflection and getParamsTypeList.
func (tc *typesCache) getByData(data interface{}) (string, error) {
	val := reflect.ValueOf(data)
	typ := val.Type()
	typesRaw, ok := tc.typesMap.Load(typ)
	if ok {
		return typesRaw.(string), nil
	}

	typesRaw, err, _ := tc.group.Do(typ, func() (interface{}, error) {
		elem := val.Elem()
		numField := elem.NumField()
		fields := make([]interface{}, numField)
		for i := 0; i < numField; i++ {
			fields[i] = elem.Field(i).Interface()
		}

		types, err := getParamsTypeList(fields)
		if err != nil {
			return "", err
		}
		tc.typesMap.Store(typ, types)

		return types, nil
	})
	if err != nil {
		return "", err
	}

	return typesRaw.(string), nil
}

// get retrieves Types string from reflect.Type directly.
// For test.
func (tc *typesCache) get(key reflect.Type) (string, bool) {
	typesRaw, ok := tc.typesMap.Load(key)
	if !ok {
		return "", false
	}

	return typesRaw.(string), true
}

// len returns the length of embedded sync.Map.
// For test.
func (tc *typesCache) len() int {
	var length int
	tc.typesMap.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

func GetTypes(data interface{}) (string, error) {
	return cache.getByData(data)
}

// GetParamsTypeList is copied from dubbo-go, it should be rewritten
func getParamsTypeList(params []interface{}) (string, error) {
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
		return typ.JavaClassName()
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
