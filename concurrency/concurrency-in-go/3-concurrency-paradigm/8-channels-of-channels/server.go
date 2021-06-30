package main

import "net/http"

// todo 应该构成rpc请求
// https://golang.org/doc/effective_go#concurrency  channels of channels
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8991", nil)
}
