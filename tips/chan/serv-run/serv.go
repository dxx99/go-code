package serv

import (
	"context"
	"os/signal"
)

// Create variable signal.Notify function so we can mock it in tests
var signalNotify = signal.Notify

type Environment interface {
	IsWindowService() bool
}

type Service interface {

	// Init 在服务启动的时候，进行初始化
	Init(Environment) error

	// Start 启动服务
	Start() error

	// Stop 停止服务，通过syscall.signal
	Stop() error
}

// Context interface contains an optional Context function which a Service can implement.
// When implemented the context.Done() channel will be used in addition to signal handling
// to exit a process.
// 用于任务的退出
type Context interface {
	Context() context.Context
}
