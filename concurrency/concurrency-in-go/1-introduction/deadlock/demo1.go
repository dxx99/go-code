package main

import "sync"

type Val struct {
	mu  sync.Mutex
	Num int
}

func (v *Val) Add() {
	v.mu.Lock()
	defer v.mu.Unlock()

	//fatal error: all goroutines are asleep - deadlock!
	n := v.Read()
	v.Num = n + 1
}

func (v *Val) Read() int {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.Num
}

// 死锁问题：给没有释放的继续加锁
func main() {
	v := new(Val)
	v.Add()
}
