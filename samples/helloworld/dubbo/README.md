# About this directory

This code is used for kitex-dubbo interoperability samples.

## Start the service consumer

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.cloudwego.kitex.samples.client.Application exec:java
```

## Start the service provider

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.cloudwego.kitex.samples.provider.Application exec:java
```
