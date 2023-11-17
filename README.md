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

|     thrift 类型      |    golang 类型     | hessian2 类型 |       默认 java 类型       |                可拓展 java 类型                 |
|:------------------:|:----------------:|:-----------:|:----------------------:|:------------------------------------------:|
|        bool        |       bool       |   boolean   |   java.lang.Boolean    |                  boolean                   |
|        byte        |       int8       |     int     |     java.lang.Byte     |                    byte                    |
|        i16         |      int16       |     int     |    java.lang.Short     |                   short                    |
|        i32         |      int32       |     int     |   java.lang.Integer    |                    int                     |
|        i64         |      int64       |    long     |     java.lang.Long     |                    long                    |
|       double       |     float64      |   double    |    java.lang.Double    |    double <br> float / java.lang.Float     |
|       string       |      string      |   string    |    java.lang.String    |                     -                      |
|       binary       |      []byte      |   binary    |         byte[]         |                     -                      |
|    list\<bool>     |      []bool      |    list     |     List\<Boolean>     |      boolean[] / ArrayList\<Boolean>       |
|     list\<i32>     |     []int32      |    list     |     List\<Integer>     |        int[] / ArrayList\<Integer>         |
|     list\<i64>     |     []int64      |    list     |      List\<Long>       |         long[] / ArrayList\<Long>          |
|   list\<double>    |    []float64     |    list     |     List\<Double>      | double[] / ArrayList\<Double> <br> float[] |
|   list\<string>    |     []string     |    list     |     List\<String>      |       String[] / ArrayList\<String>        |
|  map\<bool, bool>  |  map[bool]bool   |     map     | Map\<Boolean, Boolean> |         HashMap\<Boolean, Boolean>         |
|  map\<bool, i32>   |  map[bool]int32  |     map     | Map\<Boolean, Integer> |         HashMap\<Boolean, Integer>         |
|  map\<bool, i64>   |  map[bool]int64  |     map     |  Map\<Boolean, Long>   |          HashMap\<Boolean, Long>           |
| map\<bool, double> | map[bool]float64 |     map     | Map\<Boolean, Double>  |         HashMap\<Boolean, Double>          |
| map\<bool, string> | map[bool]string  |     map     | Map\<Boolean, String>  |         HashMap\<Boolean, String>          |

**重要提示**：

1. 映射表中的 map 类型并没有被完全列举，当前仅包含经过测试的用例。

2. 不支持在 map 类型中使用包含和 **binary** 类型的键值。

3. 由于 **float32** 在 thrift 中不是有效的类型，DubboCodec 将 **float**(java) 映射到了 **float64**(go)，可以在 idl 中使用方法注解指定 **double** 映射为 **float**，具体可参考 [api.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/tests/kitex/api.thrift)。

4. dubbo-java 不支持对包含 **byte**、**short**、**float** 键值的 Map 类型解码，建议避开 dubbo-java 不兼容的用法，可以在定义接口的响应字段时使用 **struct** 来包裹 map。


### 类型拓展

#### 自定义映射

在 **thrift** 的方法后面使用 `hessian.argsType` 注解标签可以指定每个参数映射到 **java** 的类型。

**注解格式**
```thrift
(hessian.argsType="req1JavaType,req2JavaType,req3JavaType,...")
```
其中，每个 reqJavaType 可以使用 `-` 或不填写，表示该参数将使用默认的类型映射。

在初始化 **DubboCodec** 时使用 `WithFileDescriptor` Option，传入生成的 `FileDescriptor`，即可指定 **kitex -> dubbo-java** 的类型映射。

**示例**
```thrift
namespace go echo

service EchoService {
   i64 Echo(1: i32 req1, 2: list<i32> req2, 3: map<i32, i32> req3) (hessian.argsType="int,int[],java.util.HashMap")
   // 前两个字段使用默认的类型映射（可留空或填写 "-"）
   i64 EchoDefaultType(1: i32 req1, 2: i64 req2, 3: bool req3, 4: string req4) (hessian.argsType=",-,bool,string")
}
```

#### 其它类型

