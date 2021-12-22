package main


//go:generate /usr/bin/curl localhost:8001/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["dxx99"],"id":0}'

func main() {

}
