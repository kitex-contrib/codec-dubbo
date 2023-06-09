#!/usr/bin/env bash
# Copyright 2023 CloudWeGo Authors
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


current=$(git status | head -n1 | sed 's/On branch //')
name=${1:-$current}
if [[ ! $name =~ ^(((opt(imize)?|feat(ure)?|(bug|hot)?fix|test|refact(or)?|ci)/.+)|(main|develop)|(release-v[0-9]+\.[0-9]+)|(release/v[0-9]+\.[0-9]+\.[0-9]+(-[a-z0-9.]+(\+[a-z0-9.]+)?)?)|revert-[a-z0-9]+)$ ]]; then
    echo "branch name '$name' is invalid"
    exit 1
else
    echo "branch name '$name' is valid"
fi
