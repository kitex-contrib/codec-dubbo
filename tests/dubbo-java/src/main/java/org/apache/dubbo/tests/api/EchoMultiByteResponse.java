package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiByteResponse implements Serializable {
    byte baseResp;
    List<Byte> listResp;
    Map<Byte, Byte> mapResp;

    public EchoMultiByteResponse(byte baseResp, List<Byte> listResp, Map<Byte, Byte> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public byte getBaseResp() {
        return baseResp;
    }

    public List<Byte> getListResp() {
        return listResp;
    }

    public Map<Byte, Byte> getMapResp() {
        return mapResp;
    }
}
