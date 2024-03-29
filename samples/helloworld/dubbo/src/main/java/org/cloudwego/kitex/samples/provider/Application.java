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

package org.cloudwego.kitex.samples.provider;

import org.apache.dubbo.config.ProtocolConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.config.ServiceConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.cloudwego.kitex.samples.api.GreetProvider;

public class Application {

    public static void main(String[] args) {
        ServiceConfig<GreetProvider> service = new ServiceConfig<>();
        service.setInterface(GreetProvider.class);
        service.setRef(new GreetProviderImpl());

        DubboBootstrap.getInstance()
                .application("first-dubbo-provider")
                .protocol(new ProtocolConfig("dubbo", 21001))
                .service(service)
                .start()
                .await();
    }
}