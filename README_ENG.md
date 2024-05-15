# codec-dubbo

English | [中文](README.md)

[Kitex](https://github.com/cloudwego/kitex)'s dubbo codec for **kitex \<-\> dubbo interoperability**.

## Introduction

1. Support Kitex Client to request Dubbo-Java/Dubbo-Go Server, and also support Dubbo-Java/Dubbo-Go Client to request Kitex Server.
2. Generate project scaffolding based on IDL (compatible with Thrift syntax), including kiten_gen (Client/Server Stub, codec code, etc.), main.go (server initialization), and handler.go (method handler).
3. IDL annotation extension: You can specify the corresponding Java Class for the an argument, the response or a field within a struct, and support non-standard Thrift types (such as `float32`, `interface{}` (java.lang.Object), `time.Time` (java.util.Date)).
4. Support zookeeper service registration and discovery (interface level).

## Getting Started

**Full example in**: [samples/helloworld](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld/).

### Prerequisites

```shell
# install the latest kitex cmd tool (version >= v0.8.0)
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# install thriftgo (version >= v0.3.5)
go install github.com/cloudwego/thriftgo@latest
```

Note: Customized Exception is not officially released, but you can install this version of Kitex for a try:
> go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.9.0-rc1


### Kitex Server

#### Generating kitex stub codes

Create a project directory (let's say demo-server) and initialize the go module:
```bash
mkdir ~/demo-server && cd ~/demo-server
go mod init demo-server
```
Write the IDL (compatible with Thrift syntax) under the directory as needed, for example `api.thrift`:
```Thrift
namespace go hello

struct GreetRequest {
    1: required string req,
}(JavaClassName="org.cloudwego.kitex.samples.api.GreetRequest")

struct GreetResponse {
    1: required string resp,
}(JavaClassName="org.cloudwego.kitex.samples.api.GreetResponse")

service GreetService {
    string Greet(1: string req)
    GreetResponse GreetWithStruct(1: GreetRequest req)
}
```

Generate the project scaffold using the Kitex command-line tool (note that you need to specify `-protocol Hessian2`):
```bash
kitex -module demo-server -protocol Hessian2 -service GreetService ./api.thrift
go mod tidy
```

**Notes:**

1. For interoperatability with Dubbo Java, each structure defined in your IDL should add an annotation `JavaClassName` with the value corresponding to the Java class name.

#### Server Initialization

Modify `main.go` and specify DubboCodec in `NewServer`:

```go
import (
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "demo-server/helloworld/kitex/kitex_gen/hello/greetservice"
	"log"
	"net"
)

func main() {
	// Specify the address to listen
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		// Config DubboCodec
		server.WithCodec(dubbo.NewDubboCodec(
			// Config the corresponding Java Interface; other client should call with this name.
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

**Note:**

1. Each Kitex Server corresponds to an instance of a DubboCodec, please do not share the same instance among multiple Servers.
2. For service registry, please refer to the relevant sections later.

#### Server Handler: business logic

Add your business loginc to `handler.go`, such as:

```go
import (
    "context"
    hello "demo-server/kitex_gen/hello"
)

func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
	return "Hello " + req, nil
}

func (s *GreetServiceImpl) GreetWithStruct(ctx context.Context, req *hello.GreetRequest) (resp *hello.GreetResponse, err error) {
	return &hello.GreetResponse{Resp: "Hello " + req.Req}, nil
}
````

#### Run your server

Compile: generate the `output` directory with the binary and necessary scripts
```bash
sh build.sh
```

Start the server：
```bash
sh output/bootstrap.sh
```

### Kitex Client

#### Generating kitex stub codes

(For a new project) Prepare the project directory:
```bash
mkdir demo-client && cd demo-client
go mod init demo-client
```

Prepare the IDL (please refer to the same section in `Kitex Server` above).

Generate the scaffold:
> No need the paramteter `-service`, only generating the kitex_gen directory
```bash
kitex -module demo-client -protocol Hessian2 ./api.thrift
go mod tidy
```

#### Use Kitex Dubbo Client

Code for reference
```go
package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"  
	"demo-client/kitex_gen/hello"
	"demo-client/kitex_gen/hello/greetservice"
)

func main() {
	cli, err := greetservice.NewClient("helloworld",
		// Specify the targeting service
		client.WithHostPorts("127.0.0.1:21000"),
		// Config DubboCodec
		client.WithCodec(
			dubbo.NewDubboCodec(
				// Specify the targeting interface, should be same as the server side
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
			),
		),
	)
	if err != nil {
		panic(err)
	}

	resp, err := cli.Greet(context.Background(), "world")
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("resp: %s", resp)
	
	respWithStruct, err := cli.GreetWithStruct(context.Background(), &hello.GreetRequest{Req: "world"})
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("respWithStruct: %s", respWithStruct.Resp)
}
```

**Note:**

1. Each Kitex Client corresponds to an instance of a DubboCodec, please do not share the same instance among multiple Clients.

## Feature List

### Type Mapping

|    thrift type     |   golang type    | hessian2 type |   default java type    |            extendable java type            |
|:------------------:|:----------------:|:-------------:|:----------------------:|:------------------------------------------:|
|        bool        |       bool       |    boolean    |   java.lang.Boolean    |                  boolean                   |
|        byte        |       int8       |      int      |     java.lang.Byte     |                    byte                    |
|        i16         |      int16       |      int      |    java.lang.Short     |                   short                    |
|        i32         |      int32       |      int      |   java.lang.Integer    |                    int                     |
|        i64         |      int64       |     long      |     java.lang.Long     |                    long                    |
|       double       |     float64      |    double     |    java.lang.Double    |    double <br> float / java.lang.Float     |
|       string       |      string      |    string     |    java.lang.String    |                     -                      |
|       binary       |      []byte      |    binary     |         byte[]         |                     -                      |
|    list\<bool>     |      []bool      |     list      |     List\<Boolean>     |      boolean[] / ArrayList\<Boolean>       |
|     list\<i32>     |     []int32      |     list      |     List\<Integer>     |        int[] / ArrayList\<Integer>         |
|     list\<i64>     |     []int64      |     list      |      List\<Long>       |         long[] / ArrayList\<Long>          |
|   list\<double>    |    []float64     |     list      |     List\<Double>      | double[] / ArrayList\<Double> <br> float[] |
|   list\<string>    |     []string     |     list      |     List\<String>      |       String[] / ArrayList\<String>        |
|  map\<bool, bool>  |  map[bool]bool   |      map      | Map\<Boolean, Boolean> |         HashMap\<Boolean, Boolean>         |
|  map\<bool, i32>   |  map[bool]int32  |      map      | Map\<Boolean, Integer> |         HashMap\<Boolean, Integer>         |
|  map\<bool, i64>   |  map[bool]int64  |      map      |  Map\<Boolean, Long>   |          HashMap\<Boolean, Long>           |
| map\<bool, double> | map[bool]float64 |      map      | Map\<Boolean, Double>  |         HashMap\<Boolean, Double>          |
| map\<bool, string> | map[bool]string  |      map      | Map\<Boolean, String>  |         HashMap\<Boolean, String>          |

**Important notes**:

1. The list of map types is not exhaustive and includes only tested cases.

2. Using keys of **binary** type in map types is not supported.

3. Since **float32** is not a valid type in Thrift, DubboCodec maps **float**(java) to **float64**(go). You can specify the mapping of **double** to **float** in the IDL with method annotations. For an example, please refer to [api.thrift](https://github.com/kitex-contrib/codec-dubbo-tests/blob/v0.1.0/code/kitex/api.thrift#L173).

4. dubbo-java does not support decoding map types that contain **byte**, **short**, or **float** key values. It is recommended to avoid practices incompatible with dubbo-java. You can use **struct** to wrap the map when defining response fields for interfaces.

**Null Compatibility**:

1. Due to some basic types in Go not supporting null values (e.g., **bool**, **int64**, etc.), it is not recommended for the Java side to pass `null` values to non-nullable types in Go.

2. When Java-server passes a `null` value to a non-nullable type in the Kitex-client, it will be converted to the zero value of the corresponding type.

3. Java-client does not support passing `null` values to non-nullable types in the Kitex-server. There are plans to add extended types in future versions [#69](https://github.com/kitex-contrib/codec-dubbo/issues/69) to support Java null values.

4. If there is a requirement for `null` values, it is recommended to wrap non-nullable types in a **struct**. On the Go side, the corresponding type's zero value will be received, and the DubboCodec provides good support for null values in **struct** fields.

### Type Extension

#### Custom Mapping

After a method in **thrift**, you can use the `hessian.argsType` annotation tag to specify the mapping of each parameter to Java types.

**Annotation Format**
```thrift
(hessian.argsType="req1JavaType,req2JavaType,req3JavaType,...")
```
Here, each `reqJavaType` can either be left blank or use a `-`, indicating that the default type mapping will be used for that parameter.

When initializing the DubboCodec, use the WithFileDescriptor option and pass in the generated FileDescriptor to specify the type mapping from kitex -> dubbo-java.

**Example**

```thrift
namespace go echo

service EchoService {
   i64 Echo(1: i32 req1, 2: list<i32> req2, 3: map<i32, i32> req3) (hessian.argsType="int,int[],java.util.HashMap")
   // Use the default type mapping for the first 2 arguments
   i64 EchoDefaultType(1: i32 req1, 2: i64 req2, 3: bool req3, 4: string req4) (hessian.argsType=",-,bool,string")
}
```

#### Other Types (java.lang.Object, java.util.Date)

Due to the limitations of the **thrift** type system, there are some incompatible types when mapping **kitex** to **dubbo-java**. The DubboCodec, located in the [codec-dubbo/java](https://github.com/kitex-contrib/codec-dubbo/tree/main/java) package, provides support for additional **java** types that are not supported by **thrift**.

To enable these types, you should add into **Thrift IDL** `include "java.thrift"`, and generate code with the **kitex** scaffolding tool with the `-hessian2 java_extension` parameter.

You can download [java.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/java/java.thrift) manually to the targeting path (especially when you need a special version), otherwise **kitex** will do it for you.

The currently supported types include `java.lang.Object`, `java.util.Date`. For more details, you can refer to [java.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/java/java.thrift).

**Example**
```thrift
namespace go echo
include "java.thrift"

service EchoService {
    // java.lang.Object
    i64 EchoString2ObjectMap(1: map<string, java.Object> req)
    // java.util.Date
    i64 EchoDate(1: java.Date req)
}
```

### Method Overloading

After a method in **thrift**, you can use the `JavaMethodName` annotation tag to specify the name of the method on the Java side.
This allows you to invoke overloaded methods in Java.

By doing so, you can point multiple different methods to the same Java method. DubboCodec will call the corresponding Java method on the basis of different parameter types.

**Example**
```thrift
namespace go echo

service EchoService {
    string EchoMethodA(1: bool req) (JavaMethodName="EchoMethod")
    string EchoMethodB(1: i32 req) (JavaMethodName="EchoMethod")
    string EchoMethodC(1: i32 req) (JavaMethodName="EchoMethod", hessian.argsType="int")
    string EchoMethodD(1: bool req1, 2: i32 req2) (JavaMethodName="EchoMethod")
 }
```

### Protocol Probing

Simultaneous support Dubbo and Thrift protocols, example：
```
import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/remote/trans/detection"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "demo-server/kitex_gen/hello/greetservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		server.WithTransHandlerFactory(detection.NewSvrTransHandlerFactory(netpoll.NewSvrTransHandlerFactory(),
			dubbo.NewSvrTransHandlerFactory(
                // set the Java interface name corresponding to the Kitex service
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider")),
			nphttp2.NewSvrTransHandlerFactory(),
		)),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

```

### Exception Handling

**codec-dubbo** defines exceptions as **error** that implement the following interface. You can handle exceptions in Java as you could handle **error** in Go:

```go
type Throwabler interface {
	Error() string
	JavaClassName() string
	GetStackTrace() []StackTraceElement
}
```

#### Common Exceptions

**codec-dubbo** provides commonly used Java exceptions in the [pkg/hessian2/exception](https://github.com/kitex-contrib/codec-dubbo/tree/main/pkg/hessian2/exception) directory. Currently, it supports java.lang.Exception, and more exceptions will be added in subsequent iterations.
Common exceptions do not require command line tool support and could be directly referenced.

##### Extracting Exception on the Client Side

```go
import (
	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
)

func main() {
	resp, err := cli.Greet(context.Background(), true)
	if err != nil {
        // FromError returns a Throwabler
        exceptionRaw, ok := hessian2_exception.FromError(err)
        if !ok {
        // Treat as a regular error handling	
        } else {
            // If you are not concerned with the specific type of exceptionRaw, just call the methods provided by Throwabler
            klog.Errorf("get %s type Exception", exceptionRaw.JavaClassName())
			
            // If you want to obtain the specific type of exceptionRaw, you need to perform a type conversion, but this requires knowing the specific type
            exception := exceptionRaw.(*hessian2_exception.Exception)
        }
    }
}
```

##### Returning Exception on the Server Side

```go
import (
	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
)

func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
    return "", hessian2_exception.NewException("Your detailed message")
}
```

#### Customized Exceptions

Customized business exceptions in Java often inherit a base exception. Here, we use CustomizedException as an example, which inherits from java.lang.Exception:

```java
public class CustomizedException extends Exception {
    private final String customizedMessage;
    public CustomizedException(String customizedMessage) {
        super();
        this.customizedMessage = customizedMessage;
    }

    public String getCustomizedMessage() {
        return this.customizedMessage;
    }
}
```

To define a corresponding exception on the kitex side, write the following definition in **thrift**:

```thrift
exception CustomizedException {
    1: required java.Exception exception (thrift.nested="true")
    2: required string customizedMessage
}(JavaClassName="org.cloudwego.kitex.samples.api.CustomizedException")
```

Like [other types](#other-types--javalangobject-javautildate-), you need to add the `-hessian2 java_extension` parameter when generating code with the **kitex** scaffolding tool to pull the extension package.

The usage is consistent with [Common Exceptions](#common-exceptions).

## Service Registry and Service Discovery

> Currently, only **Interface-Level** service discovery based on zookeeper is supported. **Application-Level** support is planned in a future release.


The configurations used for this functionality are divided into the following two levels:
1. The WithXXX functions in [registry/options.go](https://github.com/kitex-contrib/codec-dubbo/tree/main/registries/zookeeper/registry/options.go) and [resolver/options.go](https://github.com/kitex-contrib/codec-dubbo/tree/main/registries/zookeeper/resolver/options.go) provide registry-level configurations; use these functions to generate ```registry.Registry```
   and ```discovery.Resolver``` instances.
2. Service level configurations are passed by ```client.WithTag``` with ```server.WithRegistryInfo```, and /registries/common.go provides Tag Keys:

|            Tag Key             |                              client side effect                              |                                                              server side effect                                                               |
|:------------------------------:|:----------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------------------------:|
|    **DubboServiceGroupKey**    |                The group to which the called service belongs                 | dubbo supports the division of multiple services into groups under an Interface, specifying the group to which the registered service belongs |
|   **DubboServiceVersionKey**   |                      The version of the service called                       |                 dubbo supports versioning of multiple services under one Interface, specifying the registered service version                 |
|  **DubboServiceInterfaceKey**  | The corresponding InterfaceName of the called service under the dubbo system |                               The corresponding InterfaceName of the registered service under the dubbo system                                |
|   **DubboServiceWeightKey**    |                                                                              |                                                         Weight of registered service                                                          |
| **DubboServiceApplicationKey** |                                                                              |                                      The name of the application to which the registered service belongs                                      |

Currently only zookeeper is supported as a registry.

### Interface-Level service discovery

#### initializing client

```go
import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/codec-dubbo/registries"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	// this resolver is dedicated to interacting with the zookeeper in the dubbo system
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/resolver"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
)

func main() {
	intfName := "org.cloudwego.kitex.samples.api.GreetProvider"
	res, err := resolver.NewZookeeperResolver(
		// specify the addresses of the zookeeper servers, please specify at least one
		resolver.WithServers("127.0.0.1:2181"),
	)
	if err != nil {
		panic(err)
	}
	cli, err := greetservice.NewClient("helloworld",
		// configure ZookeeperResolver
		client.WithResolver(res),
		// configure DubboCodec
		client.WithCodec(
			dubbo.NewDubboCodec(
				// target dubbo Interface，this Interface should be consistent with the value of DubboServiceInterfaceKey below
				dubbo.WithJavaClassName(intfName),
			),
		),
		// target dubbo Interface Name
        client.WithTag(registries.DubboServiceInterfaceKey, intfName),
	)
	if err != nil {
		panic(err)
	}

	resp, err := cli.Greet(context.Background(), "world")
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("resp: %s", resp)
	
	respWithStruct, err := cli.GreetWithStruct(context.Background(), &hello.GreetRequest{Req: "world"})
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("respWithStruct: %s", respWithStruct.Resp)
}
```

Important notes:
1. The ```WithJavaClassName``` for DubboCodec should be consistent with the value of ```registries.DubboServiceInterfaceKey```.

#### initializing server

```go
import (
	"github.com/cloudwego/kitex/server"
	kitex_registry "github.com/cloudwego/kitex/pkg/registry"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries"
	// this registry is dedicated to interacting with the zookeeper in the dubbo system 
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/registry"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"log"
	"net"
)

func main() {
	intfName := "org.cloudwego.kitex.samples.api.GreetProvider"
	reg, err := registry.NewZookeeperRegistry(
        // specify the addresses of the zookeeper servers, please specify at least one 
	    registry.WithServers("127.0.0.1:2181"),
	)
	if err != nil {
	    panic(err)
	}
	// specify the address that the server will listen to
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		// configure DubboCodec
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName(intfName),
		)),
		server.WithRegistry(reg),
		// configure dubbo URL metadata
		server.WithRegistryInfo(&kitex_registry.Info{
		    Tags: map[string]string{
			    registries.DubboServiceInterfaceKey: intfName,
				// application value should be consistent with ApplicationConfig set in dubbo, this is only an example
				registries.DubboServiceApplicationKey: "application-name",
            }
        }),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

Important notes:
1. The ```WithJavaClassName``` for DubboCodec should be consistent with the value of ```registries.DubboServiceInterfaceKey```.

## Benchmark

### Benchmark Environment

CPU: **Intel(R) Xeon(R) Gold 5118 CPU @ 2.30GHz**  
Memory: **192GB**

### Benchmark Code

Referring to [dubbo-go-benchmark](https://github.com/dubbogo/dubbo-go-benchmark). Converting dubbo client and dubbo server
to kitex client and kitex server. Please see [this](https://github.com/kitex-contrib/codec-dubbo-tests/tree/v0.1.0/benchmark).

### Benchmark Result

#### kitex -> kitex

```shell
bash deploy.sh kitex_server -p 21001
bash deploy.sh kitex_client -p 21002 -addr "127.0.0.1:21001"
bash deploy.sh stress -addr '127.0.0.1:21002' -t 1000000 -p 100 -l 256
```

Result:

| average rt |  tps  | success rate |
|:----------:|:-----:|:------------:|
|  2310628   | 46015 |   1.000000   |
|  2363729   | 44202 |   1.000000   |
|  2256177   | 43280 |   1.000000   |
|  2194147   | 43171 |   1.000000   |

Resource:

|   process_name    | %CPU  | %MEM |
|:-----------------:|:-----:|:----:|
| kitex_client_main | 914.6 | 0.0  |
| kitex_server_main | 520.5 | 0.0  |
|    stress_main    | 1029  | 0.1  |

### Benchmark Summary

Since the [**dubbo-go-hessian2**](https://github.com/apache/dubbo-go-hessian2) relies on reflection for encoding/decoding,
there's great potential to improve the performance with a codec based on generated Go code.

A [**fastCodec for Hessian2**](https://github.com/kitex-contrib/codec-dubbo/issues/28) is planned for better performance.

## Acknowledgements

This is a community driven project maintained by [@DMwangnima](https://github.com/DMwangnima).

We extend our sincere appreciation to the dubbo-go development team for their valuable contribution!
- [**dubbo-go**](https://github.com/apache/dubbo-go)
- [**dubbo-go-hessian2**](https://github.com/apache/dubbo-go-hessian2)

## Reference

- [hessian serialization](http://hessian.caucho.com/doc/hessian-serialization.html)
