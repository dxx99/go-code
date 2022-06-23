package cache

import (
	"container/list"
	"fmt"
	"runtime"
)

type Cache interface {
	Set(key string, value interface{})	// 设置/添加一个缓存，如果 key 存在，用新值覆盖旧值
	Get(key string) interface{}			// 通过 key 获取一个缓存值
	Del(key string)						// 通过 key 删除一个缓存值
	DelOldest()							// 删除最无用的一个缓存值
	Len() int							// 获取缓存已存在的记录数
}

type fifo struct {
	maxBytes int		// 缓存的容量，单位字节，最大存放的entry个数

	onEvicted func(key string, value interface{})	// 当缓存中移除调用该回调函数

	usedBytes int 	// 已使用的字节数，只包括值，key不算

	ll *list.List

	cache map[string]*list.Element
}

// New 创建一个新的cache, 如果没有设置maxBytes， 这表示没有容量的限制
func New(maxBytes int, onEvicted func(key string, value interface{})) Cache {
	return &fifo{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		usedBytes: 0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}


func (f *fifo) Set(key string, value interface{}) {
	// 对已存在的值进行更新
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - CalLen(en.value) + CalLen(value)
		en.value = value
		return
	}

	en := &entry{key,value}
	e := f.ll.PushBack(en)
	f.cache[key] = e


	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}

func (f fifo) Get(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (f fifo) Del(key string) {
	//TODO implement me
	panic("implement me")
}

func (f fifo) DelOldest() {
	//TODO implement me
	panic("implement me")
}

func (f fifo) Len() int {
	//TODO implement me
	panic("implement me")
}

type entry struct {
	key string
	value interface{}
}

func (e *entry) Len() int {
	return CalLen(e.value)
}

type Value interface {
	Len() int
}

func CalLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, uint8, int8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		}else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.Value", value))
	}
	return n
}

