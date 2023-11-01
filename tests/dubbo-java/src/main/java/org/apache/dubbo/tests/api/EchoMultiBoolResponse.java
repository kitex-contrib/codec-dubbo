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

    public boolean[] getListRespToArray() {
        boolean[] arr = new boolean[listResp.size()];
        for (int i = 0; i < listResp.size(); i++) {
            arr[i] = listResp.get(i);
        }
        return arr;
    }

    public Map<Boolean, Boolean> getMapResp() {
        return mapResp;
    }
}

