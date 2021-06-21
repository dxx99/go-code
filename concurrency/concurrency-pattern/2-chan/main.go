package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, ch chan string)  {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s --> %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
// code source:
// https://talks.golang.org/2012/concurrency.slide#20
func main() {
	ch := make(chan string)

	go boring("boring", ch)

	for i := 0; i < 5; i++ {
		// <-ch read the value from boring func
		// <-ch waits for a value to be send
		//
		// Printf format param meaning
		// %q 打印单引号
		// %v 以默认的方式打印变量的值
		// %T 打印变量的类型
		// %+d 带符号打印整型  %5d 表示该整型的最大长度是5
		// %b 打印整型的二精致
		// %U 打印unicode字符
		// %x 小写的十六进制    %X 大写的十六进制
		// %o 不带零的八进制    %#o 带零的八进制
		// %s %d
		fmt.Printf("You say : %q\n", <-ch)
	}

	fmt.Println("You're boring. I'm leaving")
}
