[TOC]

## 二叉树

### 递归
**三要数**
- 参数、返回值
- 终止条件
- 单层逻辑

**递归的本质**
- 每一次递归调用都会把函数的局部变量、参数值和返回值地址等压入调用栈中
- 返回时，从栈顶弹出上一次递归的各项参数，所以这就是递归可以返回上一层位置的原因

### [遍历](./二叉树遍历/main.go)
![](./images/二叉树前中序遍历.png)

#### 递归遍历
- [前序遍历](./二叉树遍历/main.go#L16)
- [中序遍历](./二叉树遍历/main.go#L34)
- [后序遍历](./二叉树遍历/main.go#L53)

#### 迭代遍历
- [前序遍历](./二叉树遍历/main.go#L74)
- [中序遍历](./二叉树遍历/main.go#L104)
- [后序遍历](./二叉树遍历/main.go#L130)

### 层序遍历
- [bfs](./二叉树遍历/main.go#L162)

### 树的递归函数的注意点
#### 1. 怎么确定树的叶子节点
```go
if node.Left == nil && node.Right == nil {
	//todo
}
```
#### 2.怎么确定最左子节点
```go
if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
	//todo
}
```
#### 3. 怎么得到满二叉树的高度
```go
h := 0
for node != nil {
	node = node.Left
	h++
}
//todo
```
#### 4. 怎么判断当前节点下面的子树是满二叉树
```go
leftHeight, rightHeight := 0, 0
for lNode != nil {
	lNode = lNode.Left
	leftHeight++
}
for rNode != nil {
	rNode = rNode.Right
	rightHeight++
}
if leftHeight == rightHeight {
	//todo
}
```
#### 5. 什么时候需要返回值
- 如果需要遍历整棵树，递归函数就不能有返回值
- 如果需要遍历某一条固定路线，递归函数就一定要有返回值



### [LeetCode二叉树层序遍历](./力扣层序遍历/main.go)
- [102.二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/)
- [107.二叉树的层次遍历II](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/)
- [199.二叉树的右视图](https://leetcode.cn/problems/binary-tree-right-side-view/)
- [637.二叉树的层平均值](https://leetcode.cn/problems/average-of-levels-in-binary-tree/)
- [429.N叉树的层序遍历](https://leetcode.cn/problems/n-ary-tree-level-order-traversal/)
- [515.在每个树行中找最大值](https://leetcode.cn/problems/find-largest-value-in-each-tree-row/)
- [116.填充每个节点的下一个右侧节点指针](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/)
- [117.填充每个节点的下一个右侧节点指针II](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/)
- [104.二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
- [111.二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)

### [二叉树的修改与构造](./二叉树的修改与构造/main.go)
- [226.反转二叉树](https://leetcode.cn/problems/invert-binary-tree/)
- [101.对称二叉树](https://leetcode.cn/problems/symmetric-tree/)
- [222. 完全二叉树的节点个数](https://leetcode.cn/problems/count-complete-tree-nodes/)
- [110. 平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree/)
- [257. 二叉树的所有路径](https://leetcode.cn/problems/binary-tree-paths/)
- [404.左叶子节点之和](https://leetcode.cn/problems/sum-of-left-leaves/)
- [513. 找树左下角的值](https://leetcode.cn/problems/find-bottom-left-tree-value/)
- [112. 路径总和](https://leetcode.cn/problems/path-sum/)
- [113. 路径总和 II](https://leetcode.cn/problems/path-sum-ii/)
- 