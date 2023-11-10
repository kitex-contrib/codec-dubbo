package org.apache.dubbo.tests.provider;

import org.apache.dubbo.config.ProtocolConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.config.ServiceConfig;
import org.apache.dubbo.config.bootstrap.DubboBootstrap;
import org.apache.dubbo.tests.api.UserProvider;

import java.util.ArrayList;
import java.util.List;

public class RegistryApplication {
    public static void main(String[] args) {
        ServiceConfig<UserProvider> service = new ServiceConfig<>();
        service.setInterface(UserProvider.class);
        service.setRef(new UserProviderImpl());

        ServiceConfig<UserProvider> service1 = new ServiceConfig<>();
        service1.setInterface(UserProvider.class);
        service1.setRef(new UserProviderImplV1());
        service1.setGroup("g1");
        service1.setVersion("v1");

        List<ServiceConfig> list = new ArrayList<>();
        list.add(service);
        list.add(service1);

        String zookeeperAddress = "zookeeper://127.0.0.1:2181";
        RegistryConfig zookeeper = new RegistryConfig(zookeeperAddress);
        zookeeper.setGroup("myGroup");
        zookeeper.setRegisterMode("interface");

        DubboBootstrap.getInstance()
                .application("first-dubbo-provider")
                .registry(zookeeper)
                .protocol(new ProtocolConfig("dubbo", 20001))
                .services(list)
                .start()
                .await();
    }
}
