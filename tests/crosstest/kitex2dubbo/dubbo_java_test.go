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

package kitex2dubbo

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
)

func runDubboJavaServer() (context.CancelFunc, chan struct{}) {
	finishChan := make(chan struct{})
	testDir := "../../dubbo-java"
	// initialize mvn packages
	cleanCmd := exec.Command("mvn", "clean", "package")
	cleanCmd.Dir = testDir
	if _, err := cleanCmd.Output(); err != nil {
		panic(fmt.Sprintf("mvn clean package failed: %s", err))
	}
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "mvn",
		"-Djava.net.preferIPv4Stack=true",
		"-Dexec.mainClass=org.apache.dubbo.tests.provider.Application",
		"exec:java")
	cmd.Dir = testDir

	go func() {
		var exitErr *exec.ExitError
		if err := cmd.Run(); err == nil || !errors.As(err, &exitErr) {
			panic("dubbo-java server should be terminated by this test process")
		} else {
			finishChan <- struct{}{}
		}
	}()

	return cancel, finishChan
}
