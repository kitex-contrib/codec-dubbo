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

package org.apache.dubbo.samples.provider;

import org.apache.dubbo.samples.api.UserProvider;

import java.util.HashMap;
import java.util.List;

public class UserProviderImpl implements UserProvider {
    @Override
    public int EchoInt(int req) throws Exception {
        return 0;
    }

    @Override
    public byte EchoInt8(byte req) throws Exception {
        return 0;
    }

    @Override
    public boolean EchoBool(boolean req) throws Exception {
        return false;
    }

    @Override
    public Byte EchoByte(Byte req) throws Exception {
        return null;
    }

    @Override
    public Short EchoInt16(Short req) throws Exception {
        return null;
    }

    @Override
    public Integer EchoInt32(Integer req) throws Exception {
        return null;
    }

    @Override
    public Long EchoInt64(Long req) throws Exception {
        return null;
    }

    @Override
    public Double EchoDouble(Double req) throws Exception {
        return null;
    }

    @Override
    public String EchoString(String req) throws Exception {
        return null;
    }

    @Override
    public byte[] EchoBinary(byte[] req) throws Exception {
        return new byte[0];
    }

    @Override
    public List<Boolean> EchoBoolList(List<Boolean> req) throws Exception {
        return null;
    }

    @Override
    public List<Byte> EchoByteList(List<Byte> req) throws Exception {
        return null;
    }

    @Override
    public List<Short> EchoInt16List(List<Short> req) throws Exception {
        return null;
    }

    @Override
    public List<Integer> EchoInt32List(List<Integer> req) throws Exception {
        return null;
    }

    @Override
    public List<Long> EchoInt64List(List<Long> req) throws Exception {
        return null;
    }

    @Override
    public List<Double> EchoDoubleList(List<Double> req) throws Exception {
        return null;
    }

    @Override
    public List<String> EchoStringList(List<String> req) throws Exception {
        return null;
    }

    @Override
    public List<byte[]> EchoBinaryList(List<byte[]> req) throws Exception {
        return null;
    }

    @Override
    public HashMap<Boolean, Boolean> EchoBool2BoolMap(HashMap<Boolean, Boolean> req) throws Exception {
        return null;
    }
}
