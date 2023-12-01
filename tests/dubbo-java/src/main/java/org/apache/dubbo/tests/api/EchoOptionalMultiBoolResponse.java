package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoOptionalMultiBoolResponse implements Serializable {
    boolean basicResp;
    Boolean packResp;
    List<Boolean> listResp;
    Map<Boolean, Boolean> mapResp;

    public EchoOptionalMultiBoolResponse(boolean basicResp, Boolean packResp, List<Boolean> listResp, Map<Boolean, Boolean> mapResp) {
        this.basicResp = basicResp;
        this.packResp = packResp;
        this.listResp = listResp;
        this.mapResp = mapResp;
    }

    public boolean getBasicResp() {
        return basicResp;
    }

    public void setBasicResp(boolean basicResp) {
        this.basicResp = basicResp;
    }

    public Boolean getPackResp() {
        return packResp;
    }

    public void setPackResp(Boolean packResp) {
        this.packResp = packResp;
    }

    public List<Boolean> getListResp() {
        return listResp;
    }

    public void setListResp(List<Boolean> listResp) {
        this.listResp = listResp;
    }

    public Map<Boolean, Boolean> getMapResp() {
        return mapResp;
    }

    public void setMapResp(Map<Boolean, Boolean> mapResp) {
        this.mapResp = mapResp;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiBoolResponse{" +
                "basicResp=" + basicResp +
                ", packResp=" + packResp +
                ", listResp=" + listResp +
                ", mapResp=" + mapResp +
                '}';
    }
}
