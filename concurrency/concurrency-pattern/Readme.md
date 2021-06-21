## go并发模型

### 1. go func
- go func 直接协程运行
- [示例](1-boring/main.go)

### 2. go channel
- 通过 channel 将协程的结果返回给 main 协程
- [示例](2-chan/main.go)
- [示例-2](3-generator/main.go)

### 3. fanIn
- 合并两个channel的结果，输入到一个channel
- [示例](4-fanin/main.go)

### 4. struct-channel 
- 通过struct里面嵌套channel保证两个任务同步进行
- [示例](5-resotre-sequence/main.go)

### 5. select-timeout
- 通过select来进行超时处理
- [示例](6-select-timeout/main.go)

### 6. channel-exit
- 通过channel控制子协程退出
- [示例](7-quit-signal/main.go)

### 7. daisy-channel
- 链式channel, 链式传递channel
- [示例](8-daisy-chain/main.go)

### 8. google-search
- 串行请求(Web, Video, Image)
    - [示例](9-google-1.0/main.go)
- 并行请求
    - 使用一个channel接收三个goroutine的返回值
    - [示例](10-google-2.0/main.go)
- 并行请求(超时控制)
    - 按并行请求，然后添加select来进行超时控制
    - [示例](11-google-2.1/main.go)
- 并行请求(超时控制+多节点部署)，返回每个服务最快响应
    - 每个节点使用同一个channel来接收，保留最快的结果
    - [示例](12-google-3.0/main.go)
    
### 9. goroutine-communication
- goroutine之间的通信
- [示例](13-adv-pingpong/main.go)

### 10. waitGroup 
- waitGroup协程处理
- [示例](15-bounded-paraller/main.go)

### 11. channel-context
- 使用context与select进行级联控制
- [示例-server](16-context/server.go)
- [示例-client](16-context/client.go)

### 12. ring-buffer-channel
- 环形buffer的channel
- [示例](17-ring-buffer-channel/main.go)

### 13. worker-pool
- 工作池模式
- [示例](18-worker-pool/main.go)
