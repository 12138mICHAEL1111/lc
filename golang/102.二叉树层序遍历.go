package main
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{}
	if root!= nil {
		queue=append(queue, root)
	}
	for len(queue) > 0{
		size := len(queue)
		nodeList := []int{}
		for i:=0;i<size;i++{ //一次性遍历完同一层
			// 取出并删除queue第一个元素
			node:= queue[0]
			nodeList = append(nodeList,node.Val)
			queue=queue[1:]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
		}
		result = append(result, nodeList)
	}
	return result
}