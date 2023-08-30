package main
func buildTree(preorder []int, inorder []int) *TreeNode {
	rootIndex := 0 //每次递增1, 相当于遍历一次preorder
	inOrderMap :=  make(map[int]int)
	for k, v := range inorder{
		inOrderMap[v] = k
	}
	var build func(int,int)*TreeNode
	build = func(left, right int) *TreeNode{
		if left <= right{ // 背
			rootVal := preorder[rootIndex]
			rootIndex ++ 
			node := &TreeNode{Val: rootVal}
			node.Left = build(left,inOrderMap[rootVal]-1)
			node.Right = build(inOrderMap[rootVal] + 1,right)
			return node
		}else{
			return nil
		}
	}

	return build(0, len(preorder)-1)
}