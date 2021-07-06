package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"serv"
	"sync"
	"time"
)

type program struct {
	LogFile *os.File
	serv *server
	ctx context.Context
}

func NewProgram(ctx context.Context, f *os.File) *program {
	return &program{
		LogFile: f,
		serv:    NewServer(),
		ctx:     ctx,
	}
}

// Close 关闭文件输入错误
func (p *program) Close() {
	if p.LogFile != nil {
		if closeErr := p.LogFile.Close(); closeErr != nil {
			log.Printf("error closing '%s': %v\n", p.LogFile.Name(), closeErr)
		}
	}
}

func (p *program) Init(env serv.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowService())

	// write to "example.log" when running as a Windows Service
	if env.IsWindowService() {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}

		logPath := filepath.Join(dir, "example.log")

		f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		p.LogFile = f

		log.SetOutput(f)
	}
	return nil
}

func (p *program) Start() error {
	log.Printf("Starting...\n")
	go func() {
		err := p.serv.start()
		if err != nil {
			log.Printf("start failed, reason: %v\n", err.Error())
		}
	}()
	return nil
}

func (p *program) Stop() error {
	log.Printf("Stopping...\n")
	if err := p.serv.stop(); err != nil {
		log.Printf("stop failed, reason: %v\n", err.Error())
		return err
	}
	log.Printf("Stopped.\n")
	return nil
}

func (p *program) Context() context.Context {
	return p.ctx
}


type server struct {
	data chan int

	done chan struct{}
	wg sync.WaitGroup
}

func NewServer() *server {
	return &server{
		data: make(chan int),
		done: make(chan struct{}),
		wg:   sync.WaitGroup{},
	}
}


func (s *server) start() error {
	s.wg.Add(2)
	go s.startReceiver()
	go s.startSender()
	s.wg.Wait()
	return nil
}

func (s *server) startSender() {
	ticker := time.NewTicker(20 * time.Millisecond)
	defer s.wg.Done()
	count := 1
	for {
		select {
		case <-ticker.C:
			select {
			case s.data <- count:
				count++
			case <-s.done:
				return
			}
		case <-s.done:
			return
		}
	}
}

func (s *server) startReceiver() {
	defer s.wg.Done()

	for {
		select {
		case n := <-s.data:
			log.Printf("%d\n", n)
		case <-s.done:
			return
		}
	}
}

func (s *server) stop() error {
	close(s.done)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	p := NewProgram(ctx, nil)
	defer p.Close()

	if err := serv.Run(p); err != nil {
		log.Fatal(err)
	}
}
