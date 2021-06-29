package main

import "fmt"

func main() {
	done := make(chan string)
	stringSteam := make(chan string, 10)
	for _, item := range []string{"a", "b", "c"} {
		select {
		case <-done:
		case stringSteam <- item:
		}
	}
	close(stringSteam)
	//fmt
	for ch := range stringSteam {
		fmt.Println(ch)
	}

	// 无限循环等待停止
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}

			//执行非抢占式任务
		}
	}()

	//2. 变种
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// 执行非抢占任务
			}
		}
	}()

}
