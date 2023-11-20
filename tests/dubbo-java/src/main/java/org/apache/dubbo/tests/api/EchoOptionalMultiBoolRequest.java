package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class EchoOptionalMultiBoolRequest implements Serializable {
    boolean basicReq;
    Boolean packReq;
    List<Boolean> listReq;
    Map<Boolean, Boolean> mapReq;

    public EchoOptionalMultiBoolRequest(boolean basicReq, Boolean packReq, List<Boolean> listReq, Map<Boolean, Boolean> mapReq) {
        this.basicReq = basicReq;
        this.packReq = packReq;
        this.listReq = listReq;
        this.mapReq = mapReq;
    }

    public boolean getBasicReq() {
        return basicReq;
    }

    public void setBasicReq(boolean basicReq) {
        this.basicReq = basicReq;
    }

    public Boolean getPackReq() {
        return packReq;
    }

    public void setPackReq(Boolean packReq) {
        this.packReq = packReq;
    }

    public List<Boolean> getListReq() {
        return listReq;
    }

    public void setListReq(List<Boolean> listReq) {
        this.listReq = listReq;
    }

    public Map<Boolean, Boolean> getMapReq() {
        return mapReq;
    }

    public void setMapReq(Map<Boolean, Boolean> mapReq) {
        this.mapReq = mapReq;
    }

    @Override
    public String toString() {
        return "EchoOptionalMultiBoolRequest{" +
                "basicReq=" + basicReq +
                ", packReq=" + packReq +
                ", listReq=" + listReq +
                ", mapReq=" + mapReq +
                '}';
    }
}
