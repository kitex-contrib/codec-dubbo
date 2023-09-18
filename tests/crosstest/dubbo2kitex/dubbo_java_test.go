/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dubbo2kitex

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"testing"
)

type outputParser struct {
	funcMap   map[string]bool
	funcNum   int
	finishNum int
	fail      *regexp.Regexp
	exception *regexp.Regexp
	end       *regexp.Regexp
}

func (op *outputParser) init(set map[string]bool) {
	op.funcMap = set
	op.funcNum = len(set)
	op.fail = regexp.MustCompile("^{(.*?)} fail")
	op.exception = regexp.MustCompile("^{(.*?)} exception: {(.*?)}")
	op.end = regexp.MustCompile("^{(.*?)} end")
}

func (op *outputParser) parse(output string) error {
	if matches := op.end.FindStringSubmatch(output); len(matches) == 2 {
		if _, ok := op.funcMap[matches[1]]; ok {
			op.funcMap[matches[1]] = true
			op.finishNum++
		}
		return nil
	}
	if matches := op.fail.FindStringSubmatch(output); len(matches) == 2 {
		if _, ok := op.funcMap[matches[1]]; ok {
			op.funcMap[matches[1]] = true
			return fmt.Errorf("%s failed", matches[0])
		}
		return nil
	}
	if matches := op.exception.FindStringSubmatch(output); len(matches) == 3 {
		if _, ok := op.funcMap[matches[1]]; ok {
			op.funcMap[matches[1]] = true
			return fmt.Errorf("%s catch exception: %s", matches[1], matches[2])
		}
		return nil
	}

	return nil
}

func (op *outputParser) checkStop() bool {
	if op.finishNum == op.funcNum {
		return true
	}

	return false
}

func (op *outputParser) missingFuncs() []string {
	var res []string
	for funcName, flag := range op.funcMap {
		if !flag {
			res = append(res, funcName)
		}
	}

	return res
}

func TestDubboJava(t *testing.T) {
	testDir := "../../dubbo-java"
	// initialize mvn packages
	cleanCmd := exec.Command("mvn", "clean", "package")
	cleanCmd.Dir = testDir
	if _, err := cleanCmd.Output(); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "mvn",
		"-Djava.net.preferIPv4Stack=true",
		"-Dexec.mainClass=org.apache.dubbo.tests.client.Application",
		"exec:java")
	cmd.Dir = testDir
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	reader := bufio.NewReader(pipe)
	parser := new(outputParser)
	parser.init(map[string]bool{
		// comment lines mean dubbo-java can not support
		"EchoBool":     false,
		"EchoByte":     false,
		"EchoInt16":    false,
		"EchoInt32":    false,
		"EchoInt64":    false,
		"EchoDouble":   false,
		"EchoString":   false,
		"EchoBinary":   false,
		"EchoBoolList": false,
		//"EchoByteList":   false,
		//"EchoInt16List":  false,
		"EchoInt32List":  false,
		"EchoInt64List":  false,
		"EchoDoubleList": false,
		"EchoStringList": false,
		// hessian2 can not support encoding [][]byte
		//"EchoBinaryList":   false,
		"EchoBool2BoolMap": false,
		//"EchoBool2ByteMap":   false,
		//"EchoBool2Int16Map":  false,
		"EchoBool2Int32Map":  false,
		"EchoBool2Int64Map":  false,
		"EchoBool2DoubleMap": false,
		"EchoBool2StringMap": false,
		//"EchoBool2BinaryMap": false,
	})
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err := parser.parse(line); err != nil {
			t.Error(err)
		}
		if parser.checkStop() {
			break
		}
	}
	cancel()
	if missing := parser.missingFuncs(); len(missing) > 0 {
		t.Errorf("missing funcs: %s", missing)
	}
}
