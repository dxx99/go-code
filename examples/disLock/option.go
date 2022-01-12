package disLock

type Option func(o *Options)

type Options struct {
	lok Locker
}

func WithLok(lok Locker) Option {
	return func(o *Options) {
		o.lok = lok
	}
}


