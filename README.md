# codec-dubbo

中文 | [English](README_ENG.md)

[Kitex](https://github.com/cloudwego/kitex) 为了支持 **kitex \<-\> dubbo 互通** 推出的 dubbo 协议编解码器。

## 功能

### Kitex-Dubbo 互通

1. **kitex -> dubbo**

基于已有的 **dubbo Interface API** 和 [**类型映射**](#类型映射)，编写 **api.thrift**。然后使用最新的 kitex 命令行工具和 thriftgo 生成 kitex 的脚手架代码（包括用于编解码的stub代码）。

除了默认的类型映射外，还可以在 **thrift** 中使用 [**方法注解**](#方法注解) 指定请求参数映射的 java 类型。

2. **dubbo -> kitex**

基于已有的 **api.thrift** 和 [**类型映射**](#类型映射)，编写 dubbo 客户端代码。

### 类型映射

|     thrift 类型      |    golang 类型     | hessian2 类型 |       默认 java 类型       |           可拓展 java 类型           |
|:------------------:|:----------------:|:-----------:|:----------------------:|:-------------------------------:|
|        bool        |       bool       |   boolean   |   java.lang.Boolean    |             boolean             |
|        byte        |       int8       |     int     |     java.lang.Byte     |              byte               |
|        i16         |      int16       |     int     |    java.lang.Short     |              short              |
|        i32         |      int32       |     int     |   java.lang.Integer    |               int               |
|        i64         |      int64       |    long     |     java.lang.Long     |              long               |
|       double       |     float64      |   double    |    java.lang.Double    |             double              |
|       string       |      string      |   string    |    java.lang.String    |                -                |
|       binary       |      []byte      |   binary    |         byte[]         |                -                |
|    list\<bool>     |      []bool      |    list     |     List\<Boolean>     | boolean[] / ArrayList\<Boolean> |
|     list\<i32>     |     []int32      |    list     |     List\<Integer>     |   int[] / ArrayList\<Integer>   |
|     list\<i64>     |     []int64      |    list     |      List\<Long>       |    long[] / ArrayList\<Long>    |
|   list\<double>    |    []float64     |    list     |     List\<Double>      |  double[] / ArrayList\<Double>  |
|   list\<string>    |     []string     |    list     |     List\<String>      |  String[] / ArrayList\<String>  |
|  map\<bool, bool>  |  map[bool]bool   |     map     | Map\<Boolean, Boolean> |   HashMap\<Boolean, Boolean>    |
|  map\<bool, i32>   |  map[bool]int32  |     map     | Map\<Boolean, Integer> |   HashMap\<Boolean, Integer>    |
|  map\<bool, i64>   |  map[bool]int64  |     map     |  Map\<Boolean, Long>   |     HashMap\<Boolean, Long>     |
| map\<bool, double> | map[bool]float64 |     map     | Map\<Boolean, Double>  |    HashMap\<Boolean, Double>    |
| map\<bool, string> | map[bool]string  |     map     | Map\<Boolean, String>  |    HashMap\<Boolean, String>    |

**重要提示**：
1. 映射表中的 map 类型并没有被完全列举，当前仅包含经过测试的用例。
   请勿在map类型中使用包含 **i8**、**i16** 和 **binary** 的键值。
2. 目前不支持float32，因为它在 thrift 中不是有效的类型。计划在后续迭代中支持该类型。

### 方法注解

DubboCodec 支持在 **thrift** 中使用 **方法注解** 指定请求参数需要映射的 java 类型。

**方法注解格式：**
```thrift
(hessian.argsType="req1JavaType,req2JavaType,req3JavaType,...")
```
其中，每个 reqJavaType 可以使用 `-` 或不填写，表示该参数将使用默认的类型映射。

添加方法注解后，使用 kitex 命令行工具生成代码时添加选项 `-thrift with_reflection`，会在生成的脚手架代码中包含 thrift 的 **FileDescriptor**。
在初始化 client 时使用 DubboCodec 提供的 `WithFileDescriptor` Option，传入生成的 **FileDescriptor**，即可指定 **kitex -> dubbo** 的类型映射。

**举例：**
```thrift
service EchoService {
   EchoResponse Echo(1: i32 req1, 2: list<i32> req2, 3: map<i32, i32> req3) (hessian.argsType="int,int[],java.util.HashMap")
   // 使用默认的类型映射
   EchoDefaultTypeResponse EchoDefaultType(1: i32 req1, 2: i64 req2, 3: bool req3, 4: string req4) (hessian.argsType=",-,,-")
}
```

## 开始

[**完整代码**](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples//helloworld/).

### 安装命令行工具

```shell
# 安装 kitex 命令行工具，需指定版本 (kitex >= 0.7.3)
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# 安装 thriftgo 命令行工具
go install github.com/cloudwego/thriftgo@latest
```

### 生成 kitex stub 代码

```shell
mkdir ~/kitex-dubbo-demo && cd ~/kitex-dubbo-demo
go mod init kitex-dubbo-demo

# 编写你所需的 Thrift IDL，此处仅为演示
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

# 使用 `-protocol Hessian2` 配置项生成 Kitex 脚手架代码
kitex -module kitex-dubbo-demo -thrift template=slim -service GreetService -protocol Hessian2 ./api.thrift
```

**重要提示**:
1. api.thrift 中定义的每个结构体都应该有一个名为 JavaClassName 的注解，并且注解值与 Dubbo Java 中对应的类名必须一致。

### 实现业务逻辑并完成初始化

#### 业务逻辑

```go
import (
    "context"
    hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
)

func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
	return "Hello " + req, nil
}

func (s *GreetServiceImpl) GreetWithStruct(ctx context.Context, req *hello.GreetRequest) (resp *hello.GreetResponse, err error) {
	return &hello.GreetResponse{Resp: "Hello " + req.Req}, nil
}
```

实现在 **handler.go** 中定义的接口.

#### 客户端初始化

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
		// 指定想要访问的服务端地址
		client.WithHostPorts("127.0.0.1:21001"),
		// 配置 DubboCodec
		client.WithCodec(
			dubbo.NewDubboCodec(
				// 指定想要调用的 Dubbo Interface
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

**重要提示**:
1. 每个 Dubbo Interface 对应一个 DubboCodec 实例，请不要在多个客户端之间共享同一个实例。

#### 服务端初始化

```go
import (
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"log"
	"net"
)

func main() {
	// 指定服务端将要监听的地址
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		// 配置 DubboCodec
		server.WithCodec(dubbo.NewDubboCodec(
			// 配置 Kitex 服务所对应的 Interface. 其他 dubbo 客户端和 kitex 客户端可以通过这个名字进行调用。
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

**重要提示**:
1. 每个 Dubbo Interface 对应一个 DubboCodec 实例，请不要在多个服务端之间共享同一个实例。

## 性能测试

### 测试环境 

CPU: **Intel(R) Xeon(R) Gold 5118 CPU @ 2.30GHz**  
内存: **192GB**

### 测试代码

测试代码主要参考 [dubbo-go-benchmark](https://github.com/dubbogo/dubbo-go-benchmark). 将 dubbo 客户端和 dubbo 服务端替换成对应的 kitex 客户端和 kitex 服务端。
具体实现请看[代码](https://github.com/kitex-contrib/codec-dubbo/tree/main/tests/benchmark)。

### 测试结果

#### kitex -> kitex

```shell
bash deploy.sh kitex_server -p 21001
bash deploy.sh kitex_client -p 21002 -addr "127.0.0.1:21001"
bash deploy.sh stress -addr '127.0.0.1:21002' -t 1000000 -p 100 -l 256
```

结果:

| 平均响应时间(ns) |  tps  |   成功率    |
|:----------:|:-----:|:--------:|
|  2310628   | 46015 | 1.000000 |
|  2363729   | 44202 | 1.000000 |
|  2256177   | 43280 | 1.000000 |
|  2194147   | 43171 | 1.000000 |

资源占用:

|        进程名        | %CPU  | %内存 |
|:-----------------:|:-----:|:---:|
| kitex_client_main | 914.6 | 0.0 |
| kitex_server_main | 520.5 | 0.0 |
|    stress_main    | 1029  | 0.1 |

### 测试总结

[**dubbo-go-hessian2**](https://github.com/apache/dubbo-go-hessian2) 依赖反射进行编解码，因此可以通过基于生成代码的编解码器来提升性能。

我们将在后续迭代中推出[**Hessian2 fastCodec**](https://github.com/kitex-contrib/codec-dubbo/issues/28)。

## 致谢

这是一个由社区驱动的项目，由 @DMwangnima 维护。

我们衷心感谢dubbo-go开发团队的宝贵贡献！

- [**dubbo-go**](https://github.com/apache/dubbo-go)
- [**dubbo-go-hessian2**](https://github.com/apache/dubbo-go-hessian2)

## 参考

- [hessian序列化协议](http://hessian.caucho.com/doc/hessian-serialization.html)
