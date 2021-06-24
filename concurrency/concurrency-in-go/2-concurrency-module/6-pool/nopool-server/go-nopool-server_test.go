package main_test

import (
	"io/ioutil"
	"net"
	"testing"
)

//go:generate go test -bench=. -benchtime=10s
func BenchmarkNetworkNoPoolRequest(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8011")
		if err != nil {
			b.Fatal("cannot dial host: ", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read : %v\n", err)
		}
		conn.Close()
	}
}




