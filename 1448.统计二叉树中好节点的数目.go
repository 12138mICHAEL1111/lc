package main

import "math"

func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int)
	ans := 0
	dfs = func(root *TreeNode, pathMax int) {
		if root == nil {
			return
		}

		if root.Val > pathMax {
			ans++
			pathMax = root.Val
		}
		dfs(root.Left, pathMax)
		dfs(root.Right, pathMax)
	}
	dfs(root, math.MinInt)
	return ans
}