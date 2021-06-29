package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring is a func that return a channel to communicate with it.
// <-chan string means receives-only channel of string
func boring(msg string) <-chan string {
	ch := make(chan string)

	// we launch(启动) goroutine inside a func
	// that sends the data a channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s --> %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		// the sender should close the channel
		close(ch)
	}()
	return ch
}

// code source:
// https://talks.golang.org/2012/concurrency.slide#26
func main() {

	joe := boring("Joe")
	ahn := boring("Ahn")

	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ahn)
	}

	// or we can simple use the for range
	//for msg := range joe{
	//	fmt.Println(msg)
	//}
	fmt.Println("You're both boring. I'm leaving")

}
