package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 通过并行处理，解决管道阻塞问题
// fan-out 启动多个goroutines来处理来自管道的输入过程
// fan-in 将多个结果组合到一个通道中的过程
//
// 产生这种计算的场景？
// 1. 不依赖模块之前的计算结果
// 2. 运行需要很长的时间
func main() {

	// 产生一个随机数
	randNum := func() interface{} {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(5 * 1e7)
	}

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

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	// 素数查找函数
	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				integer -= 1
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
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

	startTime := time.Now()
	randIntStream := toInt(done, repeatFn(done, randNum))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(startTime))

}
