protoc-gen-kitex-dubbo
========

`protoc-gen-kitex-dubbo` is a protoc plugin that can generate go file to implement the dubbo extension or java dubbo API.

Installation
------------

Note: before executing the following commands, **make sure your `GOPATH` environment is properly set**.

Using `go install`:

`GO111MODULE=on go install github.com/kitex-contrib/codec-dubbo/tools/protoc-gen-kitex-dubbo`

Or build from source:

```shell
git clone https://github.com/kitex-contrib/codec-dubbo.git
cd codec-dubbo/tools/protoc-gen-kitex-dubbo
export GO111MODULE=on
go mod tidy && go mod vendor
go build
go install
```

Usage
-----

```
kitex -type protobuf -protobuf-plugin=kitex-dubbo:ext_lang=java:. -service GreetService  ./api.proto
```
