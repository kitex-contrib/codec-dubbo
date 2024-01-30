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

package sql

import (
	"time"

	"github.com/apache/dubbo-go-hessian2"
)

func init() {
	hessian.RegisterPOJO(&Timestamp{})
	hessian.SetJavaSqlTimeSerialize(&Timestamp{})
}

type Timestamp struct {
	time.Time
}

func (d *Timestamp) GetTime() time.Time {
	return d.Time
}

func (d *Timestamp) SetTime(time time.Time) {
	d.Time = time
}

func (Timestamp) JavaClassName() string {
	return "java.sql.Timestamp"
}

func (d *Timestamp) ValueOf(dateStr string) error {
	// todo(DMwangnima): change layout
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Time = date
	return nil
}

// nolint
func (d *Timestamp) Year() int {
	return d.Time.Year()
}

// nolint
func (d *Timestamp) Month() time.Month {
	return d.Time.Month()
}

// nolint
func (d *Timestamp) Day() int {
	return d.Time.Day()
}
