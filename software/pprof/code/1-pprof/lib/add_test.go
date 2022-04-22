package lib

import "testing"

func TestAdd(t *testing.T)  {
	_ = Add("go-programing-tour-book")
}

func BenchmarkAdd(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Add("go-programing-tour-book")
	}
}

// cpu采样
//go:generate go test -bench=. -cpuprofile=cpu.cpuprofile

// memory采样
//go:generate go test -bench=. -memprofile=mem.profile

//go:generate go tool pprof -http=:6006 cpu.cpuprofile

//go:generate go tool pprof -http=:6007 mem.profile
