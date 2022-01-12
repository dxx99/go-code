package disLock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	RedisPreemptLockFailed = errors.New("Redis failed to preempt lock ")
	RedisUnLockFailed = errors.New("Redis unlock failed ")
)

const (
	RedisLockValue = 1
)

type Redis struct {
	cli *redis.Client
	l sync.Mutex

	ctx context.Context
	timeout time.Duration
}

func NewRedis(cli *redis.Client) *Redis {
	return &Redis{
		cli: cli,
		l:   sync.Mutex{},
		ctx: context.Background(),
		timeout: 10 * time.Second,
	}
}

func (r *Redis) Lock(key string) error {
	r.l.Lock()
	defer r.l.Unlock()

	ok, err := r.cli.SetNX(r.ctx, key, RedisLockValue, r.timeout).Result()
	if err != nil {
		return err
	}

	if !ok {
		return RedisPreemptLockFailed
	}
	return nil
}

func (r *Redis) Unlock(key string) error {
	r.l.Lock()
	defer r.l.Unlock()

	n, err := r.cli.Del(r.ctx, key).Result()
	if err != nil {
		return err
	}
	if n == RedisLockValue {
		return nil
	}
	return RedisUnLockFailed
}
