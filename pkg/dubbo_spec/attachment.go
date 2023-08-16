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

package dubbo_spec

import (
	"strconv"
	"time"
)

const (
	PATH_KEY      = "path"
	GROUP_KEY     = "group"
	INTERFACE_KEY = "interface"
	VERSION_KEY   = "version"
	TIMEOUT_KEY   = "timeout"
)

type Attachment = map[string]interface{}

func NewAttachment(path, group, iface, version string, timeout time.Duration) Attachment {
	result := Attachment{}
	if len(path) > 0 {
		result[PATH_KEY] = path
	}
	if len(group) > 0 {
		result[GROUP_KEY] = group
	}
	if len(iface) > 0 {
		result[INTERFACE_KEY] = iface
	}
	if len(version) > 0 {
		result[VERSION_KEY] = version
	}
	if timeout > 0 {
		result[TIMEOUT_KEY] = strconv.Itoa(int(timeout.Milliseconds()))
	}
	return result
}
