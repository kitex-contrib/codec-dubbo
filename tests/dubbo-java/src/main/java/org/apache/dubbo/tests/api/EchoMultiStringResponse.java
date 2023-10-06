package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiStringResponse implements Serializable {
    String baseResp;
    ArrayList<String> listResp;
    HashMap<String, String> mapResp;

    public EchoMultiStringResponse(String baseResp, ArrayList<String> listResp, HashMap<String, String> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public String getBaseResp() {
        return baseResp;
    }

    public ArrayList<String> getListResp() {
        return listResp;
    }

    public HashMap<String, String> getMapResp() {
        return mapResp;
    }
}
