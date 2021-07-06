## 带信号的服务启动模式

### 接口定义
```go
package serv
type Service interface {

	// Init 在服务启动的时候，进行初始化
	Init(Environment) error

	// Start 启动服务
	Start() error

	// Stop 停止服务，通过syscall.signal
	Stop() error
}

// Environment 服务启动环境
type Environment interface {
	IsWindowService() bool
}
```

### 实现接口
- 通过ctx 和 signal来控制服务退出
```go
	// 通过ctx或sig信号关闭
	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case os.Interrupt, os.Kill, syscall.SIGABRT, syscall.SIGTERM:
				log.Printf("receive exit signal, sig=%v\n", sig)
				return serv.Stop()
			case syscall.SIGURG:
				log.Println("sigurg")
				runtime.Gosched()
			default:	//io信号直接退出：https://github.com/golang/go/issues/38290
				log.Printf("unhandled signal: %v\n", sig)
			}
		case d := <-ctx.Done():
			log.Printf("timout exit, done=%v\n", d)
			return serv.Stop()
		}
	}
```