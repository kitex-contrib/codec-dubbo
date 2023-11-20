package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoOptionalMultiStringResponse implements Serializable {
    String baseResp;
    List<String> listResp;
    Map<Boolean, String> mapResp;

    public EchoOptionalMultiStringResponse(String baseResp, List<String> listResp, Map<Boolean, String> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public String getBaseResp() {
        return baseResp;
    }

    public void setBaseResp(String baseResp) {
        this.baseResp = baseResp;
    }

    public List<String> getListResp() {
        return listResp;
    }

    public void setListResp(List<String> listResp) {
        this.listResp = listResp;
    }

    public Map<Boolean, String> getMapResp() {
        return mapResp;
    }

    public void setMapResp(Map<Boolean, String> mapResp) {
        this.mapResp = mapResp;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiStringResponse{" +
                "baseResp='" + baseResp + '\'' +
                ", listResp=" + listResp +
                ", mapResp=" + mapResp +
                '}';
    }
}