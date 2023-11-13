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
import java.util.*;

import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.apache.dubbo.tests.api.*;
import org.eclipse.jetty.server.Authentication;

public class Application {
    public static void main(String[] args) throws IOException {
        DubboBootstrap instance = DubboBootstrap.getInstance()
                    .application("dubbo");
        ReferenceConfig<UserProvider> reference = new ReferenceConfig<>();
        reference.setInterface(UserProvider.class);

        boolean withRegistryFlag = false;
        if (args.length >= 1 && args[0].equals("withRegistry")) {
            withRegistryFlag = true;
            // initialize zookeeper registry
            String zookeeperAddress = "zookeeper://127.0.0.1:2181";
            RegistryConfig zookeeper = new RegistryConfig(zookeeperAddress);
            zookeeper.setGroup("myGroup");
            zookeeper.setRegisterMode("interface");
            instance = instance.registry(zookeeper);
        } else {
            reference.setUrl("127.0.0.1:20000");
        }

        instance.reference(reference)
                .start();
        UserProvider service = reference.get();
        if (withRegistryFlag) {
            testEchoBool(service);
            return;
        }

        testBaseTypes(service);
        testContainerListType(service);
        testContainerMapType(service);
        testMultiParams(service);
        testMethodAnnotation(service);
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
        testEchoFloat(svc);
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
//         testEchoFloatList(svc);
        testEchoDoubleList(svc);
        testEchoStringList(svc);
//        testEchoBinaryList(svc);
    }

    public static void testContainerMapType(UserProvider svc) {
        testEchoBool2BoolMap(svc);
//        testEchoBool2ByteMap(svc);
//        testEchoBool2Int16Map(svc);
        testEchoBool2Int32Map(svc);
        testEchoBool2Int64Map(svc);
//         testEchoBool2FloatMap(svc);
        testEchoBool2DoubleMap(svc);
        testEchoBool2StringMap(svc);
//        testEchoBool2BinaryMap(svc);
    }

    public static void testMultiParams(UserProvider svc) {
        testEchoMultiBool(svc);
        testEchoMultiByte(svc);
        testEchoMultiInt16(svc);
        testEchoMultiInt32(svc);
        testEchoMultiInt64(svc);
        testEchoMultiFloat(svc);
        testEchoMultiDouble(svc);
        testEchoMultiString(svc);
    }

