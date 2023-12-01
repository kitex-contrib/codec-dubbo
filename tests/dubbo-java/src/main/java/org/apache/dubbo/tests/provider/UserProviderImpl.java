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

import org.apache.dubbo.tests.api.*;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class UserProviderImpl implements UserProvider {
    @Override
    public Boolean EchoBool(Boolean req) throws Exception {
        return req;
    }

    @Override
    public Byte EchoByte(Byte req) throws Exception {
        return req;
    }

    @Override
    public Short EchoInt16(Short req) throws Exception {
        return req;
    }

    @Override
    public Integer EchoInt32(Integer req) throws Exception {
        return req;
    }

    @Override
    public Long EchoInt64(Long req) throws Exception {
        return req;
    }

    @Override
    public Float EchoFloat(Float req) throws Exception {
        return req;
    }

    @Override
    public Double EchoDouble(Double req) throws Exception {
        return req;
    }

    @Override
    public String EchoString(String req) throws Exception {
        return req;
    }

    @Override
    public byte[] EchoBinary(byte[] req) throws Exception {
        return req;
    }

    @Override
    public List<Boolean> EchoBoolList(List<Boolean> req) throws Exception {
        return req;
    }

    @Override
    public List<Byte> EchoByteList(List<Byte> req) throws Exception {
        return req;
    }

    @Override
    public List<Short> EchoInt16List(List<Short> req) throws Exception {
        return req;
    }

    @Override
    public List<Integer> EchoInt32List(List<Integer> req) throws Exception {
        return req;
    }

    @Override
    public List<Long> EchoInt64List(List<Long> req) throws Exception {
        return req;
    }

    @Override
    public List<Float> EchoFloatList(List<Float> req) throws Exception {
        return req;
    }

    @Override
    public List<Double> EchoDoubleList(List<Double> req) throws Exception {
        return req;
    }

    @Override
    public List<String> EchoStringList(List<String> req) throws Exception {
        return req;
    }

    @Override
    public List<byte[]> EchoBinaryList(List<byte[]> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Boolean> EchoBool2BoolMap(Map<Boolean, Boolean> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Byte> EchoBool2ByteMap(Map<Boolean, Byte> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Short> EchoBool2Int16Map(Map<Boolean, Short> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Integer> EchoBool2Int32Map(Map<Boolean, Integer> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Long> EchoBool2Int64Map(Map<Boolean, Long> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Float> EchoBool2FloatMap(Map<Boolean, Float> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Double> EchoBool2DoubleMap(Map<Boolean, Double> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, String> EchoBool2StringMap(Map<Boolean, String> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, byte[]> EchoBool2BinaryMap(Map<Boolean, byte[]> req) throws Exception {
        return req;
    }

    @Override
    public EchoMultiBoolResponse EchoMultiBool(Boolean baseReq, List<Boolean> listReq, Map<Boolean, Boolean> mapReq) throws Exception {
        return new EchoMultiBoolResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiByteResponse EchoMultiByte(Byte baseReq, List<Byte> listReq, Map<Byte, Byte> mapReq) throws Exception {
        return new EchoMultiByteResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt16Response EchoMultiInt16(Short baseReq, List<Short> listReq, Map<Short, Short> mapReq) throws Exception {
        return new EchoMultiInt16Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt32Response EchoMultiInt32(Integer baseReq, List<Integer> listReq, Map<Integer, Integer> mapReq) throws Exception {
        return new EchoMultiInt32Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt64Response EchoMultiInt64(Long baseReq, List<Long> listReq, Map<Long, Long> mapReq) throws Exception {
        return new EchoMultiInt64Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiFloatResponse EchoMultiFloat(Float baseReq, List<Float> listReq, Map<Float, Float> mapReq) throws Exception {
        return new EchoMultiFloatResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiDoubleResponse EchoMultiDouble(Double baseReq, List<Double> listReq, Map<Double, Double> mapReq) throws Exception {
        return new EchoMultiDoubleResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiStringResponse EchoMultiString(String baseReq, List<String> listReq, Map<String, String> mapReq) throws Exception {
        return new EchoMultiStringResponse(baseReq, listReq, mapReq);
    }


    @Override
    public boolean EchoBaseBool(boolean req) throws Exception {
        return req;
    }

    @Override
    public byte EchoBaseByte(byte req) throws Exception {
        return req;
    }

    @Override
    public short EchoBaseInt16(short req) throws Exception {
        return req;
    }

    @Override
    public int EchoBaseInt32(int req) throws Exception {
        return req;
    }

    @Override
    public long EchoBaseInt64(long req) throws Exception {
        return req;
    }

    @Override
    public float EchoBaseFloat(float req) throws Exception {
        return req;
    }

    @Override
    public double EchoBaseDouble(double req) throws Exception {
        return req;
    }

    @Override
    public boolean[] EchoBaseBoolList(boolean[] req) throws Exception {
        return req;
    }

    @Override
    public byte[] EchoBaseByteList(byte[] req) throws Exception {
        return req;
    }

    @Override
    public short[] EchoBaseInt16List(short[] req) throws Exception {
        return req;
    }

    @Override
    public int[] EchoBaseInt32List(int[] req) throws Exception {
        return req;
    }

    @Override
    public long[] EchoBaseInt64List(long[] req) throws Exception {
        return req;
    }

    @Override
    public float[] EchoBaseFloatList(float[] req) throws Exception {
        return req;
    }

    @Override
    public double[] EchoBaseDoubleList(double[] req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Boolean> EchoBool2BoolBaseMap(HashMap<Boolean, Boolean> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Byte> EchoBool2ByteBaseMap(HashMap<Boolean, Byte> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Short> EchoBool2Int16BaseMap(HashMap<Boolean, Short> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Integer> EchoBool2Int32BaseMap(HashMap<Boolean, Integer> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Long> EchoBool2Int64BaseMap(HashMap<Boolean, Long> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Float> EchoBool2FloatBaseMap(HashMap<Boolean, Float> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Double> EchoBool2DoubleBaseMap(HashMap<Boolean, Double> req) throws Exception {
        return req;
    }

    @Override
    public EchoMultiBoolResponse EchoMultiBaseBool(boolean baseReq, boolean[] listReq, HashMap<Boolean, Boolean> mapReq) throws Exception {
        ArrayList<Boolean> arr = new ArrayList<>();
        for (boolean b : listReq) arr.add(b);
        return new EchoMultiBoolResponse(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiByteResponse EchoMultiBaseByte(byte baseReq, byte[] listReq, HashMap<Byte, Byte> mapReq) throws Exception {
        ArrayList<Byte> arr = new ArrayList<>();
        for (byte b : listReq) arr.add(b);
        return new EchoMultiByteResponse(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiInt16Response EchoMultiBaseInt16(short baseReq, short[] listReq, HashMap<Short, Short> mapReq) throws Exception {
        ArrayList<Short> arr = new ArrayList<>();
        for (short s : listReq) arr.add(s);
        return new EchoMultiInt16Response(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiInt32Response EchoMultiBaseInt32(int baseReq, int[] listReq, HashMap<Integer, Integer> mapReq) throws Exception {
        ArrayList<Integer> arr = new ArrayList<>();
        for (int i : listReq) arr.add(i);
        return new EchoMultiInt32Response(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiInt64Response EchoMultiBaseInt64(long baseReq, long[] listReq, HashMap<Long, Long> mapReq) throws Exception {
        ArrayList<Long> arr = new ArrayList<>();
        for (long l : listReq) arr.add(l);
        return new EchoMultiInt64Response(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiFloatResponse EchoMultiBaseFloat(float baseReq, float[] listReq, HashMap<Float, Float> mapReq) throws Exception {
        ArrayList<Float> arr = new ArrayList<>();
        for (float d : listReq) arr.add(d);
        return new EchoMultiFloatResponse(baseReq, arr, mapReq);
    }

    @Override
    public EchoMultiDoubleResponse EchoMultiBaseDouble(double baseReq, double[] listReq, HashMap<Double, Double> mapReq) throws Exception {
        ArrayList<Double> arr = new ArrayList<>();
        for (double d : listReq) arr.add(d);
        return new EchoMultiDoubleResponse(baseReq, arr, mapReq);
    }

    @Override
    public String EchoMethod(Boolean req) throws Exception {
        return String.format("A:%b", req);
    }

    @Override
    public String EchoMethod(Integer req) throws Exception {
        return String.format("B:%d", req);
    }

    @Override
    public String EchoMethod(int req) throws Exception {
        return String.format("C:%d", req);
    }

    @Override
    public String EchoMethod(Boolean req1, Integer req2) throws Exception {
        return String.format("D:%b,%d", req1, req2);
    }

    @Override
    public Boolean EchoOptionalBool(Boolean req) throws Exception {
        return null;
    }

    @Override
    public Integer EchoOptionalInt32(Integer req) throws Exception {
        return null;
    }

    @Override
    public String EchoOptionalString(String req) throws Exception {
        return null;
    }

    @Override
    public List<Boolean> EchoOptionalBoolList(List<Boolean> req) throws Exception {
        return null;
    }

    @Override
    public List<Integer> EchoOptionalInt32List(List<Integer> req) throws Exception {
        return null;
    }

    @Override
    public List<String> EchoOptionalStringList(List<String> req) throws Exception {
        return null;
    }

    @Override
    public Map<Boolean, Boolean> EchoOptionalBool2BoolMap(Map<Boolean, Boolean> req) throws Exception {
        return null;
    }

    @Override
    public Map<Boolean, Integer> EchoOptionalBool2Int32Map(Map<Boolean, Integer> req) throws Exception {
        return null;
    }

    @Override
    public Map<Boolean, String> EchoOptionalBool2StringMap(Map<Boolean, String> req) throws Exception {
        return null;
    }

    @Override
    public EchoOptionalStructResponse EchoOptionalStruct(EchoOptionalStructRequest req) throws Exception {
        return null;
    }

    @Override
    public Boolean EchoOptionalMultiBoolRequest(EchoOptionalMultiBoolRequest req) throws Exception {
        System.out.println(req);
        return req.getBasicReq();
    }

    @Override
    public Integer EchoOptionalMultiInt32Request(EchoOptionalMultiInt32Request req) throws Exception {
        System.out.println(req);
        return req.getBasicReq();
    }

    @Override
    public String EchoOptionalMultiStringRequest(EchoOptionalMultiStringRequest req) throws Exception {
        System.out.println(req);
        return req.getBaseReq();
    }

    @Override
    public EchoOptionalMultiBoolResponse EchoOptionalMultiBoolResponse(Boolean req) throws Exception {
        return new EchoOptionalMultiBoolResponse(false, null, null, null);
    }

    @Override
    public EchoOptionalMultiInt32Response EchoOptionalMultiInt32Response(Integer req) throws Exception {
        return new EchoOptionalMultiInt32Response(0, null, null, null);
    }

    @Override
    public EchoOptionalMultiStringResponse EchoOptionalMultiStringResponse(String req) throws Exception {
        return new EchoOptionalMultiStringResponse(null, null, null);
    }
}