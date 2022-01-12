package disLock

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
)

type Etcd struct {
	l sync.Mutex
	ctx context.Context
	cli *clientv3.Client
	etcdSession *concurrency.Session
}

func NewEtcd(cli *clientv3.Client) *Etcd {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		panic(err)
	}
	return &Etcd{
		l: sync.Mutex{},
		ctx: context.Background(),
		cli: cli,
		etcdSession: session}
}

func (e *Etcd) Lock(key string) error {
	e.l.Lock()
	defer e.l.Unlock()

	m := concurrency.NewMutex(e.etcdSession, key)
	return m.Lock(e.ctx)
}

func (e *Etcd) Unlock(key string) error {
	e.l.Lock()
	defer e.l.Unlock()

	m := concurrency.NewMutex(e.etcdSession, key)
	return m.Unlock(e.ctx)
}





