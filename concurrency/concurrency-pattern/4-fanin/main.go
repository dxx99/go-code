package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s --> %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		close(ch)
	}()
	return ch
}

// fanIn merge many result
// <-chan string only get the receive value
// fanIn spawns(产生) 2 goroutine to reads the value from 2 channels
// then it send to value to result channel(`ch` channel)
func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for true {
			ch <- <-input1
		}
	}()

	go func() {
		for true {
			ch <- <-input2
		}
	}()

	return ch
}

// fanIn pattern
//┌──────┐
//│input1│────┐
//└──────┘    │   ┌────┐
//            ├───▶ ch │─────▶
//┌──────┐    │   └────┘
//│input2│────┘
//└──────┘
func fanSimple(chs ...<-chan string) <-chan string {
	ch := make(chan string)

	for _, chi := range chs {
		go func(chi <-chan string) {
			ch <- <-chi
		}(chi)
	}

	return ch
}

// code source:
// https://talks.golang.org/2012/concurrency.slide#27
func main() {
	ch := fanIn(boring("Joe"), boring("Ahn"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("You're both boring, I'm leaving")
}
