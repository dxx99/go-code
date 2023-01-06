package container

type MonitorQueueInterface interface {
	Push(x interface{})
	Max() interface{}
	Pop(x interface{})
}

// MonitorQueue 从尾部到头部单调递增
type MonitorQueue struct {
	bucket []interface{}
}

func (m *MonitorQueue) Push(x interface{}) {
	for len(m.bucket) != 0 || m.CmpLast(x) {
		m.bucket = m.bucket[:len(m.bucket)-1]
	}
	m.bucket = append(m.bucket, x)
}

func (m *MonitorQueue) CmpLast(x interface{}) bool {
	last := m.bucket[len(m.bucket)-1]
	switch x.(type) {
	case int:
		return last.(int) < x.(int)
	case int8:
		return last.(int8) < x.(int8)
	case int16:
		return last.(int16) < x.(int16)
	case int32:
		return last.(int16) < x.(int16)
	case int64:
		return last.(int16) < x.(int16)
	case uint8:
		return last.(uint8) < x.(uint8)
	case uint16:
		return last.(uint16) < x.(uint16)
	case uint32:
		return last.(uint32) < x.(uint32)
	case uint64:
		return last.(uint64) < x.(uint64)
	case float32:
		return last.(float32) < last.(float32)
	case float64:
		return last.(float64) < last.(float64)
	default:
		return false
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
