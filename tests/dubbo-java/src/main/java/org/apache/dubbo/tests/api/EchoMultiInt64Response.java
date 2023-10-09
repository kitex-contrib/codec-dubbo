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

    public Map<Long, Long> getMapResp() {
        return mapResp;
    }
}
