## go基础
### 数组与切片
- 对比
  - 数组为固定长度，不能修改，是一片连续的内存
  - 切片是一个结构体，含有三个字段(长度，容量，底层数组)
    - 底层数组是一个指针 `unsafe.Pointer`
    - len切片元素的数量，cap切片容量
- 切片扩容逻辑
  - 如果传入的cap大于doublecap，则直接使用传入的cap
    - `append([]int{}, 1,2,3,4,5)` 操作后，len=5, cap=6
  - v1.18之前
    - `old.cap < 1024` 2倍老容量 `newcap = doublecap`
    - `old.cap > 1024` 1.25倍老容量 `newcap = 1.25*old.cap`
  - v1.18之后
    - `old.cap < 256` 2倍老容量 `newcap = doublecap`
    - `old.cap > 256` `newcap = old.cap + (old.cap+3*256)/4`
  - 后面也会进行内存对齐操作
- 切片作为函数的参数
  - 当切片作为一个参数的时候，其实就像一个结构体作为函数的参数
  - 注意：
    - slice底层数据在数组上，slice存储的是一个指针
    - 传参的时候，尽快slice的数据没发生改变，也就是指针没变
    - 但是`s[i]=10`这种操作直接改变了数组的值，因此内容也发生了变化

### 哈希表
![img.png](images/gobase/hash/1.png)
#### 哈希函数
- 分类
  - 加密类，md5、sha1、sha256
  - 非加密类，查找类
- 考察点
  - 性能
  - 碰撞概率
#### map的get的两种操作
