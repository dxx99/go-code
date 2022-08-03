package core

import (
	"context"
	"fmt"
)

type worker struct {
}

func NewWorker() *worker {
	return &worker{}
}

func (w *worker) Run(ctx context.Context) {
	fmt.Println("worker running...")
	fmt.Printf("worker context is %s\n", ctx.Value("parent").(string))
	select {
	case <-ctx.Done():
		fmt.Println("child goroutine exit!", ctx.Err().Error())

	}
}
