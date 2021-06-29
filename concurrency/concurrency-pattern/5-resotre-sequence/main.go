// send a channel on a channel, making goroutine wait its turn
// receive all messages, then enable them again by sending on a private channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message define a message type that contains a channel for the reply
type Message struct {
	str  string
	wait chan bool
}

// fanIn message many chan Message
func fanIn(inputs ...<-chan Message) <-chan Message {
	ch := make(chan Message)
	for _, input := range inputs {
		go func(input <-chan Message) {
			for {
				ch <- <-input
			}
		}(input)
	}
	return ch
}

func boring(msg string) <-chan Message {
	ch := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			ch <- Message{
				str:  fmt.Sprintf("%s --> %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// every time the goroutine send message
			// this code waits until the value to be received
			<-waitForIt
		}
	}()

	// return a channel to caller
	return ch
}

// code source:
// https://talks.golang.org/2012/concurrency.slide#30
func main() {
	ch := fanIn(boring("Joe"), boring("Ahn"))

	// main: goroutine                                          boring: goroutine
	//    |                                                           |
	//    |                                                           |
	// wait for receiving msg from channel c                    c <- Message{} // send message
	//   <-c                                                          |
	//    |                                                           |
	//    |                                                     <-waitForIt // wait for wake up signal
	// send value to channel                                          |
	// hey, boring. You can send next value to me                     |
	//   wait <-true                                                  |
	///                            REPEAT THE PROCESS
	for i := 0; i < 5; i++ {
		msg1 := <-ch
		fmt.Println(msg1.str)
		msg2 := <-ch
		fmt.Println(msg2.str)

		// each goroutine have to wait
		// main goroutine allows the boring goroutine to send next value to message channel
		msg1.wait <- true
		msg2.wait <- true
	}

}
