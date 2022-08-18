## 二分查找

### 1. 左闭右闭查找
```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1
		}else if nums[mid] > target {
			right = mid-1
		}else {
			return mid	//找到，直接退出
		}
	}
	return -1	//没有找到
}
```

### 2. 查找左边界
- **[left, right) 左闭右开版本**
```go
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {		// 相等之后就不会右移，因此找到的是左边界
			left = mid+1
		}else {
			right = mid
		}
	}

	// 找不到返回-1，优化返回结果集
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
```

- **[left, right] 左闭右闭版本**
```go
func leftBoundV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1			// [mid+1, right]
		}else if nums[mid] > target {
			right = mid-1			// [left, mid-1]
		}else if nums[mid] == target {
			right = mid-1			// 【注意】: 收缩右边界，这样才能退出
		}
	}

	// 找不到返回-1，优化返回结果集
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
```

### 3. 查找右边界
- **[left, right) 左闭右开版本**
```go
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := int(uint(left+right)>>1)
		if nums[mid] <= target {	// 注意这里的等号，也需要右移
			left = mid+1
		}else {
			right = mid
		}
	}

	// 找不到返回-1，优化返回结果集
	if left <= 0 || nums[left-1] != target {
		return -1
	}
	return left-1
}
```

- **[left, right] 左闭右闭版本**
```go
func rightBoundV2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := int(uint(left+right)>>1)
		if nums[mid] < target {
			left = mid+1
		}else if nums[mid] > target {
			right = mid-1
		}else if nums[mid] == target {
			left = mid+1		// 【注意】: 收缩左边界，这样才能退出
		}
	}

	// 找不到返回-1，优化返回结果集
	if left <= 0 || nums[left-1] != target {
		return -1
	}
	return left-1
}
```


### 4. 二分搜索的套路框架
- `通过用来求解最值问题`
```go
// 函数f关于自变量x的单调函数
func f(int x) int {
	//....
}

// 主函数，在 f(x) == target 的约束下求 x 的最值
func solution(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// 问自己：自变量 x 的最大值是多少？自变量 x 的最小值是多少？
	left, right := 1, int(1e9)
	for left < right {
		mid := int(uint(left+right)>>1)
		if	f(mid) == target {
			// 问自己：题目是求左边界还是右边界？
		} else if f(mid) > target {
			// 问自己：怎么让 f(x) 小一点？
		}else if f(mid) < target {
			// 问自己：怎么让 f(x) 大一点？		
		}
	}	
	
	return left
}

```

