/*
 * Copyright 2024 CloudWeGo Authors
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

package org.cloudwego.kitex.samples.enumeration;

import java.io.Serializable;

public enum KitexEnum implements Serializable {

    ONE("1"),TWO("2"),THREE("3"),FOUR("4"),FIVE("5");

    final String codeStr ;

    KitexEnum(String number) {
        this.codeStr = number;
    }

    // 枚举类型的 getter 方法
    public String getCode() {
        return this.codeStr;
    }
}
