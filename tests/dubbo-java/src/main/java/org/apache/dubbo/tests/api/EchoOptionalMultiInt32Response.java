package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoOptionalMultiInt32Response implements Serializable {
    int basicResp;
    Integer packResp;
    List<Integer> listResp;
    Map<Boolean, Integer> mapResp;

    public EchoOptionalMultiInt32Response(int basicResp, Integer packResp, List<Integer> listResp, Map<Boolean, Integer> mapResp) {
        this.basicResp = basicResp;
        this.packResp = packResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public int getBasicResp() {
        return basicResp;
    }

    public void setBasicResp(int basicResp) {
        this.basicResp = basicResp;
    }

    public Integer getPackResp() {
        return packResp;
    }

    public void setPackResp(Integer packResp) {
        this.packResp = packResp;
    }

    public List<Integer> getListResp() {
        return listResp;
    }

    public void setListResp(List<Integer> listResp) {
        this.listResp = listResp;
    }

    public Map<Boolean, Integer> getMapResp() {
        return mapResp;
    }

    public void setMapResp(Map<Boolean, Integer> mapResp) {
        this.mapResp = mapResp;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiInt32Response{" +
                "basicResp=" + basicResp +
                ", packResp=" + packResp +
                ", listResp=" + listResp +
                ", mapResp=" + mapResp +
                '}';
    }
}
