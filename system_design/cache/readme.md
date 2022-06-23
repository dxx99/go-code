## 进程内缓存

### 应用场景
- 只读数据、可以考虑在进程启动时加载到内存(配置文件)
- 高并发(秒杀)
- 不一致性要求不高场景(计数场景，运营场景)
- 对性能要求高，redis等网络延迟开销无法满足需求

### 缓存淘汰算法
- FIFO(先进先出)
- LFU(最少使用)
- LRU(最近最少使用)

### 参考的第三方库
- [groupcache](https://github.com/golang/groupcache)
- [bigcache](https://github.com/allegro/bigcache)
- [fastcache](https://github.com/VictoriaMetrics/fastcache)
- [freecache](https://github.com/coocood/freecache)