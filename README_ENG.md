# codec-dubbo

English | [中文](README.md)

[Kitex](https://github.com/cloudwego/kitex)'s dubbo codec for **kitex \<-\> dubbo interoperability**.


## Feature List

### Kitex-Dubbo Interoperability

1. **kitex -> dubbo**

Write **api.thrift** based on existing **dubbo Interface API** and [**Type Mapping Table**](#type-mapping). Then use
the latest kitex command tool and thriftgo to generate Kitex's scaffold (including stub code).

In addition to the default type mapping, you can also specify the Java type for request parameter mapping in **thrift** using [**Method Annotation**](#method-annotation).

2. **dubbo -> kitex**

Write dubbo client code based on existing **api.thrift** and [**Type Mapping Table**](#type-mapping).

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

3. Since **float32** is not a valid type in Thrift, DubboCodec maps **float**(java) to **float64**(go). You can specify the mapping of **double** to **float** in the idl using method annotations, Please see [api.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/tests/kitex/api.thrift).

4. dubbo-java does not support decoding map types that contain **byte**, **short**, or **float** key values. It is recommended to avoid practices incompatible with dubbo-java. You can use **struct** to wrap the map when defining response fields for interfaces.

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

#### Other Types

Due to the limitations of the **thrift** type system, there are many incompatible types when mapping **kitex** to **dubbo-java**. The DubboCodec, located in the [codec-dubbo/java](https://github.com/kitex-contrib/codec-dubbo/tree/main/java) package, provides support for additional **java** types that are not supported by **thrift**.

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

### Service Registry and Service Discovery

Currently, only **Interface-Level** service discovery based on zookeeper is supported.
**Application-Level** service discovery and service registration will be supported in subsequent iterations.

## Getting Started

[**Example**](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples//helloworld/).

### Prerequisites

```shell
# install the latest kitex cmd tool (switch to `@latest` after v0.8.0 is released)
go install github.com/cloudwego/kitex/tool/cmd/kitex@4b3520fbdb5a7d347df1de79d6252efed08ebdf2

# install thriftgo (switch to `@latest` after v0.3.3 is released)
go install github.com/cloudwego/thriftgo@d3508eeb6136bc20ba2f79a04ac878a1595c1cc5
```

### Generating kitex stub codes

```shell
mkdir ~/kitex-dubbo-demo && cd ~/kitex-dubbo-demo
go mod init kitex-dubbo-demo

# Replace with your own Thrift IDL
cat > api.thrift << EOF
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

EOF

# Generate Kitex scaffold with the `-protocol Hessian2` option
kitex -module kitex-dubbo-demo -protocol Hessian2 -service GreetService ./api.thrift

```

Important Notes:

1. Each struct in the `api.thrift` should have an annotation named `JavaClassName`, with a value consistent with the target class name in Dubbo Java.

### Finishing business logic and configuration

#### business logic

```go
import (
    "context"
    hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
)

// implement interface in handler.go
func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
	return "Hello " + req, nil
}

func (s *GreetServiceImpl) GreetWithStruct(ctx context.Context, req *hello.GreetRequest) (resp *hello.GreetResponse, err error) {
	return &hello.GreetResponse{Resp: "Hello " + req.Req}, nil
}
```

Implement the interface in **handler.go**.

#### initializing client

```go
import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"  
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
)

func main() {
	cli, err := greetservice.NewClient("helloworld",
		// specify address of target server
		client.WithHostPorts("127.0.0.1:21001"),
		// configure dubbo codec
		client.WithCodec(
			dubbo.NewDubboCodec(
				// target dubbo Interface Name
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

Important notes:
1. Each dubbo Interface corresponds to a `DubboCodec` instance. Please do not share the instance between multiple clients.

#### initializing server

```go
import (
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"log"
	"net"
)

func main() {
	// server address to listen on
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		// configure DubboCodec
		server.WithCodec(dubbo.NewDubboCodec(
			// Interface Name of kitex service. Other dubbo clients and kitex clients can refer to this name for invocation.
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

Important notes:
1. Each Interface Name corresponds to a `DubboCodec` instance. Please do not share the instance between multiple servers.

## Service Registry and Service Discovery

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
to kitex client and kitex server. Please see [this](https://github.com/kitex-contrib/codec-dubbo/tree/main/tests/benchmark).

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
