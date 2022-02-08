package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		count := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exit...")
				done <- struct{}{}
				return
			default:
				fmt.Println(count)
				count++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	<-done
}
