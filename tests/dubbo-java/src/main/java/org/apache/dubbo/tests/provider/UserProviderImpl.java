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

package org.apache.dubbo.tests.provider;

import org.apache.dubbo.tests.api.UserProvider;

import java.util.HashMap;
import java.util.List;

public class UserProviderImpl implements UserProvider {
    @Override
    public int EchoInt(int req) throws Exception {
        return req;
    }

    @Override
    public byte EchoInt8(byte req) throws Exception {
        return req;
    }

    @Override
    public boolean EchoBool(boolean req) throws Exception {
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
    public HashMap<Boolean, Boolean> EchoBool2BoolMap(HashMap<Boolean, Boolean> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Byte> EchoBool2ByteMap(HashMap<Boolean, Byte> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Short> EchoBool2Int16Map(HashMap<Boolean, Short> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Integer> EchoBool2Int32Map(HashMap<Boolean, Integer> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Long> EchoBool2Int64Map(HashMap<Boolean, Long> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, Double> EchoBool2DoubleMap(HashMap<Boolean, Double> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, String> EchoBool2StringMap(HashMap<Boolean, String> req) throws Exception {
        return req;
    }

    @Override
    public HashMap<Boolean, byte[]> EchoBool2BinaryMap(HashMap<Boolean, byte[]> req) throws Exception {
        return req;
    }
}
