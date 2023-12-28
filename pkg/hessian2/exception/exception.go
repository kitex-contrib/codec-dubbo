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

package exception

import (
	"github.com/apache/dubbo-go-hessian2/java_exception"
)

type Throwabler interface {
	java_exception.Throwabler
}

func NewException(detailMessage string) Throwabler {
	return java_exception.NewException(detailMessage)
}

// FromError extracts Throwabler from passed err.
//
//   - If err is nil, it returns nil and false
//
//   - If err implents Unwrap(), it would unwrap err until getting the real cause.
//     Then it would check cause whether implementing Throwabler. If yes, it returns
//     Throwabler and true.
//
//     If not, it checks err whether implementing Throwabler directly. If yes,
//     it returns Throwabler and true.
func FromError(err error) (Throwabler, bool) {
	if err == nil {
		return nil, false
	}
	for {
		if wrapper, ok := err.(interface{ Unwrap() error }); ok {
			err = wrapper.Unwrap()
		} else {
			break
		}
	}
	if exception, ok := err.(Throwabler); ok {
		return exception, true
	}
	return nil, false
}
