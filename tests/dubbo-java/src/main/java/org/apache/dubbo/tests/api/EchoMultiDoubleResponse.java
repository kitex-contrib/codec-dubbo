package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiDoubleResponse implements Serializable {
    double baseResp;
    List<Double> listResp;
    Map<Double, Double> mapResp;

    public EchoMultiDoubleResponse(double baseResp, List<Double> listResp, Map<Double, Double> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public double getBaseResp() {
        return baseResp;
    }

    public List<Double> getListResp() {
        return listResp;
    }

    public Map<Double, Double> getMapResp() {
        return mapResp;
    }
}
