## concurrency in go

### 介绍
- 概念介绍
    - 并发
    - 并行
    - 异步
    - 线程化
- 并发困难的原因
    - 数据竞争
        - 内存访问同步(访问同一内存区域)
    - 原子性
    - 锁的问题
        - [死锁](./1-introduction/deadlock)：所有的并发进程都彼此等待
        - [活锁](./1-introduction/livelock)：正在主动执行并发操作，但这些操作无法向前移动
        - [饥饿](./1-introduction/hungerlock)：并发进程无法获得执行工作所需要的任何资源
    - 并发安全性
- 科父曼条件(帮助检测死锁)
    - 相互排斥：并发进程在任何时候都拥有资源的独占权
    - 等待条件：并发进程必须同时拥有资源并等待额外的资源
    - 没有抢占：并发进程持有该资源只能由该进程释放
    - 循环等待：并发P1等待P2, 同时p2也等待p1, 因此符合循环等待
    
### 代码建模
- 并发与并行

- csp并发模型

### 并发构建模块
- [goroutines](./2-concurrency-module/1-goroutines)
    - 非抢占式的并发子程序，不能被中断，有runtime管理
    - runtime自动挂起它们，然后恢复它们

- sync(低级别内存访问同步)
    - [WaitGroup](./2-concurrency-module/2-waitgroup)
    - [Mutex和RWMutex](./2-concurrency-module/3-mutex-rwmutex)
    - [Cond](./2-concurrency-module/4-cond)
    - [Once](./2-concurrency-module/5-once)
    - [Pool](./2-concurrency-module/6-pool) ???

- channels

- [select](./2-concurrency-module/7-select)
  - 取消  
  - 超时
  - 等待
  - 默认值 default

- [GOMAXPROCS](./2-concurrency-module/8-gomaxprocs)

### 并发编程范式
- 访问范围约束
  - 共享内存同步原语(eg. sync.Mutex)
  - 通过通信同步(eg. channel)
  - 不可变数据(静态分析，很难)
  - 受限制条件保护的数据

- [for-select循环](./3-concurrency-paradigm/1-for-select/go-for-select.go)
  - 在channel发送迭代变量
  - 无线循环等待
  
- 防止goroutine泄露
  - goroutine终止的可能
    - 当它完成任务
    - 当它遇到不可恢复的错误无法继续它的任务
    - 当它被告知停止当前任务

### 可伸缩并发设计

### goroutine in runtime
