package container

type MonitorQueueInterface interface {
	Push(x interface{})
	Max() interface{}
	Pop(x interface{})
}

type MonitorQueue struct {
	bucket []interface{}
}

func (m MonitorQueue) Push(x interface{}) {
	if len(m.bucket) == 0 {
		m.bucket = append(m.bucket, x)
	}
}

func (m MonitorQueue) Max() interface{} {
	//TODO implement me
	panic("implement me")
}

func (m MonitorQueue) Pop(x interface{}) {
	//TODO implement me
	panic("implement me")
}
