package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiStringResponse implements Serializable {
    String baseResp;
    List<String> listResp;
    Map<String, String> mapResp;

    public EchoMultiStringResponse(String baseResp, List<String> listResp, Map<String, String> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public String getBaseResp() {
        return baseResp;
    }

    public List<String> getListResp() {
        return listResp;
    }

    public Map<String, String> getMapResp() {
        return mapResp;
    }
}
