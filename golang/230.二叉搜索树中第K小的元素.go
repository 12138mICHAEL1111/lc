package main

import "math"

func kthSmallest(root *TreeNode, k int) int {
	time := 0
	ans := math.MinInt64
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil{
			return
		}

		inOrder(root.Left)
		if time == k {
			return
		} else {
			time++
		}

		if root.Val > ans {
			ans = root.Val
		}
		inOrder(root.Right)
	}
	inOrder(root)
	return ans
}