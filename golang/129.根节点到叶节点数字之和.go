package main

func sumNumbers(root *TreeNode) int {
	return dfsSum(root,0)
}

func dfsSum(root *TreeNode, total int) int {
	if root == nil {
		return 0
	}

	total = total * 10 + root.Val

	if root.Left == nil && root.Right == nil {
		return total
	}

	return dfsSum(root.Left,total) + dfsSum(root.Right,total)
}