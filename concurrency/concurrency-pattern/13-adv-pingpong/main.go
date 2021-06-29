package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		// receive channel, add one , send channel
		ball := <-table
		ball.hits++
		fmt.Printf("%s --> %d\n", name, ball.hits)
		time.Sleep(time.Millisecond * 1e2)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game start; toss the ball
	time.Sleep(1 * time.Second)
	<-table //game over; grab the ball

	panic("show me the stacks")
}
