package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := 1 + height(root.Left)
		rightHeight := 1 + height(root.Right)

		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 || leftHeight <= 0 || rightHeight <= 0 { // 小于等于0是因为递归到上一个节点 1 + （-1）=0，这时候已经不符合条件了
			return -1
		} else {
			return max(leftHeight, rightHeight)
		}
	}
	return height(root) > 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
