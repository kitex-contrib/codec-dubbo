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

package java

import (
	"time"

	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
	hessian2_math "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/math"
)

type Object = interface{}

func NewObject() *Object {
	return new(Object)
}

type Date = time.Time

func NewDate() *Date {
	return new(Date)
}

type Exception = hessian2_exception.Exception

func NewException(detailMessage string) *Exception {
	return hessian2_exception.NewException(detailMessage)
}

type BigDecimal = hessian2_math.BigType

func NewBigDecimal(str string) (hessian2_math.BigType, error) {
	return hessian2_math.NewBigDecimalFromString(str)
}

type BigInteger = hessian2_math.BigType

func NewBigInteger(str string) (hessian2_math.BigType, error) {
	return hessian2_math.NewBigIntegerFromString(str)
}
