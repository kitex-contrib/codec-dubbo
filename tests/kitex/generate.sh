#!/bin/bash

cd $(dirname "$0")

kitex -module github.com/kitex-contrib/codec-dubbo/tests/kitex -protocol hessian2 -hessian2 java_extension -service TestService api.thrift