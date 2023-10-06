package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiBoolResponse implements Serializable {
    boolean baseResp;
    ArrayList<Boolean> listResp;
    HashMap<Boolean, Boolean> mapResp;

    public EchoMultiBoolResponse(boolean baseResp, ArrayList<Boolean> listResp, HashMap<Boolean, Boolean> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public boolean getBaseResp() {
        return baseResp;
    }

    public ArrayList<Boolean> getListResp() {
        return listResp;
    }

    public HashMap<Boolean, Boolean> getMapResp() {
        return mapResp;
    }
}

