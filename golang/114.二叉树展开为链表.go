package main

// 右节点 -> 左节点 -> 根节点
func flatten(root *TreeNode)  {
	var solve func(*TreeNode)
	var preNode *TreeNode
	solve = func(root *TreeNode) {
		if root == nil {
			return 
		}

		solve(root.Right)
		solve(root.Left)

		root.Left = nil 
		root.Right = preNode
		preNode = root
	}
	solve(root)
}