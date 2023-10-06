package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiInt16Response implements Serializable {
    short baseResp;
    ArrayList<Short> listResp;
    HashMap<Short, Short> mapResp;

    public EchoMultiInt16Response(short baseResp, ArrayList<Short> listResp, HashMap<Short, Short> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public short getBaseResp() {
        return baseResp;
    }

    public ArrayList<Short> getListResp() {
        return listResp;
    }

    public HashMap<Short, Short> getMapResp() {
        return mapResp;
    }
}
