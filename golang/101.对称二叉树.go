package main
func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool 
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil{
			return true
		}

		if left == nil || right == nil { // 因为已经运行过上一个if，所以这里不会出现两个都是nil的情况
			return false
		}

		if left.Val != right.Val{
			return false
		}

		return dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
	}
	if root==nil{
		return false
	}

	return dfs(root.Left,root.Right)
}