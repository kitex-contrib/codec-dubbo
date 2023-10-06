package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiByteResponse implements Serializable {
    byte baseResp;
    ArrayList<Byte> listResp;
    HashMap<Byte, Byte> mapResp;

    public EchoMultiByteResponse(byte baseResp, ArrayList<Byte> listResp, HashMap<Byte, Byte> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public byte getBaseResp() {
        return baseResp;
    }

    public ArrayList<Byte> getListResp() {
        return listResp;
    }

    public HashMap<Byte, Byte> getMapResp() {
        return mapResp;
    }
}
