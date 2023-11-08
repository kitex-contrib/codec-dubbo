package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiFloatResponse implements Serializable {
    float baseResp;
    List<Float> listResp;
    Map<Float, Float> mapResp;

    public EchoMultiFloatResponse(float baseResp, List<Float> listResp, Map<Float, Float> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public float getBaseResp() {
        return baseResp;
    }

    public List<Float> getListResp() {
        return listResp;
    }

    public float[] getListRespToArray() {
        float[] arr = new float[listResp.size()];
        for (int i = 0; i < listResp.size(); i++) {
            arr[i] = listResp.get(i);
        }
        return arr;
    }

    public Map<Float, Float> getMapResp() {
        return mapResp;
    }
}
