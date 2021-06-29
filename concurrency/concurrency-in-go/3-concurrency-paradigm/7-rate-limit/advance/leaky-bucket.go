package advance

import (
	"math"
	"time"
)

//使用channel 与 tick实现漏桶算法

type Limit float64

const Inf = Limit(math.MaxFloat64)

// Every 每秒钟能够通过多少请求
func Every(interval time.Duration) Limit {
	if interval <= 0 {
		return Inf
	}
	return 1 / Limit(1 * interval.Seconds())
}

// EveryTime 每个请求需要的时间
func EveryTime(num int64) time.Duration {
	if num <= 0 {
		return time.Duration(math.MaxInt64)
	}
	return time.Second / time.Duration(num)
}

// Payload 请求消息
type Payload struct {

}

// Client 定义执行客户端
type Client interface {
	Call(*Payload)
}

func RateLimitCall(clinet Client, payloads []*Payload)  {
	t := time.Tick(EveryTime(10))
	for _, payload := range payloads {
		<-t
		go clinet.Call(payload)
	}
}
