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

|    thrift type     |   golang type    | hessian2 type |   default java type    |      extendable java type       |
|:------------------:|:----------------:|:-------------:|:----------------------:|:-------------------------------:|
|        bool        |       bool       |    boolean    |   java.lang.Boolean    |             boolean             |
|        byte        |       int8       |      int      |     java.lang.Byte     |              byte               |
|        i16         |      int16       |      int      |    java.lang.Short     |              short              |
|        i32         |      int32       |      int      |   java.lang.Integer    |               int               |
|        i64         |      int64       |     long      |     java.lang.Long     |              long               |
|       double       |     float64      |    double     |    java.lang.Double    |             double              |
|       string       |      string      |    string     |    java.lang.String    |                -                |
|       binary       |      []byte      |    binary     |         byte[]         |                -                |
|    list\<bool>     |      []bool      |     list      |     List\<Boolean>     | boolean[] / ArrayList\<Boolean> |
|     list\<i32>     |     []int32      |     list      |     List\<Integer>     |   int[] / ArrayList\<Integer>   |
|     list\<i64>     |     []int64      |     list      |      List\<Long>       |    long[] / ArrayList\<Long>    |
|   list\<double>    |    []float64     |     list      |     List\<Double>      |  double[] / ArrayList\<Double>  |
|   list\<string>    |     []string     |     list      |     List\<String>      |  String[] / ArrayList\<String>  |
|  map\<bool, bool>  |  map[bool]bool   |      map      | Map\<Boolean, Boolean> |   HashMap\<Boolean, Boolean>    |
|  map\<bool, i32>   |  map[bool]int32  |      map      | Map\<Boolean, Integer> |   HashMap\<Boolean, Integer>    |
|  map\<bool, i64>   |  map[bool]int64  |      map      |  Map\<Boolean, Long>   |     HashMap\<Boolean, Long>     |
| map\<bool, double> | map[bool]float64 |      map      | Map\<Boolean, Double>  |    HashMap\<Boolean, Double>    |
| map\<bool, string> | map[bool]string  |      map      | Map\<Boolean, String>  |    HashMap\<Boolean, String>    |

**Important notes**:
1. The list of map types is not exhaustive and includes only tested cases.

2. Using keys of **binary** type in map types is not supported.

3. Since **float32** is not a valid type in Thrift, DubboCodec maps **float**(java) to **float64**(go). You can specify the mapping of **double** to **float** in the idl using method annotations, Please see [api.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/tests/kitex/api.thrift).

4. dubbo-java does not support decoding map types that contain **byte**, **short**, or **float** key values. It is recommended to avoid practices incompatible with dubbo-java. You can use **struct** to wrap the map when defining response fields for interfaces.

### Method Annotation

DubboCodec supports specifying the Java types needed for request parameter mapping in **thrift** using **method annotations**.

**Method Annotation Format:**
```thrift
(hessian.argsType="req1JavaType,req2JavaType,req3JavaType,...")
```
Here, each `reqJavaType` can either be left blank or use a `-`, indicating that the default type mapping will be used for that parameter.

After adding method annotations, use the kitex command line tool to generate code and add the option `-thrift with_reflection`. This will include the **FileDescriptor** of thrift in the generated scaffold code.
When initializing the client, use DubboCodec's `WithFileDescriptor` Option, and pass in the generated **FileDescriptor** to specify the **kitex -> dubbo** type mapping.

**Example:**
```thrift
service EchoService {
   EchoResponse Echo(1: i32 req1, 2: list<i32> req2, 3: map<i32, i32> req3) (hessian.argsType="int,int[],java.util.HashMap")
   // Use the default type mapping
   EchoDefaultTypeResponse EchoDefaultType(1: i32 req1, 2: i64 req2, 3: bool req3, 4: string req4) (hessian.argsType=",-,,-")
}
```

## Getting Started

[**Concrete sample**](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples//helloworld/).

### Prerequisites

```shell
# install the latest kitex cmd tool (kitex >= v0.7.3)
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# install thriftgo
go install github.com/cloudwego/thriftgo@latest
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
# With `-thrift template=slim,with_reflection`, generate code without thrift encoder/decoder && support thrift reflection
kitex -module kitex-dubbo-demo -thrift template=slim,with_reflection -protocol Hessian2 -service GreetService ./api.thrift

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
