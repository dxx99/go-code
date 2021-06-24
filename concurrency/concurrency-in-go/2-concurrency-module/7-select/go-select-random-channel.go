package main

import "fmt"

// select-case的选择是随机的
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	close(ch1)
	close(ch2)

	ch1Num, ch2Num := 0, 0

	for i := 0; i < 1000; i++ {
		select {
		case <-ch1:
			ch1Num++
		case <-ch2:
			ch2Num++
		}
	}

	fmt.Printf("ch1执行的次数%d, ch2执行的次数%d\n", ch1Num, ch2Num)
}
