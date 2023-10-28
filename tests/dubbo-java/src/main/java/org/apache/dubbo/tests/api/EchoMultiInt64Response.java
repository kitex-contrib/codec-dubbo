package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoMultiInt64Response implements Serializable {
    long baseResp;
    List<Long> listResp;
    Map<Long, Long> mapResp;

    public EchoMultiInt64Response(long baseResp, List<Long> listResp, Map<Long, Long> mapResp) {
        this.baseResp = baseResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public long getBaseResp() {
        return baseResp;
    }

    public List<Long> getListResp() {
        return listResp;
    }

    public long[] getListRespToArray() {
        long[] arr = new long[listResp.size()];
        for (int i = 0; i < listResp.size(); i++) {
            arr[i] = listResp.get(i);
        }
        return arr;
    }

    public Map<Long, Long> getMapResp() {
        return mapResp;
    }
}
