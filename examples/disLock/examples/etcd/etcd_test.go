package etcd

import (
	disLock "distributed-lock"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"sync"
	"testing"
	"time"
)

var (
	counter = 0
)

func Test_EtcdLock(t *testing.T)  {
	client, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		panic(err)
	}

	l := disLock.NewLocker(
		disLock.WithLok(disLock.NewEtcd(client)))
	l = l.SetKey("uid:123123")

	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(l *disLock.Lok) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				fmt.Println("start: ", i)
				err := l.Lock()
				if err != nil {
					fmt.Println(err.Error())
				}
				counter++
				err = l.Unlock()
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}(l)
	}
	wg.Wait()
	fmt.Println(counter)
}
