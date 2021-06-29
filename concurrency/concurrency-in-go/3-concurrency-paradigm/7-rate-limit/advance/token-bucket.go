package advance

import (
	"context"
	"time"
)

//使用channel 与 tick实现token算法

// BurstNum 并发的数量
const BurstNum = 10

func BurstRateLimitCall(ctx context.Context, client Client,  payloads []*Payload)  {
	ch := make(chan time.Time, BurstNum)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rateLimit := EveryTime(10)

	// 固定的时间往协程里面写数据
	go func() {
		ticker := time.NewTicker(rateLimit)
		defer ticker.Stop()
		for t := range ticker.C {
			select {
			case ch <- t:
			case <-ctx.Done():
				return //终止协程
			}
		}
	}()

	for _, payload := range payloads {
		<-ch
		go client.Call(payload)
	}
}
