package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;

public class EchoMultiDoubleResponse implements Serializable {
    double baseResp;
    ArrayList<Double> listResp;
    HashMap<Double, Double> mapResp;

    public EchoMultiDoubleResponse(double baseResp, ArrayList<Double> listResp, HashMap<Double, Double> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public double getBaseResp() {
        return baseResp;
    }

    public ArrayList<Double> getListResp() {
        return listResp;
    }

    public HashMap<Double, Double> getMapResp() {
        return mapResp;
    }
}