    public static void testMethodAnnotation(UserProvider svc) {
        testEchoBaseBool(svc);
        testEchoBaseByte(svc);
        testEchoBaseInt16(svc);
        testEchoBaseInt32(svc);
        testEchoBaseInt64(svc);
        testEchoBaseFloat(svc);
        testEchoBaseDouble(svc);
        testEchoBaseBoolList(svc);
//         testEchoBaseByteList(svc);
        testEchoBaseInt16List(svc);
        testEchoBaseInt32List(svc);
        testEchoBaseInt64List(svc);
        testEchoBaseFloatList(svc);
        testEchoBaseDoubleList(svc);
        testEchoBool2BoolBaseMap(svc);
//        testEchoBool2ByteBaseMap(svc);
//        testEchoBool2Int16BaseMap(svc);
        testEchoBool2Int32BaseMap(svc);
        testEchoBool2Int64BaseMap(svc);
//         testEchoBool2FloatBaseMap(svc);
        testEchoBool2DoubleBaseMap(svc);
        testEchoMultiBaseBool(svc);
//         testEchoMultiBaseByte(svc);
        testEchoMultiBaseInt16(svc);
        testEchoMultiBaseInt32(svc);
        testEchoMultiBaseInt64(svc);
        testEchoMultiBaseFloat(svc);
        testEchoMultiBaseDouble(svc);
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

    public static void testEchoFloat(UserProvider svc) {
        String methodName = "EchoFloat";
        try {
            Float req = 12.34F;
            Float resp = svc.EchoFloat(req);
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
            List<Boolean> resp = svc.EchoBoolList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoByteList(UserProvider svc) {
        String methodName = "EchoByteList";
        try {
            ArrayList<Byte> req = new ArrayList<>();
            req.add((byte) 12);
            List<Byte> resp = svc.EchoByteList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt16List(UserProvider svc) {
        String methodName = "EchoInt16List";
        try {
            ArrayList<Short> req = new ArrayList<>();
            req.add((short) 12);
            List<Short> resp = svc.EchoInt16List(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
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
            List<Integer> resp = svc.EchoInt32List(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoInt64List(UserProvider svc) {
        String methodName = "EchoInt64List";
        try {
            ArrayList<Long> req = new ArrayList<>();
            req.add((long) 12);
            List<Long> resp = svc.EchoInt64List(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoFloatList(UserProvider svc) {
        String methodName = "EchoFloatList";
        try {
            ArrayList<Float> req = new ArrayList<>();
            req.add(12.34F);
            List<Float> resp = svc.EchoFloatList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
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
            List<Double> resp = svc.EchoDoubleList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
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
            List<String> resp = svc.EchoStringList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
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
            List<byte[]> resp = svc.EchoBinaryList(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
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
            Map<Boolean, Boolean> resp = svc.EchoBool2BoolMap(req);
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
            Map<Boolean, Byte> resp = svc.EchoBool2ByteMap(req);
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
            Map<Boolean, Short> resp = svc.EchoBool2Int16Map(req);
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
            Map<Boolean, Integer> resp = svc.EchoBool2Int32Map(req);
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
            req.put(true, (long) 1);
            Map<Boolean, Long> resp = svc.EchoBool2Int64Map(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2FloatMap(UserProvider svc) {
        String methodName = "EchoBool2FloatMap";
        try {
            HashMap<Boolean, Float> req = new HashMap<>();
            req.put(true, 12.34F);
            Map<Boolean, Float> resp = svc.EchoBool2FloatMap(req);
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
            Map<Boolean, Double> resp = svc.EchoBool2DoubleMap(req);
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
            Map<Boolean, String> resp = svc.EchoBool2StringMap(req);
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
            Map<Boolean, byte[]> resp = svc.EchoBool2BinaryMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBool(UserProvider svc) {
        String methodName = "EchoMultiBool";
        try {
            boolean baseReq = true;
            ArrayList<Boolean> listReq = new ArrayList<>();
            listReq.add(true);
            listReq.add(true);
            HashMap<Boolean, Boolean> mapReq = new HashMap<>();
            mapReq.put(true, true);
            EchoMultiBoolResponse resp = svc.EchoMultiBool(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiByte(UserProvider svc) {
        String methodName = "EchoMultiByte";
        try {
            byte baseReq = 1;
            ArrayList<Byte> listReq = new ArrayList<>();
            listReq.add((byte) 12);
            listReq.add((byte) 34);
            HashMap<Byte, Byte> mapReq = new HashMap<>();
            mapReq.put((byte) 12, (byte) 34);
            EchoMultiByteResponse resp = svc.EchoMultiByte(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiInt16(UserProvider svc) {
        String methodName = "EchoMultiInt16";
        try {
            short baseReq = 1;
            ArrayList<Short> listReq = new ArrayList<>();
            listReq.add((short) 12);
            listReq.add((short) 34);
            HashMap<Short, Short> mapReq = new HashMap<>();
            mapReq.put((short) 12, (short) 34);
            EchoMultiInt16Response resp = svc.EchoMultiInt16(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiInt32(UserProvider svc) {
        String methodName = "EchoMultiInt32";
        try {
            int baseReq = 1;
            ArrayList<Integer> listReq = new ArrayList<>();
            listReq.add(12);
            listReq.add(34);
            HashMap<Integer, Integer> mapReq = new HashMap<>();
            mapReq.put(12, 34);
            EchoMultiInt32Response resp = svc.EchoMultiInt32(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiInt64(UserProvider svc) {
        String methodName = "EchoMultiInt64";
        try {
            long baseReq = 1;
            ArrayList<Long> listReq = new ArrayList<>();
            listReq.add((long) 12);
            listReq.add((long) 34);
            HashMap<Long, Long> mapReq = new HashMap<>();
            mapReq.put((long) 12, (long) 34);
            EchoMultiInt64Response resp = svc.EchoMultiInt64(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiFloat(UserProvider svc) {
        String methodName = "EchoMultiFloat";
        try {
            float baseReq = 12.34F;
            ArrayList<Float> listReq = new ArrayList<>();
            listReq.add(12.34F);
            listReq.add(56.78F);
            HashMap<Float, Float> mapReq = new HashMap<>();
            mapReq.put(12.34F, 56.78F);
            EchoMultiFloatResponse resp = svc.EchoMultiFloat(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiDouble(UserProvider svc) {
        String methodName = "EchoMultiDouble";
        try {
            double baseReq = 12.34;
            ArrayList<Double> listReq = new ArrayList<>();
            listReq.add(12.34);
            listReq.add(56.78);
            HashMap<Double, Double> mapReq = new HashMap<>();
            mapReq.put(12.34, 56.78);
            EchoMultiDoubleResponse resp = svc.EchoMultiDouble(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiString(UserProvider svc) {
        String methodName = "EchoMultiString";
        try {
            String baseReq = "1";
            ArrayList<String> listReq = new ArrayList<>();
            listReq.add("12");
            listReq.add("34");
            HashMap<String, String> mapReq = new HashMap<>();
            mapReq.put("12", "34");
            EchoMultiStringResponse resp = svc.EchoMultiString(baseReq, listReq, mapReq);
            if (!baseReq.equals(resp.getBaseResp()) || !listReq.equals(resp.getListResp()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseBool(UserProvider svc) {
        String methodName = "EchoBaseBool";
        try {
            boolean req = true;
            boolean resp = svc.EchoBaseBool(req);
            if (req != resp) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseByte(UserProvider svc) {
        String methodName = "EchoBaseByte";
        try {
            Byte req = '1';
            Byte resp = svc.EchoBaseByte(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt16(UserProvider svc) {
        String methodName = "EchoBaseInt16";
        try {
            Short req = 12;
            Short resp = svc.EchoBaseInt16(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt32(UserProvider svc) {
        String methodName = "EchoBaseInt32";
        try {
            Integer req = 12;
            Integer resp = svc.EchoBaseInt32(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt64(UserProvider svc) {
        String methodName = "EchoBaseInt64";
        try {
            Long req = 12L;
            Long resp = svc.EchoBaseInt64(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }


    public static void testEchoBaseFloat(UserProvider svc) {
        String methodName = "EchoBaseFloat";
        try {
            Float req = 12.34F;
            Float resp = svc.EchoBaseFloat(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseDouble(UserProvider svc) {
        String methodName = "EchoBaseDouble";
        try {
            Double req = 12.34;
            Double resp = svc.EchoBaseDouble(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseBoolList(UserProvider svc) {
        String methodName = "EchoBaseBoolList";
        try {
            boolean[] req = {true, false};
            boolean[] resp = svc.EchoBaseBoolList(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseByteList(UserProvider svc) {
        String methodName = "EchoBaseByteList";
        try {
            byte[] req = {1, 2};
            byte[] resp = svc.EchoBaseByteList(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt16List(UserProvider svc) {
        String methodName = "EchoBaseInt16List";
        try {
            short[] req = {1, 2};
            short[] resp = svc.EchoBaseInt16List(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt32List(UserProvider svc) {
        String methodName = "EchoBaseInt32List";
        try {
            int[] req = {1, 2};
            int[] resp = svc.EchoBaseInt32List(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseInt64List(UserProvider svc) {
        String methodName = "EchoBaseInt64List";
        try {
            long[] req = {1, 2};
            long[] resp = svc.EchoBaseInt64List(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseFloatList(UserProvider svc) {
        String methodName = "EchoBaseFloatList";
        try {
            float[] req = {1.2F, 3.4F};
            float[] resp = svc.EchoBaseFloatList(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBaseDoubleList(UserProvider svc) {
        String methodName = "EchoBaseDoubleList";
        try {
            double[] req = {1.2, 3.4};
            double[] resp = svc.EchoBaseDoubleList(req);
            if (!Arrays.equals(req, resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2BoolBaseMap(UserProvider svc) {
        String methodName = "EchoBool2BoolBaseMap";
        try {
            HashMap<Boolean, Boolean> req = new HashMap<>();
            req.put(true, true);
            Map<Boolean, Boolean> resp = svc.EchoBool2BoolBaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2ByteBaseMap(UserProvider svc) {
        String methodName = "EchoBool2ByteBaseMap";
        try {
            HashMap<Boolean, Byte> req = new HashMap<>();
            req.put(true, (byte) 12);
            Map<Boolean, Byte> resp = svc.EchoBool2ByteBaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int16BaseMap(UserProvider svc) {
        String methodName = "EchoBool2Int16BaseMap";
        try {
            HashMap<Boolean, Short> req = new HashMap<>();
            req.put(true, (short) 12);
            Map<Boolean, Short> resp = svc.EchoBool2Int16BaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int32BaseMap(UserProvider svc) {
        String methodName = "EchoBool2Int32BaseMap";
        try {
            HashMap<Boolean, Integer> req = new HashMap<>();
            req.put(true, 1);
            Map<Boolean, Integer> resp = svc.EchoBool2Int32BaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2Int64BaseMap(UserProvider svc) {
        String methodName = "EchoBool2Int64BaseMap";
        try {
            HashMap<Boolean, Long> req = new HashMap<>();
            req.put(true, (long) 1);
            Map<Boolean, Long> resp = svc.EchoBool2Int64BaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2FloatBaseMap(UserProvider svc) {
        String methodName = "EchoBool2FloatBaseMap";
        try {
            HashMap<Boolean, Float> req = new HashMap<>();
            req.put(true, 12.34F);
            Map<Boolean, Float> resp = svc.EchoBool2FloatBaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoBool2DoubleBaseMap(UserProvider svc) {
        String methodName = "EchoBool2DoubleBaseMap";
        try {
            HashMap<Boolean, Double> req = new HashMap<>();
            req.put(true, 12.34);
            Map<Boolean, Double> resp = svc.EchoBool2DoubleBaseMap(req);
            if (!req.equals(resp)) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseBool(UserProvider svc) {
        String methodName = "EchoMultiBaseBool";
        try {
            boolean baseReq = true;
            boolean[] listReq = {true, false};
            HashMap<Boolean, Boolean> mapReq = new HashMap<>();
            mapReq.put(true, true);
            EchoMultiBoolResponse resp = svc.EchoMultiBaseBool(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseByte(UserProvider svc) {
        String methodName = "EchoMultiBaseByte";
        try {
            byte baseReq = 1;
            byte[] listReq = {12, 34};
            HashMap<Byte, Byte> mapReq = new HashMap<>();
            mapReq.put((byte) 12, (byte) 34);
            EchoMultiByteResponse resp = svc.EchoMultiBaseByte(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseInt16(UserProvider svc) {
        String methodName = "EchoMultiBaseInt16";
        try {
            short baseReq = 1;
            short[] listReq = {12, 34};
            HashMap<Short, Short> mapReq = new HashMap<>();
            mapReq.put((short) 12, (short) 34);
            EchoMultiInt16Response resp = svc.EchoMultiBaseInt16(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseInt32(UserProvider svc) {
        String methodName = "EchoMultiBaseInt32";
        try {
            int baseReq = 1;
            int[] listReq = {12, 34};
            HashMap<Integer, Integer> mapReq = new HashMap<>();
            mapReq.put(12, 34);
            EchoMultiInt32Response resp = svc.EchoMultiBaseInt32(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseInt64(UserProvider svc) {
        String methodName = "EchoMultiBaseInt64";
        try {
            long baseReq = 1;
            long[] listReq = {12, 34};
            HashMap<Long, Long> mapReq = new HashMap<>();
            mapReq.put((long) 12, (long) 34);
            EchoMultiInt64Response resp = svc.EchoMultiBaseInt64(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseFloat(UserProvider svc) {
        String methodName = "EchoMultiBaseFloat";
        try {
            float baseReq = 12.34F;
            float[] listReq = {12.34F, 56.78F};
            HashMap<Float, Float> mapReq = new HashMap<>();
            mapReq.put(12.34F, 56.78F);
            EchoMultiFloatResponse resp = svc.EchoMultiBaseFloat(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }

    public static void testEchoMultiBaseDouble(UserProvider svc) {
        String methodName = "EchoMultiBaseDouble";
        try {
            double baseReq = 12.34;
            double[] listReq = {12.34, 56.78};
            HashMap<Double, Double> mapReq = new HashMap<>();
            mapReq.put(12.34, 56.78);
            EchoMultiDoubleResponse resp = svc.EchoMultiBaseDouble(baseReq, listReq, mapReq);
            if (baseReq != resp.getBaseResp() || !Arrays.equals(listReq, resp.getListRespToArray()) || !mapReq.equals(resp.getMapResp())) {
                logEchoFail(methodName);
            }
        } catch (Exception e) {
            logEchoException(methodName, e);
        }
        logEchoEnd(methodName);
    }
}