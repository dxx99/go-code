package main

import (
	"context"
	"fmt"
	"github.com/dxx99/go-code/examples/base/context/core"
	"time"
)

// context 数据传输
func main() {
	ctx := context.WithValue(context.Background(), "parent", "parent")
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	go func() {
		time.Sleep(5*time.Second)
		cancel()
	}()

	go core.NewWorker().Run(ctx)
	fmt.Println("main goroutine waiting...")


	time.Sleep(10*time.Second)

}


