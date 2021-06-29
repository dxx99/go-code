package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var doWork = func(str <-chan string) <-chan interface{} {
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

// 访问 http://127.0.0.1:6060/debug/pprof
// go内存泄露的简单例子
func main() {
	http.HandleFunc("/go-memory-leak", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan string)
		for i := 0; i < 10000; i++ {
			go doWork(ch)
			ch <- fmt.Sprintf("输入Id:%d\n", i)
		}
		//1. 不执行close会出现goroutine泄露
		//close(ch)
		_, _ = w.Write([]byte("执行成功"))
	})
	log.Fatal(http.ListenAndServe(":6060", nil))
}
