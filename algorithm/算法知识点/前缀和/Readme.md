## 前缀和

### 构建条件
- 所有的元素没有负数
- 通过前缀和构建一个递增的数组，然后可以利用二分查找

### 示例代码
```go
func prefix(nums []int ){
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
}
```

### leetCode真题
- [1894. 找到需要补充粉笔的学生编号](https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk/)