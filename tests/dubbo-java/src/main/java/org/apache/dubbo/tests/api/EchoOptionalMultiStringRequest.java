package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoOptionalMultiStringRequest implements Serializable {
    String baseReq;
    List<String> listReq;
    Map<Boolean, String> mapReq;

    public EchoOptionalMultiStringRequest(String baseReq, List<String> listReq, Map<Boolean, String> mapReq) {
        this.baseReq = baseReq;
        this.listReq = listReq;
        this.mapReq = mapReq;
    }

    public String getBaseReq() {
        return baseReq;
    }

    public void setBaseReq(String baseReq) {
        this.baseReq = baseReq;
    }

    public List<String> getListReq() {
        return listReq;
    }

    public void setListReq(List<String> listReq) {
        this.listReq = listReq;
    }

    public Map<Boolean, String> getMapReq() {
        return mapReq;
    }

    public void setMapReq(Map<Boolean, String> mapReq) {
        this.mapReq = mapReq;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiStringRequest{" +
                "baseReq='" + baseReq + '\'' +
                ", listReq=" + listReq +
                ", mapReq=" + mapReq +
                '}';
    }
}
