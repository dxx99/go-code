package main

import "fmt"

// code source:
// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

type RingBuffer struct {
	inputChannel  <-chan int
	outputChannel chan int
}

func NewRingBuffer(inputChannel <-chan int, outputChannel chan int) *RingBuffer {
	return &RingBuffer{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
	}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			fmt.Printf("丢弃任务：%d\n", <-r.outputChannel)
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}

func main() {
	in := make(chan int)
	out := make(chan int, 5)
	rb := NewRingBuffer(in, out)
	go rb.Run()

	go func() {
		// do task
		for res := range out {
			fmt.Printf("working %d\n", res)
		}
	}()

	// send task
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)

}
