package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiInt32Response implements Serializable {
    int baseResp;
    ArrayList<Integer> listResp;
    HashMap<Integer, Integer> mapResp;

    public EchoMultiInt32Response(int baseResp, ArrayList<Integer> listResp, HashMap<Integer, Integer> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public int getBaseResp() {
        return baseResp;
    }

    public ArrayList<Integer> getListResp() {
        return listResp;
    }

    public HashMap<Integer, Integer> getMapResp() {
        return mapResp;
    }
}
