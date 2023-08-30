package main
//
//    2
// 1     3
//4  5 null 8
//先交换 4 5， 再交换 1 3
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rightTree := root.Right

	root.Right = invertTree(root.Left)
	root.Left = invertTree(rightTree)

	return root
}