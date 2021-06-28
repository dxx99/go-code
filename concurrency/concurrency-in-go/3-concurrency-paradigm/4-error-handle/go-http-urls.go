package main

import (
	"fmt"
	"net/http"
)

func main() {

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan	*http.Response {
		resps := make(chan *http.Response)

		go func() {
			defer close(resps)
			for _, url := range urls {
				resp, err := http.Get(url)
				if	err != nil {
					fmt.Println(err)	// 这里就是bad主机的报错
					continue
				}
				select {
				case <-done:
					return
				case resps <- resp:

				}
			}
		}()
		return resps
	}

	done := make(chan interface{})
	defer close(done)	// 这里也可以用作超时控制


	urls := []string{"https://baidu.com", "https://badhost"}

	for response := range checkStatus(done, urls...) {
		fmt.Printf("response:%v\n", response.Status)
	}

}
