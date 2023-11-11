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

package org.apache.dubbo.tests.provider;

import org.apache.dubbo.config.ProtocolConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.config.ServiceConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.apache.dubbo.tests.api.UserProvider;

import java.util.ArrayList;
import java.util.List;

public class Application {

public static void main(String[] args) {
        List<ServiceConfig> list = new ArrayList<>();
        ServiceConfig<UserProvider> service = new ServiceConfig<>();
        service.setInterface(UserProvider.class);
        service.setRef(new UserProviderImpl());
        list.add(service);

        DubboBootstrap instance = DubboBootstrap.getInstance()
                .application("first-dubbo-provider")
                .protocol(new ProtocolConfig("dubbo", 20001));


        boolean withRegistryFlag = false;
        if (args.length >= 1 && args[0].equals("withRegistry")) {
            // initialize another version of UserProvider
            ServiceConfig<UserProvider> service1 = new ServiceConfig<>();
            service1.setInterface(UserProvider.class);
            service1.setRef(new UserProviderImplV1());
            service1.setGroup("g1");
            service1.setVersion("v1");
            service1.setWeight(1000);
            list.add(service1);

            // initialize zookeeper registry
            String zookeeperAddress = "zookeeper://127.0.0.1:2181";
            RegistryConfig zookeeper = new RegistryConfig(zookeeperAddress);
            zookeeper.setGroup("myGroup");
            zookeeper.setRegisterMode("interface");

            instance = instance.registry(zookeeper);
        }

        instance.services(list).
                start().
                await();
    }
}
