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

    public double[] getListRespToArray() {
        double[] arr = new double[listResp.size()];
        for (int i = 0; i < listResp.size(); i++) {
            arr[i] = listResp.get(i);
        }
        return arr;
    }

    public Map<Double, Double> getMapResp() {
        return mapResp;
    }
}
