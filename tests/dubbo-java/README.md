# About this directory

This code is used for crosstest between dubbo and kitex. It is under developing.

## Start the service provider

```bash
mvn clean package
mvn -Djava.net.preferIPv4Stack=true -Dexec.mainClass=org.apache.dubbo.samples.provider.Application exec:java
```
