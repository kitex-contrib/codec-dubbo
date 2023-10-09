package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiInt16Response implements Serializable {
    short baseResp;
    List<Short> listResp;
    Map<Short, Short> mapResp;

    public EchoMultiInt16Response(short baseResp, List<Short> listResp, Map<Short, Short> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public short getBaseResp() {
        return baseResp;
    }

    public List<Short> getListResp() {
        return listResp;
    }

    public Map<Short, Short> getMapResp() {
        return mapResp;
    }
}
