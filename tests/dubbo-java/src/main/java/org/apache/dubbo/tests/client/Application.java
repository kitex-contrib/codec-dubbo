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

package org.apache.dubbo.tests.client;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.apache.dubbo.tests.api.UserProvider;

public class Application {
    public static void main(String[] args) throws IOException {
        ReferenceConfig<UserProvider> reference = new ReferenceConfig<>();
        reference.setInterface(UserProvider.class);
        reference.setUrl("127.0.0.1:20000");

        DubboBootstrap.getInstance()
                .application("first-dubbo-consumer")
                .reference(reference)
                .start();

        UserProvider service = reference.get();
        testBaseTypes(service);
        testContainerListType(service);
        testContainerMapType(service);
    }

    public static void logEchoFail(String methodName) {
        System.out.printf("{%s} fail\n", methodName);
    }

    public static void logEchoException(String methodName, Exception e) {
        System.out.printf("{%s} exception: {%s}\n", methodName, e);
    }

    public static void logEchoEnd(String methodName) {
        System.out.printf("{%s} end\n", methodName);
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

    public static void testContainerListType(UserProvider svc) {
        testEchoBoolList(svc);
        testEchoByteList(svc);
        testEchoInt16List(svc);
        testEchoInt32List(svc);
        testEchoInt64List(svc);
        testEchoDoubleList(svc);
        testEchoStringList(svc);
//        testEchoBinaryList(svc);
    }

    public static void testContainerMapType(UserProvider svc) {
        testEchoBool2BoolMap(svc);
        testEchoBool2ByteMap(svc);
        testEchoBool2Int16Map(svc);
        testEchoBool2Int32Map(svc);
        testEchoBool2Int64Map(svc);
        testEchoBool2DoubleMap(svc);
        testEchoBool2StringMap(svc);
        testEchoBool2BinaryMap(svc);
    }

    public static void testEchoBool(UserProvider svc) {
        String methodName = "EchoBool";
        try {
            boolean req = true;
            boolean resp = svc.EchoBool(req);
            if (req != resp) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoByte(UserProvider svc) {
        String methodName = "EchoByte";
        try {
            Byte req = '1';
            Byte resp = svc.EchoByte(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt16(UserProvider svc) {
        String methodName = "EchoInt16";
        try {
            Short req = 12;
            Short resp = svc.EchoInt16(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt32(UserProvider svc) {
        String methodName = "EchoInt32";
        try {
            Integer req = 12;
            Integer resp = svc.EchoInt32(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt64(UserProvider svc) {
        String methodName = "EchoInt64";
        try {
            Long req = 12L;
            Long resp = svc.EchoInt64(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoDouble(UserProvider svc) {
        String methodName = "EchoDouble";
        try {
            Double req = 12.34;
            Double resp = svc.EchoDouble(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoString(UserProvider svc) {
        String methodName = "EchoString";
        try {
            String req = "12";
            String resp = svc.EchoString(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBinary(UserProvider svc) {
        String methodName = "EchoBinary";
        try {
            byte[] req = "1234".getBytes();
            byte[] resp = svc.EchoBinary(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBoolList(UserProvider svc) {
        String methodName = "EchoBoolList";
        try {
            ArrayList<Boolean> req = new ArrayList<>();
            req.add(true);
            ArrayList<Boolean> resp = svc.EchoBoolList(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoByteList(UserProvider svc) {
        String methodName = "EchoByteList";
        try {
            ArrayList<Byte> req = new ArrayList<>();
            req.add((byte)12);
            ArrayList<Byte> resp = svc.EchoByteList(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt16List(UserProvider svc) {
        String methodName = "EchoInt16List";
        try {
            ArrayList<Short> req = new ArrayList<>();
            req.add((short)12);
            ArrayList<Short> resp = svc.EchoInt16List(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt32List(UserProvider svc) {
        String methodName = "EchoInt32List";
        try {
            ArrayList<Integer> req = new ArrayList<>();
            req.add(12);
            ArrayList<Integer> resp = svc.EchoInt32List(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt64List(UserProvider svc) {
        String methodName = "EchoInt64List";
        try {
            ArrayList<Long> req = new ArrayList<>();
            req.add((long)12);
            ArrayList<Long> resp = svc.EchoInt64List(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoDoubleList(UserProvider svc) {
        String methodName = "EchoDoubleList";
        try {
            ArrayList<Double> req = new ArrayList<>();
            req.add(12.34);
            ArrayList<Double> resp = svc.EchoDoubleList(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoStringList(UserProvider svc) {
        String methodName = "EchoStringList";
        try {
            ArrayList<String> req = new ArrayList<>();
            req.add("12");
            ArrayList<String> resp = svc.EchoStringList(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBinaryList(UserProvider svc) {
        String methodName = "EchoBinaryList";
        try {
            ArrayList<byte[]> req = new ArrayList<>();
            byte[] bs = new byte[]{1, 2};
            req.add(bs);
            ArrayList<byte[]> resp = svc.EchoBinaryList(req);
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2BoolMap(UserProvider svc) {
        String methodName = "EchoBool2BoolMap";
        try {
            HashMap<Boolean, Boolean> req = new HashMap<>();
            req.put(true, true);
            HashMap<Boolean, Boolean> resp = svc.EchoBool2BoolMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2ByteMap(UserProvider svc) {
        String methodName = "EchoBool2ByteMap";
        try {
            HashMap<Boolean, Byte> req = new HashMap<>();
            req.put(true, (byte) 12);
            HashMap<Boolean, Byte> resp = svc.EchoBool2ByteMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int16Map(UserProvider svc) {
        String methodName = "EchoBool2Int16Map";
        try {
            HashMap<Boolean, Short> req = new HashMap<>();
            req.put(true, (short) 12);
            HashMap<Boolean, Short> resp = svc.EchoBool2Int16Map(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int32Map(UserProvider svc) {
        String methodName = "EchoBool2Int32Map";
        try {
            HashMap<Boolean, Integer> req = new HashMap<>();
            req.put(true, 1);
            HashMap<Boolean, Integer> resp = svc.EchoBool2Int32Map(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int64Map(UserProvider svc) {
        String methodName = "EchoBool2Int64Map";
        try {
            HashMap<Boolean, Long> req = new HashMap<>();
            req.put(true, (long)1);
            HashMap<Boolean, Long> resp = svc.EchoBool2Int64Map(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2DoubleMap(UserProvider svc) {
        String methodName = "EchoBool2DoubleMap";
        try {
            HashMap<Boolean, Double> req = new HashMap<>();
            req.put(true, 12.34);
            HashMap<Boolean, Double> resp = svc.EchoBool2DoubleMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2StringMap(UserProvider svc) {
        String methodName = "EchoBool2StringMap";
        try {
            HashMap<Boolean, String> req = new HashMap<>();
            req.put(true, "12");
            HashMap<Boolean, String> resp = svc.EchoBool2StringMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2BinaryMap(UserProvider svc) {
        String methodName = "EchoBool2BinaryMap";
        try {
            HashMap<Boolean, byte[]> req = new HashMap<>();
            byte[] bs = new byte[]{1, 2};
            req.put(true, bs);
            HashMap<Boolean, byte[]> resp = svc.EchoBool2BinaryMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }
}