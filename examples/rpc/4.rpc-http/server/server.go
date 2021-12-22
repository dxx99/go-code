package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {

}

func (h *HelloService) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}

func main() {
	err := rpc.Register(new(HelloService))
	if err != nil {
		log.Fatal("register error:", err)
	}

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer: writer,
			ReadCloser: request.Body,
		}

		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8001", nil)
}
