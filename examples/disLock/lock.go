package disLock

import (
	"errors"
	"sync"
	"time"
)

var (
	KeyEmpty = errors.New("locker key is empty ")
)

type Locker interface {
	Lock(key string) error
	Unlock(key string) error
}

type Lok struct {
	opts Options

	mutex sync.Mutex
	key string
	timeout time.Duration
}

func NewLocker(opts ...Option) *Lok {
	o := Options{}

	for _, opt := range opts {
		opt(&o)
	}

	lok :=  &Lok{
		opts: o,
		timeout: time.Second,
	}
	return lok
}

func (l *Lok) SetKey(key string) *Lok {
	l.key = key
	return l
}

func (l *Lok) Lock() error {
	if l.key == "" {
		return KeyEmpty
	}
	return l.opts.lok.Lock(l.key)
}

func (l *Lok) Unlock() error {
	if l.key == "" {
		return KeyEmpty
	}
	return l.opts.lok.Unlock(l.key)
}





