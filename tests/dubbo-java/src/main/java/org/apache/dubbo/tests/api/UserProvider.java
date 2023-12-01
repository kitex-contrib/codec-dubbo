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

package org.apache.dubbo.tests.api;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public interface UserProvider {
    Boolean EchoBool(Boolean req) throws Exception;

    Byte EchoByte(Byte req) throws Exception;

    Short EchoInt16(Short req) throws Exception;

    Integer EchoInt32(Integer req) throws Exception;

    Long EchoInt64(Long req) throws Exception;

    Float EchoFloat(Float req) throws Exception;

    Double EchoDouble(Double req) throws Exception;

    String EchoString(String req) throws Exception;

    byte[] EchoBinary(byte[] req) throws Exception;

    List<Boolean> EchoBoolList(List<Boolean> req) throws Exception;

    List<Byte> EchoByteList(List<Byte> req) throws Exception;

    List<Short> EchoInt16List(List<Short> req) throws Exception;

    List<Integer> EchoInt32List(List<Integer> req) throws Exception;

    List<Long> EchoInt64List(List<Long> req) throws Exception;

    List<Float> EchoFloatList(List<Float> req) throws Exception;

    List<Double> EchoDoubleList(List<Double> req) throws Exception;

    List<String> EchoStringList(List<String> req) throws Exception;

    List<byte[]> EchoBinaryList(List<byte[]> req) throws Exception;


    Map<Boolean, Boolean> EchoBool2BoolMap(Map<Boolean, Boolean> req) throws Exception;

    Map<Boolean, Byte> EchoBool2ByteMap(Map<Boolean, Byte> req) throws Exception;

    Map<Boolean, Short> EchoBool2Int16Map(Map<Boolean, Short> req) throws Exception;

    Map<Boolean, Integer> EchoBool2Int32Map(Map<Boolean, Integer> req) throws Exception;

    Map<Boolean, Long> EchoBool2Int64Map(Map<Boolean, Long> req) throws Exception;

    Map<Boolean, Float> EchoBool2FloatMap(Map<Boolean, Float> req) throws Exception;

    Map<Boolean, Double> EchoBool2DoubleMap(Map<Boolean, Double> req) throws Exception;

    Map<Boolean, String> EchoBool2StringMap(Map<Boolean, String> req) throws Exception;

    Map<Boolean, byte[]> EchoBool2BinaryMap(Map<Boolean, byte[]> req) throws Exception;

    EchoMultiBoolResponse EchoMultiBool(Boolean baseReq, List<Boolean> listReq, Map<Boolean, Boolean> mapReq) throws Exception;

    EchoMultiByteResponse EchoMultiByte(Byte baseReq, List<Byte> listReq, Map<Byte, Byte> mapReq) throws Exception;

    EchoMultiInt16Response EchoMultiInt16(Short baseReq, List<Short> listReq, Map<Short, Short> mapReq) throws Exception;

    EchoMultiInt32Response EchoMultiInt32(Integer baseReq, List<Integer> listReq, Map<Integer, Integer> mapReq) throws Exception;

    EchoMultiInt64Response EchoMultiInt64(Long baseReq, List<Long> listReq, Map<Long, Long> mapReq) throws Exception;

    EchoMultiFloatResponse EchoMultiFloat(Float baseReq, List<Float> listReq, Map<Float, Float> mapReq) throws Exception;

    EchoMultiDoubleResponse EchoMultiDouble(Double baseReq, List<Double> listReq, Map<Double, Double> mapReq) throws Exception;

    EchoMultiStringResponse EchoMultiString(String baseReq, List<String> listReq, Map<String, String> mapReq) throws Exception;

    boolean EchoBaseBool(boolean req) throws Exception;

    byte EchoBaseByte(byte req) throws Exception;

    short EchoBaseInt16(short req) throws Exception;

    int EchoBaseInt32(int req) throws Exception;

    long EchoBaseInt64(long req) throws Exception;

    float EchoBaseFloat(float req) throws Exception;

    double EchoBaseDouble(double req) throws Exception;

    boolean[] EchoBaseBoolList(boolean[] req) throws Exception;

    byte[] EchoBaseByteList(byte[] req) throws Exception;

    short[] EchoBaseInt16List(short[] req) throws Exception;

    int[] EchoBaseInt32List(int[] req) throws Exception;

    long[] EchoBaseInt64List(long[] req) throws Exception;

    float[] EchoBaseFloatList(float[] req) throws Exception;

    double[] EchoBaseDoubleList(double[] req) throws Exception;

    HashMap<Boolean, Boolean> EchoBool2BoolBaseMap(HashMap<Boolean, Boolean> req) throws Exception;

    HashMap<Boolean, Byte> EchoBool2ByteBaseMap(HashMap<Boolean, Byte> req) throws Exception;

    HashMap<Boolean, Short> EchoBool2Int16BaseMap(HashMap<Boolean, Short> req) throws Exception;

    HashMap<Boolean, Integer> EchoBool2Int32BaseMap(HashMap<Boolean, Integer> req) throws Exception;

    HashMap<Boolean, Long> EchoBool2Int64BaseMap(HashMap<Boolean, Long> req) throws Exception;

    HashMap<Boolean, Float> EchoBool2FloatBaseMap(HashMap<Boolean, Float> req) throws Exception;

    HashMap<Boolean, Double> EchoBool2DoubleBaseMap(HashMap<Boolean, Double> req) throws Exception;

    EchoMultiBoolResponse EchoMultiBaseBool(boolean baseReq, boolean[] listReq, HashMap<Boolean, Boolean> mapReq) throws Exception;

    EchoMultiByteResponse EchoMultiBaseByte(byte baseReq, byte[] listReq, HashMap<Byte, Byte> mapReq) throws Exception;

    EchoMultiInt16Response EchoMultiBaseInt16(short baseReq, short[] listReq, HashMap<Short, Short> mapReq) throws Exception;

    EchoMultiInt32Response EchoMultiBaseInt32(int baseReq, int[] listReq, HashMap<Integer, Integer> mapReq) throws Exception;

    EchoMultiInt64Response EchoMultiBaseInt64(long baseReq, long[] listReq, HashMap<Long, Long> mapReq) throws Exception;

    EchoMultiFloatResponse EchoMultiBaseFloat(float baseReq, float[] listReq, HashMap<Float, Float> mapReq) throws Exception;

    EchoMultiDoubleResponse EchoMultiBaseDouble(double baseReq, double[] listReq, HashMap<Double, Double> mapReq) throws Exception;

    String EchoMethod(Boolean req) throws Exception;

    String EchoMethod(Integer req) throws Exception;

    String EchoMethod(int req) throws Exception;

    String EchoMethod(Boolean req1, Integer req2) throws Exception;

    Boolean EchoOptionalBool(Boolean req) throws Exception;

    Integer EchoOptionalInt32(Integer req) throws Exception;

    String EchoOptionalString(String req) throws Exception;

    List<Boolean> EchoOptionalBoolList(List<Boolean> req) throws Exception;

    List<Integer> EchoOptionalInt32List(List<Integer> req) throws Exception;

    List<String> EchoOptionalStringList(List<String> req) throws Exception;

    Map<Boolean, Boolean> EchoOptionalBool2BoolMap(Map<Boolean, Boolean> req) throws Exception;

    Map<Boolean, Integer> EchoOptionalBool2Int32Map(Map<Boolean, Integer> req) throws Exception;

    Map<Boolean, String> EchoOptionalBool2StringMap(Map<Boolean, String> req) throws Exception;

    EchoOptionalStructResponse EchoOptionalStruct(EchoOptionalStructRequest req) throws Exception;

    Boolean EchoOptionalMultiBoolRequest(EchoOptionalMultiBoolRequest req) throws Exception;

    Integer EchoOptionalMultiInt32Request(EchoOptionalMultiInt32Request req) throws Exception;

    String EchoOptionalMultiStringRequest(EchoOptionalMultiStringRequest req) throws Exception;

    EchoOptionalMultiBoolResponse EchoOptionalMultiBoolResponse(Boolean req) throws Exception;

    EchoOptionalMultiInt32Response EchoOptionalMultiInt32Response(Integer req) throws Exception;

    EchoOptionalMultiStringResponse EchoOptionalMultiStringResponse(String req) throws Exception;
}