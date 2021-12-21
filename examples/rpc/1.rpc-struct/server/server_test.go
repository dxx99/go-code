package main

import (
	"fmt"
	"reflect"
	"rpc-struct/server/lib"
	"sync"
	"testing"
)

func TestCal_Reflect(t *testing.T) {
	cal := new(lib.Cal)

	typ := reflect.TypeOf(cal)
	val := reflect.ValueOf(cal)
	fmt.Println(typ)
	fmt.Println(val)
	fmt.Println(reflect.Indirect(val).Type().Name())

	// get function
	fmt.Println(typ.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		fmt.Println("method type: ", m.Type)
		fmt.Println("method name: ", m.Name)
	}
}

// Server represents an RPC Server.
type Server struct {
	reqLock    sync.Mutex // protects freeReq
	freeReq    *Request
}

func NewServer() *Server {
	return &Server{freeReq: &Request{
		ServiceMethod: "",
		Seq:           0,
		next:          nil,
	}}
}

type Request struct {
	ServiceMethod string   // format: "Service.Method"
	Seq           uint64   // sequence number chosen by client
	next          *Request // for free list in Server
}

func (s *Server) SetRequest(serviceMethod string) {
	s.reqLock.Lock()
	defer s.reqLock.Unlock()

	req := &Request{
		ServiceMethod: serviceMethod,
		Seq:           0,
		next:          nil,
	}
	lastNode := s.freeReq
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = req
}

func (s *Server) GetRequest() *Request {
	s.reqLock.Lock()
	req := s.freeReq
	if req == nil {
		req = new(Request)
	} else {
		s.freeReq = req.next
		*req = Request{}
	}
	s.reqLock.Unlock()
	return req
}



func TestRpc_Request(t *testing.T)  {
	s := NewServer()
	s.SetRequest("math.Add")
	s.SetRequest("math.Sub")
	s.SetRequest("math.Div")

	fmt.Println("getRequest:", s.GetRequest())
	fmt.Println("getRequest:", s.GetRequest())


	node := s.freeReq
	fmt.Println("first Node: ", node)
	for node != nil {
		node = node.next
		fmt.Println(node)
	}
}




