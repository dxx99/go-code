package main

import (
	"fmt"
	"time"
)

func main() {
	for _, s := range []string{"AAA", "BBB", "CCC"} {
		func() {
			fmt.Printf("func say: %s\n", s)
		}()

		// 这是因为循环在任何的goroutine开始运行之前退出，所以s转移到堆中，并保存对字符串切片"CCC"中最后一个值的引用
		go func() {
			fmt.Printf("goroutine say: %s\n", s)
		}()
	}
	time.Sleep(1 * time.Millisecond)
}
