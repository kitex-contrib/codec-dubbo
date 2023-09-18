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
import java.util.ArrayList;

public interface UserProvider {
    int EchoInt(int req) throws Exception;
    byte EchoInt8(byte req) throws Exception;

    boolean EchoBool(boolean req) throws Exception;

    byte EchoByte(byte req) throws Exception;

    short EchoInt16(short req) throws Exception;

    int EchoInt32(int req) throws Exception;

    long EchoInt64(long req) throws Exception;

    double EchoDouble(double req) throws Exception;

    String EchoString(String req) throws Exception;

    byte[] EchoBinary(byte[] req) throws Exception;

    ArrayList<Boolean> EchoBoolList(ArrayList<Boolean> req) throws Exception;

    ArrayList<Byte> EchoByteList(ArrayList<Byte> req) throws Exception;

    ArrayList<Short> EchoInt16List(ArrayList<Short> req) throws Exception;

    ArrayList<Integer> EchoInt32List(ArrayList<Integer> req) throws Exception;

    ArrayList<Long> EchoInt64List(ArrayList<Long> req) throws Exception;

    ArrayList<Double> EchoDoubleList(ArrayList<Double> req) throws Exception;

    ArrayList<String> EchoStringList(ArrayList<String> req) throws Exception;

    ArrayList<byte[]> EchoBinaryList(ArrayList<byte[]> req) throws Exception;


    HashMap<Boolean, Boolean> EchoBool2BoolMap(HashMap<Boolean, Boolean> req) throws Exception;

    HashMap<Boolean, Byte> EchoBool2ByteMap(HashMap<Boolean, Byte> req) throws Exception;

    HashMap<Boolean, Short> EchoBool2Int16Map(HashMap<Boolean, Short> req) throws Exception;

    HashMap<Boolean, Integer> EchoBool2Int32Map(HashMap<Boolean, Integer> req) throws Exception;

    HashMap<Boolean, Long> EchoBool2Int64Map(HashMap<Boolean, Long> req) throws Exception;

    HashMap<Boolean, Double> EchoBool2DoubleMap(HashMap<Boolean, Double> req) throws Exception;

    HashMap<Boolean, String> EchoBool2StringMap(HashMap<Boolean, String> req) throws Exception;

    HashMap<Boolean, byte[]> EchoBool2BinaryMap(HashMap<Boolean, byte[]> req) throws Exception;
}