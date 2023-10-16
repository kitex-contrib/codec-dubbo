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

package org.cloudwego.kitex.samples.client;

import java.io.IOException;

import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.cloudwego.kitex.samples.api.*;

public class Application {
    public static void main(String[] args) throws IOException {
        ReferenceConfig<GreetProvider> reference = new ReferenceConfig<>();
        reference.setInterface(GreetProvider.class);
        reference.setUrl("127.0.0.1:21000");

        DubboBootstrap.getInstance()
                .application("first-dubbo-consumer")
                .reference(reference)
                .start();

        GreetProvider service = reference.get();
        try {
            String req = "world";
            String resp = service.Greet(req);
            System.out.printf("resp: %s\n", resp);
        } catch (Exception e) {
            System.out.printf("catch exception: %s\n", e);
        }

        try {
            GreetRequest reqWithStruct = new GreetRequest("world");
            GreetResponse respWithStruct = service.GreetWithStruct(reqWithStruct);
            System.out.printf("respWithStruct: %s\n", respWithStruct.getResp());
        } catch (Exception e) {
            System.out.printf("catch exception: %s\n", e);
        }

        return;
    }
}