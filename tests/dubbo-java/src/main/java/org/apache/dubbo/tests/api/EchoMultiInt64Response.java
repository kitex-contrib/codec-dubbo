package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiInt64Response implements Serializable {
    long baseResp;
    ArrayList<Long> listResp;
    HashMap<Long, Long> mapResp;

    public EchoMultiInt64Response(long baseResp, ArrayList<Long> listResp, HashMap<Long, Long> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public long getBaseResp() {
        return baseResp;
    }

    public ArrayList<Long> getListResp() {
        return listResp;
    }

    public HashMap<Long, Long> getMapResp() {
        return mapResp;
    }
}
