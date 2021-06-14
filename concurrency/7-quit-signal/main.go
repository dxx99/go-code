package main

import "fmt"

// for-select-quit control goroutine exit
func boring(msg string, quit chan string) <-chan string {
	ch := make(chan string)

	//we launch goroutine inside a func
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s --> %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("clean up")
				quit <- "see you"
				return
			}
		}
	}()
	// return a channel to caller
	return ch
}

//code source:
// https://talks.golang.org/2012/concurrency.slide#38
func main() {
	quit := make(chan string)
	ch := boring("Joe", quit)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	quit <- "Bye"
	fmt.Printf("Joe say: %s\n", <-quit)
}