由于 **thrift** 类型的局限性，**kitex** 与 **dubbo-java** 映射时有很多不兼容的类型。 
DubboCodec 在 [codec-dubbo/java](https://github.com/kitex-contrib/codec-dubbo/tree/main/java) 包中提供了更多 **thrift** 不支持的 **java** 类型。

为了启用这些类型，你可以在 **Thrift IDL** 中使用 `include "java.thrift"` 导入它们，并且在使用 **kitex** 脚手架工具生成代码时添加 `-hessian2 java_extension` 参数来拉取该拓展包。

kitex 脚手架工具会自动下载 [java.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/java/java.thrift)，你也可以手动下载后放到对应位置。

目前支持的类型包含 `java.lang.Object`、`java.util.Date` 等，更多类型可以参考 [java.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/java/java.thrift)。

**示例**
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

### 方法重载

在 **thrift** 的方法后面使用 `JavaMethodName` 注解标签可以指定该方法在 java 侧的名称。
通过此方式可以实现调用 java 的重载方法。

你可以将多个不同的方法指向 java 的同一个方法，DubboCodec 会根据不同的参数类型调用到 java 侧对应的方法。

**示例**
```thrift
namespace go echo

service EchoService {
    string EchoMethodA(1: bool req) (JavaMethodName="EchoMethod")
    string EchoMethodB(1: i32 req) (JavaMethodName="EchoMethod")
    string EchoMethodC(1: i32 req) (JavaMethodName="EchoMethod", hessian.argsType="int")
    string EchoMethodD(1: bool req1, 2: i32 req2) (JavaMethodName="EchoMethod")
 }
```

### 服务注册与服务发现

目前仅支持基于 zookeeper 的**接口级**服务发现与服务注册，**应用级**服务发现以及服务注册将在后续迭代中支持。

## 开始

[**完整代码**](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples//helloworld/).

### 安装命令行工具

```shell
# 安装 kitex 命令行工具 (注：待发布 v0.8.0 后可改为 `@latest` )
go install github.com/cloudwego/kitex/tool/cmd/kitex@4b3520fbdb5a7d347df1de79d6252efed08ebdf2

# 安装 thriftgo 命令行工具 (注：待发布 v0.3.3 后可改为 `@latest` )
go install github.com/cloudwego/thriftgo@d3508eeb6136bc20ba2f79a04ac878a1595c1cc5
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
kitex -module kitex-dubbo-demo -protocol Hessian2 -service GreetService ./api.thrift
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

## 服务注册与发现

用于该功能的配置分为以下两个层次：
1. [registry/options.go](https://github.com/kitex-contrib/codec-dubbo/tree/main/registries/zookeeper/registry/options.go) 与 [resolver/options.go](https://github.com/kitex-contrib/codec-dubbo/tree/main/registries/zookeeper/resolver/options.go) 中的WithXXX函数提供注册中心级别的配置，请使用这些函数生成```registry.Registry```
和```discovery.Resolver```实例。
2. 服务级别的配置由```client.WithTag```与```server.WithRegistryInfo```进行传递，/registries/common.go提供Tag Keys:

|            Tag Key             |           client侧作用            |                server侧作用                 |
|:------------------------------:|:------------------------------:|:----------------------------------------:|
|    **DubboServiceGroupKey**    |           调用的服务所属的组            | dubbo支持在一个Interface下对多个服务划分组，指定注册的服务所属的组 |
|   **DubboServiceVersionKey**   |            调用的服务版本             | dubbo支持在一个Interface下对多个服务划分版本，指定注册的服务版本  |
|  **DubboServiceInterfaceKey**  | 调用的服务在dubbo体系下对应的InterfaceName |      注册的服务在dubbo体系下对应的InterfaceName      |
|   **DubboServiceWeightKey**    |                                |                注册的服务具有的权重                |
| **DubboServiceApplicationKey** |                                |               注册的服务所属的应用名                |

目前仅支持 zookeeper 作为注册中心。

### 接口级服务发现

#### 客户端初始化

```go
import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries"
	// 该resolver专门用于与dubbo体系下的zookeeper进行交互
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/resolver"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
)

func main() {
	intfName := "org.cloudwego.kitex.samples.api.GreetProvider"
	res, err := resolver.NewZookeeperResolver(
		// 指定 zookeeper 服务器的地址，可指定多个，请至少指定一个 
		resolver.WithServers("127.0.0.1:2181"),
	)
	if err != nil {
		panic(err)
	}
	cli, err := greetservice.NewClient("helloworld",
		// 配置 ZookeeperResolver
		client.WithResolver(res),
		// 配置 DubboCodec
		client.WithCodec(
			dubbo.NewDubboCodec(
				// 指定想要调用的 dubbo Interface，该 Interface 请与下方的 DubboServiceInterfaceKey 值保持一致
				dubbo.WithJavaClassName(intfName),
			),
		),
		// 指定想要调用的 dubbo Interface
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

**重要提示**
1. 用于 DubboCodec 的```WithJavaClassName```应与用于```regitries.DubboServiceInterfaceKey```的值保持一致。

### 接口级服务注册

#### 服务端初始化

```go
import (
	"github.com/cloudwego/kitex/server"
	kitex_registry "github.com/cloudwego/kitex/pkg/registry"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries"
	// 该resolver专门用于与dubbo体系下的zookeeper进行交互 
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/registry"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"log"
	"net"
)

func main() {
	intfName := "org.cloudwego.kitex.samples.api.GreetProvider"
	reg, err := registry.NewZookeeperRegistry(
	    // 指定 zookeeper 服务器的地址，可指定多个，请至少指定一个 
	    registry.WithServers("127.0.0.1:2181"),
	)
	if err != nil {
	    panic(err)
	}
	// 指定服务端将要监听的地址
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		// 配置 DubboCodec
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName(intfName),
		)),
		server.WithRegistry(reg),
		// 配置dubbo URL元数据
		server.WithRegistryInfo(&kitex_registry.Info{
		    Tags: map[string]string{
			    registries.DubboServiceInterfaceKey: intfName,
				// application请与dubbo所设置的ApplicationConfig保持一致，此处仅为示例
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

**重要提示**
1. 用于 DubboCodec 的```WithJavaClassName```应与用于```regitries.DubboServiceInterfaceKey```的值保持一致。

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
