package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiBoolResponse implements Serializable {
    boolean baseResp;
    List<Boolean> listResp;
    Map<Boolean, Boolean> mapResp;

    public EchoMultiBoolResponse(boolean baseResp, List<Boolean> listResp, Map<Boolean, Boolean> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public boolean getBaseResp() {
        return baseResp;
    }

    public List<Boolean> getListResp() {
        return listResp;
    }

    public Map<Boolean, Boolean> getMapResp() {
        return mapResp;
    }
}

