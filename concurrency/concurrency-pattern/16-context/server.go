package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8011", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	log.Println("handler starting")
	defer log.Println("handler stopping")

	// todo 传参有问题
	vTime, ok := (ctx.Value("time_out")).(time.Duration)
	if ok {
		log.Println("time out is ", vTime)
	} else {
		vTime = time.Duration(1000)
	}
	ctx, cancel := context.WithCancel(r.Context())
	time.AfterFunc(vTime * time.Millisecond, cancel)

	select {
	case <-time.After(5 * time.Second):
		_, _ = fmt.Fprintf(w, "hello")
	case <-ctx.Done():
		log.Printf("请求失败：%s\n", ctx.Err())
		_, _ = fmt.Fprintf(w, ctx.Err().Error())
	}
}
