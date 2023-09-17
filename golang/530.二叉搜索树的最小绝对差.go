package main

import "math"

// 不存在右节点-左节点比根节点-左节点小的情况
// 所以用中序遍历， 不用运算左右节点，只用运算根节点和子节点
func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt64
	var pre *TreeNode
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return 
		}
		inOrder(root.Left)
		if pre != nil  &&  root.Val - pre.Val < min{
			min = root.Val - pre.Val 
		}
		pre = root
		inOrder(root.Right)
	}
	inOrder(root)
	return min
}