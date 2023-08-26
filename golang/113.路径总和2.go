package main

func pathSum(root *TreeNode, targetSum int) [][]int {
    path := []int{}
	ans := [][]int{}
    var dfs func(*TreeNode, int,int)
	dfs = func(root *TreeNode, sum,targetSum int) {
		if root == nil{
			return 
	   }
	   path = append(path, root.Val)
	   if root.Left == nil && root.Right==nil{
		   sum = sum + root.Val
		   if sum == targetSum{
			ans = append(ans, append([]int(nil), path...))
		   }
	   } 
		dfs(root.Left,sum+root.Val,targetSum)  
		dfs(root.Right,sum+root.Val,targetSum)
		//移除
		path = path[:len(path)-1 ]
	}
	dfs(root,0,targetSum)
	return ans
}

