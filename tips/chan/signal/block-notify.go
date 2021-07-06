package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	// block until received signal
	s := <-ch
	fmt.Println("get signal: ", s)
}
