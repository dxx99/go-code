package trylock

type TryLock struct {
	ch chan struct{}
}

func NewTryLock() *TryLock {
	return &TryLock{ch: make(chan struct{}, 1)}
}

func (t *TryLock) Lock() bool {
	select {
	case t.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (t *TryLock) Unlock() bool {
	select {
	case <-t.ch:
		return true
	default:
		return false
	}
}
