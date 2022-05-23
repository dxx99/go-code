package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 贪心算法求解
// 局部最优解：让叶子节点的父节点安装摄像头，所用的摄像头最少
// 整体最优：全部摄像头数量最少
// 题目难点：
//	1. 二叉树遍历
// 	2. 如何隔两个节点放一个摄像头
//
// todo:
func minCameraCover(root *TreeNode) int {
	return 0
}
