package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

// 生成prof文件：
//go:generate go run go-gen-profile.go --cpuprofile=cpu.prof

//执行命令
//go:generate go build go-gen-profile.go
//go:generate go tool pprof go-gen-profile cpu.prof
func main() {

	// 通过传参生成文件
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	doWork := func(str <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range str {
				fmt.Println("执行任务: ", s)
			}
		}()
		return completed
	}

	for i := 0; i < 100; i++ {
		doWork(nil)
	}

	time.Sleep(3 * time.Second)

}
