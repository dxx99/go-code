package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 数据重复函数
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	// 处理的数据传给某channel
	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int, ) <-chan interface{} {

		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("开始执行改进版函数")

	// 执行repeatFn
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {

		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	doneFn := make(chan interface{})
	defer close(doneFn)

	rand := func() interface{} {
		rand.Seed(time.Now().UnixNano())
		return rand.Int()
	}

	for num := range take(doneFn, repeatFn(doneFn, rand), 10) {
		fmt.Println(num)
	}
}
