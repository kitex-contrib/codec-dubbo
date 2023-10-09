package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiInt32Response implements Serializable {
    int baseResp;
    List<Integer> listResp;
    Map<Integer, Integer> mapResp;

    public EchoMultiInt32Response(int baseResp, List<Integer> listResp, Map<Integer, Integer> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public int getBaseResp() {
        return baseResp;
    }

    public List<Integer> getListResp() {
        return listResp;
    }

    public Map<Integer, Integer> getMapResp() {
        return mapResp;
    }
}
