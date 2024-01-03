# codec-dubbo

中文 | [English](README_ENG.md)

[Kitex](https://github.com/cloudwego/kitex) 为了支持 **kitex \<-\> dubbo 互通** 推出的 dubbo 协议编解码器。


## 简介

1. 支持 Kitex Client 请求 Dubbo-Java、Dubbo-Go Server，也支持 Dubbo-Java、Dubbo-Go Client 请求 Kitex Server。
2. 基于 IDL（兼容 Thrift 语法）生成项目脚手架，包括 kiten_gen (Client/Server Stub，编解码代码等）、main.go（server 初始化）和 handler.go（method handler）。
3. IDL 注解扩展：可指定类型对应的 Java Class，扩展支持 Thrift 非标类型（如 `float32`、`interface{}`(java.lang.Object)、`time.Time`(java.util.Date)）。
4. 支持 zookeeper 服务注册和发现（接口级别）

## 开始

- 完整代码参见: [samples/helloworld](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld/)
- 由于 dubbo-java、dubbo-go-hessian2 编解码器的原因，使用时存在一些限制，详见后文

### 安装命令行工具

```shell
# 安装 kitex 命令行工具 (version >= v0.8.0)
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# 安装 thriftgo 命令行工具 (version >= v0.3.3)
go install github.com/cloudwego/thriftgo@latest
```

### Server 端

#### 生成脚手架

创建项目目录（以 `demo-server` 为例），并初始化 go module:
```bash
mkdir ~/demo-server && cd ~/demo-server
go mod init demo-server
```

在目录下按需编写 IDL (兼容 Thrift 语法），例如 `api.thrift`：
```thrift
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

Kitex 命令行工具生成项目脚手架（注意需指定 `-protocol Hessian2`）：
> 需 `-service` 参数，生成 `kitex_gen` 目录、server 初始化代码 `main.go` 和 method handler `handler.go`
```bash
kitex -module demo-server -protocol Hessian2 -service GreetService ./api.thrift
go mod tidy
```

**注意**:

1. 如需与 Dubbo Java 互通，`api.thrift` 中定义的每个结构体都应添加注解 `JavaClassName`，值为对应的 Java 类名称。

#### Server 初始化

修改 `main.go`，在 `NewServer` 中指定 DubboCodec：

```go
import (
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "demo-server/helloworld/kitex/kitex_gen/hello/greetservice"
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
			// 配置 Kitex 服务所对应的 Java Interface. 其他 dubbo 客户端和 kitex 客户端可以通过这个名字进行调用。
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```

**注意**:

1. 每个 Kitex Server 对应一个 DubboCodec 实例，请不要在多个 Server 间共享同一个实例。
2. 如需注册到服务发现中心，请参考后文相关章节。


#### Server Handler: 业务逻辑

在 **handler.go** 中添加业务逻辑，例如：

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
```

#### 启动 server：

编译：生成到 output 目录
```bash
sh build.sh
```

启动 Server：
```bash
sh output/bootstrap.sh
```

### Client 端

#### 生成项目脚手架

（新项目）准备项目目录：
```bash
mkdir demo-client && cd demo-client
go mod init demo-client
```

准备 IDL（参考上文 server 端相关内容）。

生成脚手架：
> 无需 `-service`，只生成 `kitex_gen` 目录
```bash
kitex -module demo-client -protocol Hessian2 ./api.thrift
go mod tidy
```

#### 使用 Kitex Dubbo Client

参考代码：

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
		// 指定想要访问的服务端地址
		client.WithHostPorts("127.0.0.1:21000"),
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

**注意**:

1. 每个 Kitex Client 对应一个 DubboCodec 实例，请不要在多个客户端之间共享同一个实例。

## 功能特性

### 类型映射

|     thrift 类型     |    golang 类型   | hessian2 类型 |       默认 java 类型   |                可拓展 java 类型              |
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

3. 由于 **float32** 在 thrift 支持的类型，DubboCodec 将 **float**(java) 映射到了 **float64**(go)，可在 idl 中使用方法注解指定 **double** 映射为 **float**，具体可参考 [api.thrift](https://github.com/kitex-contrib/codec-dubbo/blob/main/tests/kitex/api.thrift)。

4. dubbo-java 不支持对包含 **byte**、**short**、**float** 键值的 Map 类型解码，建议避开 dubbo-java 不兼容的用法，可以在定义接口的响应字段时使用 **struct** 来包裹 map。

**空值(null)兼容性**：

1. 由于 go 中部分基础类型不支持空值（如：**bool**、**int64**等），不建议 java 端向 go 端不可为空的类型传递 `null` 值。

2. java-server 向 kitex-client 不可为空的类型传递 `null` 值时，会被转换为对应类型的零值。

3. 不支持 java-client 向 kitex-server 不可为空的类型传递 `null` 值，计划于后续版本 [#69](https://github.com/kitex-contrib/codec-dubbo/issues/69) 添加拓展类型以支持 java 空值。

4. 如果对 `null` 值有需求，建议将不可为空的类型包装在 **struct** 中，在 go 端将接收到对应类型的零值，DubboCodec 对 **struct** 中字段的空值有较好的支持。

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

#### 其它类型（java.lang.Object, java.util.Date）

由于 **thrift** 类型的局限性，**kitex** 与 **dubbo-java** 映射时有一些不兼容的类型。 
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

### 异常处理

**codec-dubbo** 将异常定义为实现了以下接口的错误，你可以像处理错误一样处理 java 中的异常：
```go
type Throwabler interface {
	Error() string
	JavaClassName() string
	GetStackTrace() []StackTraceElement
}
```

#### 常见异常

**codec-dubbo** 在[pkg/hessian2/exception](https://github.com/kitex-contrib/codec-dubbo/tree/main/pkg/hessian2/exception)目录下提供了java中常见的异常，目前支持 java.lang.Exception ，更多异常将在后续迭代中加入。
常见异常无需命令行工具的支持，直接引用即可。

##### client端提取异常

```go
import (
	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
)

func main() {
	resp, err := cli.Greet(context.Background(), true)
	if err != nil {
		// FromError 返回 Throwabler
        exceptionRaw, ok := hessian2_exception.FromError(err)
        if !ok {
        // 视作常规错误处理	
        } else {
            // 若不关心 exceptionRaw 的具体类型，直接调用 Throwabler 提供的方法即可
			klog.Errorf("get %s type Exception", exceptionRaw.JavaClassName())
			
			// 若想获得 exceptionRaw 的具体类型，需要进行类型转换，但前提是已知该具体类型
			exception := exceptionRaw.(*hessian2_exception.Exception)
        }
    }

}
```

##### server端返回异常

```go
import (
	hessian2_exception "github.com/kitex-contrib/codec-dubbo/pkg/hessian2/exception"
)

func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
    return "", hessian2_exception.NewException("Your detailed message")
}
```

#### 自定义异常

java 中的自定义业务异常往往会继承一个基础异常，这里以 CustomizedException 为例，CustomizedException 继承了 java.lang.Exception：
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

为了在 kitex 侧定义与之对应的异常，在 **thrift** 中编写如下定义：

```thrift
exception CustomizedException {
    1: required java.Exception exception (thrift.nested="true")
    2: required string customizedMessage
}(JavaClassName="org.cloudwego.kitex.samples.api.CustomizedException")
```

和[其它类型](#其它类型javalangobject-javautildate)一样，需要在使用 **kitex** 脚手架工具生成代码时添加 `-hessian2 java_extension` 参数来拉取拓展包。 

使用方法与[常见异常](#常见异常)一致。

## 服务注册与发现

> 目前仅支持基于 zookeeper 的**接口级**服务发现与服务注册，**应用级**服务发现以及服务注册计划在后续迭代中支持。

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
