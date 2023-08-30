package main
//二叉树深度：从当前节点到根节点
//高度：从当节点到最长的叶子节点
func maxDepth(root *TreeNode) int {
	if root != nil {
		return 0
	}
	leftMax := 1 + maxDepth(root.Left)
	rightMax := 1 + maxDepth(root.Right)
	if leftMax >= rightMax{
		return leftMax
	}else{
		return rightMax
	}
}