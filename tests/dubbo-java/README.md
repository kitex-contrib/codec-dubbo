# About this directory

This code is used for crosstest between dubbo and kitex.

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

## Start the service provider

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.apache.dubbo.tests.provider.RegistryApplication exec:java
```