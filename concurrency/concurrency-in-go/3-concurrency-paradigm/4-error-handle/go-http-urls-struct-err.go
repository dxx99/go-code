package main

import (
	"fmt"
	"net/http"
)

// Result 通过结构体包装错误
type Result struct {
	Error error
	Response *http.Response
}


func main() {

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		res := make(chan Result)

		go func() {
			defer close(res)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{
					Error:    err,
					Response: resp,
				}
				select {
				case <-done:
					return
				case res <- result:

				}
			}
		}()
		return res
	}

	done := make(chan interface{})
	defer close(done)	// 这里也可以用作超时控制


	urls := []string{"https://baidu.com", "https://badhost"}

	for res := range checkStatus(done, urls...) {
		fmt.Printf("err=%s, response:%v\n", res.Error, res.Response)
	}

}
