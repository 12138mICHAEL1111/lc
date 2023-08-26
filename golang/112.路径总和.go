package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root,0,targetSum)
}

func dfs(root *TreeNode, sum,targetSum int) bool{
	// 只有左节点或右节点，或第一个根节点
	if root == nil{
		 return false
	}

	//叶子节点
	if root.Left == nil && root.Right==nil{
		sum = sum + root.Val
		if sum == targetSum{
			return true
		}else{
			return false
		}
	}


	return dfs(root.Left,sum+root.Val,targetSum) || dfs(root.Right,sum+root.Val,targetSum)
}