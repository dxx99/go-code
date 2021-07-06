package serv

import (
	"context"
	"log"
	"os"
	"runtime"
	"syscall"
)

type environment struct {

}

func (e environment) IsWindowService() bool {
	return false
}

func Run(serv Service, sig ...os.Signal) error {
	env := environment{}

	// 初始化操作
	if err := serv.Init(env); err != nil {
		return err
	}

	// 服务启动
	if err := serv.Start(); err != nil {
		return err
	}

	signalChan := make(chan os.Signal, 1)
	signalNotify(signalChan, sig...)

	ctx := context.Background()
	if s, ok := serv.(Context); ok {
		ctx = s.Context()
	}

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

}
