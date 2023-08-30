package main

func buildTree2(inorder []int, postorder []int)*TreeNode{
	inorderMap := make(map[int]int)
	for k,v:= range inorder{
		inorderMap[v] = k
	}
	var build func(int, int) *TreeNode
	rootIndex := len(postorder) - 1 
	build = func(left int, right int) *TreeNode {
		if(left <= right){
			  rootVal := postorder[rootIndex]
			  rootIndex --
			  node := &TreeNode{Val: rootVal}
			  node.Left = build(left,inorderMap[rootVal]-1) // 就这里不一样，先遍历right
			  node.Right = build(inorderMap[rootVal] + 1,right)
			  return node
		}else{
			return nil
		}
	}
	return build(0, len(postorder)-1)
}