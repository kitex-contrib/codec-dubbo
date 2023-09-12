# About this directory

This code is used for crosstest between dubbo and kitex. It is under developing.

## Start the service consumer

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.apache.dubbo.tests.client.Application exec:java
```

## Start the service provider

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.apache.dubbo.tests.provider.Application exec:java
```
