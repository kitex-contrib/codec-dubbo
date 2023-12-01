package util

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

// RunAndTestDubboJavaClient is responsible for starting dubbo-java echo consumer.
//
//	-params:
//	    dir: maven dubbo-java package directory
//	    class: the java class you want to start
//	    args: additional cmd-line arguments, if there is not, please set nil
//	    methods: the methods that you want to test, e.g. EchoBool
func RunAndTestDubboJavaClient(t *testing.T, dir, class string, args, methods []string) {
	mvnClean(t, dir)
	pipe, cancel := mvnExec(t, dir, class, args)
	methodSet := make(map[string]bool, len(methods))
	for _, method := range methods {
		methodSet[method] = false
	}
	op := new(outputParser)
	op.init(methodSet)
	op.run(t, pipe, cancel)
}

func mvnClean(t *testing.T, dir string) {
	cleanCmd := exec.Command("mvn", "clean", "package")
	cleanCmd.Dir = dir
	if _, err := cleanCmd.Output(); err != nil {
		t.Fatal(err)
	}
}

func mvnExec(t *testing.T, dir, class string, args []string) (io.ReadCloser, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	cmdArgs := []string{
		"-Djava.net.preferIPv4Stack=true",
		fmt.Sprintf("-Dexec.mainClass=%s", class),
	}
	if len(args) > 0 {
		addArgs := strings.Join(args, " ")
		cmdArgs = append(cmdArgs, fmt.Sprintf("-Dexec.args=\"%s\"", addArgs))
	}
	cmdArgs = append(cmdArgs, "exec:java")
	cmd := exec.CommandContext(ctx, "mvn", cmdArgs...)
	cmd.Dir = dir
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	return pipe, cancel
}

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

func (op *outputParser) run(t *testing.T, rd io.Reader, cancel context.CancelFunc) {
	reader := bufio.NewReader(rd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err := op.parse(line); err != nil {
			t.Error(err)
		}
		if op.checkStop() {
			break
		}
	}
	cancel()
	if missing := op.missingFuncs(); len(missing) > 0 {
		t.Errorf("missing funcs: %s", missing)
	}
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
