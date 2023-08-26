package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 当往左查找时，更新上限
// 通过判断上下限来确定是否符合
func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	return lower < root.Val && root.Val < upper && helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}