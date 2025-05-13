//go:build mage
// +build mage

package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alimy/mir/v5/core"
	"github.com/alimy/mir/v5/engine"
	"github.com/livekit/mageutil"
)

var Default = Generate

// regenerate mir and protobuf
func Generate() (err error) {
	return anyFun(cleanAuto, Mir, Rpc, Enum, Fmt)
}

// regenerate mir RESTful API interface
func Mir() error {
	fmt.Println("Generating mir")
	if err := engine.Generate(
		core.UseGin(),
		core.Schema("mirc"),
		core.SinkPath("auto"),
		core.WatchCtxDone(true),
		core.RunMode(core.InSerialMode),
	); err != nil {
		return err
	}
	return gofmt("./auto/api/...")
}

// regenerate RPC API interface
func Rpc() error {
	fmt.Println("Generating gRPC source code")
	cmd := exec.Command("buf", "generate", "proto")
	mageutil.ConnectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	return gofmt("./auto/rpc/...")
}

func Enum() error {
	fmt.Println("Generating Enum source code")
	cmd := exec.Command("go", "generate", "./internal/model/enum/...")
	mageutil.ConnectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	return gofmt("./internal/model/enum/...")
}

// run tests
func Test() error {
	fmt.Println("Testing start")
	cmd := exec.Command("go", "test", "-race", "./...")
	connectStd(cmd)
	return cmd.Run()
}

// format source code
func Fmt() error {
	return gofmt("./...")
}

func gofmt(pattern string) error {
	fmt.Printf("Formating source code [%s]\n", pattern)
	cmd := exec.Command("go", "fmt", pattern)
	mageutil.ConnectStd(cmd)
	return cmd.Run()
}

func getToolPath(name string) (string, error) {
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	// check under gopath
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	p := filepath.Join(gopath, "bin", name)
	if _, err := os.Stat(p); err != nil {
		return "", err
	}
	return p, nil
}

func cleanAuto() error {
	fmt.Println("Delete auto directory")
	cmd := exec.Command("rm", "-rf", "auto")
	connectStd(cmd)
	return cmd.Run()
}

func connectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func anyFun(fn ...func() error) (err error) {
	for _, f := range fn {
		if err = f(); err != nil {
			break
		}
	}
	return
}
