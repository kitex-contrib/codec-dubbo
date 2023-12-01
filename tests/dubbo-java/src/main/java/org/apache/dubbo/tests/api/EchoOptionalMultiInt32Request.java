package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.Map;
import java.util.List;

public class EchoOptionalMultiInt32Request implements Serializable {
    int basicReq;
    Integer packReq;
    List<Integer> listReq;
    Map<Boolean, Integer> mapReq;

    public EchoOptionalMultiInt32Request(int basicReq, Integer packReq, List<Integer> listReq, Map<Boolean, Integer> mapReq) {
        this.basicReq = basicReq;
        this.packReq = packReq;
        this.listReq = listReq;
        this.mapReq = mapReq;
    }

    public int getBasicReq() {
        return basicReq;
    }

    public void setBasicReq(int basicReq) {
        this.basicReq = basicReq;
    }

    public Integer getPackReq() {
        return packReq;
    }

    public void setPackReq(Integer packReq) {
        this.packReq = packReq;
    }

    public List<Integer> getListReq() {
        return listReq;
    }

    public void setListReq(List<Integer> listReq) {
        this.listReq = listReq;
    }

    public Map<Boolean, Integer> getMapReq() {
        return mapReq;
    }

    public void setMapReq(Map<Boolean, Integer> mapReq) {
        this.mapReq = mapReq;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiInt32Request{" +
                "basicReq=" + basicReq +
                ", packReq=" + packReq +
                ", listReq=" + listReq +
                ", mapReq=" + mapReq +
                '}';
    }
}
