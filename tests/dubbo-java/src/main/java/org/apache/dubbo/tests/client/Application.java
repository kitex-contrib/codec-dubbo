/*
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

package org.apache.dubbo.samples.client;

import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;

import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.apache.dubbo.samples.api.UserProvider;

public class Application {
    private static final String ZOOKEEPER_HOST = System.getProperty("zookeeper.address", "127.0.0.1");
    private static final String ZOOKEEPER_PORT = System.getProperty("zookeeper.port", "2181");
    private static final String ZOOKEEPER_ADDRESS = "zookeeper://" + ZOOKEEPER_HOST + ":" + ZOOKEEPER_PORT;

    public static void main(String[] args) throws IOException {
        ReferenceConfig<UserProvider> reference = new ReferenceConfig<>();
        reference.setInterface(UserProvider.class);
        reference.setUrl("127.0.0.1:20000");

        DubboBootstrap.getInstance()
                .application("first-dubbo-consumer")
//                .registry(new RegistryConfig(ZOOKEEPER_ADDRESS))
                .reference(reference)
                .start();

        UserProvider service = reference.get();
        testBaseTypes(service);
    }

    public static void testBaseTypes(UserProvider svc) {
        testEchoBool(svc);
        testEchoByte(svc);
        testEchoInt16(svc);
        testEchoInt32(svc);
        testEchoInt64(svc);
        testEchoDouble(svc);
        testEchoString(svc);
        testEchoBinary(svc);
    }

    public static void testEchoBool(UserProvider svc) {
        String methodName = "EchoBool";
        try {
            boolean req = true;
            boolean resp = svc.EchoBool(req);
            if (req != resp) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoByte(UserProvider svc) {
        String methodName = "EchoByte";
        try {
            Byte req = '1';
            Byte resp = svc.EchoByte(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoInt16(UserProvider svc) {
        String methodName = "EchoInt16";
        try {
            Short req = 12;
            Short resp = svc.EchoInt16(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoInt32(UserProvider svc) {
        String methodName = "EchoInt32";
        try {
            Integer req = 12;
            Integer resp = svc.EchoInt32(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoInt64(UserProvider svc) {
        String methodName = "EchoInt64";
        try {
            Long req = 12L;
            Long resp = svc.EchoInt64(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoDouble(UserProvider svc) {
        String methodName = "EchoDouble";
        try {
            Double req = 12.34;
            Double resp = svc.EchoDouble(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoString(UserProvider svc) {
        String methodName = "EchoString";
        try {
            String req = "12";
            String resp = svc.EchoString(req);
            if (!req.equals(resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

    public static void testEchoBinary(UserProvider svc) {
        String methodName = "EchoBinary";
        try {
            byte[] req = "1234".getBytes();
            byte[] resp = svc.EchoBinary(req);
            if (!Arrays.equals(req, resp)) {
                System.out.printf("%s req %s is not equal to resp %s", methodName, req, resp);
            }
        } catch (Exception e) {
            System.out.printf("%s received Exception: %s", methodName, e);
        }
    }

}
