# Steps

## Generate code

Generate kitex_gen without thrift encode/decode code (with frugal slim template)

```bash
kitex -module xxx -thrift template=slim -service kitex-server api.thrift
```

## Remove thrift related code

* imports
* thrift & frugal tag
* ServiceInfo: `PayloadCodec:    kitex.Thrift`

Just search for `thrift` in this directory.

## Implement iface.Message

Implement `iface.Message` interface (Encode, Decode, GetTypes) for all `KitexArgs` and `KitexResult` types in `kitex_gen/echo/k-api.go`

## Register POJO

Register all Request & Response structs (defined in `kitex_gen/echo/api.go`) in `kitex_gen/echo/register.go:init()`

Need to implement `hessian.POJO` for these structs (`JavaClassName() string`)

## Add PayloadCodec option

* Client: kitex_gen/echo/testservice/client.go 

```go
options = append(options, client.WithCodec(hessian2.NewHessian2Codec()))
```

* Server: kitex_gen/echo/testservice/server.go

```go
options = append(options, server.WithCodec(hessian2.NewHessian2Codec()))
```
