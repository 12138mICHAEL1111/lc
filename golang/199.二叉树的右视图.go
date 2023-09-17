package main

func rightSideView(root *TreeNode) []int {
	result := []int{}
	var dfs func (*TreeNode,int) 
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}

		if height > len(result){
			result = append(result, root.Val)
		}

		dfs(root.Right,height+1)
		dfs(root.Left,height+1)
	}
	dfs(root,1)
	return result

	// 从下往上
	// 就插入在数组最前面
}
