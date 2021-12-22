package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	argv0, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatal("lookPath err: ", err)
	}

	var env []string
	env = append(env, os.Environ()...)
	originWd, _ := os.Getwd()
	pid, err := os.StartProcess(argv0, os.Args, &os.ProcAttr{
		Dir:   originWd,
		Env:   env,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatal("get pid err: ", err)
	}

	fmt.Println(0x08)
	fmt.Println("process:", pid)
	fmt.Println(os.Geteuid())
}
