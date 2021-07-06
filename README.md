# go编程学习
Collecting code related snippets for easy development

## 并发编程
### [concurrency-pattern](concurrency/concurrency-pattern/Readme.md)
- [1-bubble](concurrency/concurrency-pattern/1-boring/main.go)
- [2-channel](concurrency/concurrency-pattern/2-chan/main.go)
- [3-generator](concurrency/concurrency-pattern/3-generator/main.go)
- [4-fanIn](concurrency/concurrency-pattern/4-fanin/main.go)
- [5-restore-sequence](concurrency/concurrency-pattern/5-resotre-sequence/main.go)
- [6-select-timeout](concurrency/concurrency-pattern/6-select-timeout/main.go)
- [7-quit-signal](concurrency/concurrency-pattern/7-quit-signal/main.go)
- [8-daisy-chain](concurrency/concurrency-pattern/8-daisy-chain)
- [9-google-search](concurrency/concurrency-pattern/12-google-3.0/main.go)
- [10-bounded-paraller](concurrency/concurrency-pattern/15-bounded-paraller/main.go)
- [11-context](concurrency/concurrency-pattern/16-context/main.go)
- [12-ring-buffer-channel](concurrency/concurrency-pattern/17-ring-buffer-channel/main.go)
- [13-worker-pool](concurrency/concurrency-pattern/18-worker-pool/main.go)

### [concurrency-in-go](concurrency/concurrency-in-go/Readme.md)
- 锁的相关问题
    - [死锁](concurrency/concurrency-in-go/1-introduction/deadlock)：所有的并发进程都彼此等待
    - [活锁](concurrency/concurrency-in-go/1-introduction/livelock)：正在主动执行并发操作，但这些操作无法向前移动
    - [饥饿](concurrency/concurrency-in-go/1-introduction/hungerlock)：并发进程无法获得执行工作所需要的任何资源
- 低级别的同步原语
    - [WaitGroup](concurrency/concurrency-in-go/2-concurrency-module/2-waitgroup)
    - [Mutex和RWMutex](concurrency/concurrency-in-go/2-concurrency-module/3-mutex-rwmutex)
    - [Cond](concurrency/concurrency-in-go/2-concurrency-module/4-cond)
    - [Once](concurrency/concurrency-in-go/2-concurrency-module/5-once)
    - [Pool](concurrency/concurrency-in-go/2-concurrency-module/6-pool) 
- 高级别的同步
    - [select](concurrency/concurrency-in-go/2-concurrency-module/7-select)
    - [for-select循环](concurrency/concurrency-in-go/3-concurrency-paradigm/1-for-select/go-for-select.go)

### 高级话题
- [速率限制](concurrency/concurrency-in-go/3-concurrency-paradigm/7-rate-limit/Readme.md)


## [数据结构与算法](./algorithm/Readme.md)
### 数据结构
- 链表
  - [单链表](./algorithm/linkList/single-linked-table/single-link.go)
  - [双链表](./algorithm/linkList/double-linked-table/double-link.go)
  - [循环单链表](./algorithm/linkList/cycle-single-link/cycle-single-linke.go)
- 栈
  - [数组栈](./algorithm/stack/slice-stack/slice-stack.go) 
- 队列
  - [环形队列](./algorithm/queue/cycle-queue/cycle-queue.go)
  

### 算法实现
- [排序](./algorithm/sort/Readme.md)
    - [冒泡排序](./algorithm/sort/1-bubble/main.go)
    - [选择排序](./algorithm/sort/2-selection/main.go)
    - [插入排序](./algorithm/sort/3-insertion/main.go)
    - [希尔排序](./algorithm/sort/4-shell/main.go)
    - [归并排序](./algorithm/sort/5-merge/main.go)
    - [快速排序](./algorithm/sort/6-quick/main.go)
    - [桶排序](./algorithm/sort/7-bucket/main.go)
    - [计数排序](./algorithm/sort/8-count/main.go)
    - [基数排序](./algorithm/sort/9-radix/main.go)
    - [堆排序](./algorithm/sort/10-heap/main.go)

## 问答收集
### [官方faq](faq/official/Readme.md)

## [编程小技巧](./tips/Readme.md)
### channel
- [tryLock](./tips/chan/trylock/trylock.go)
- [服务启动](./tips/chan/serv-run/Readme.md)

### slice
- [切片基本操作](./tips/slice/baseop/main.go)

### interface
- interface类型转换
    - [转换为T](./tips/interface/interface-implements/convert_T_same_underlying_type.go)
    - [验证接口](./tips/interface/interface-implements/implements-verify.go)


### int 
- [整型溢出](./tips/int/overflow.go)

