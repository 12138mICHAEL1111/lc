package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	queue := []*Node{}

	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0{
		size := len(queue)

		for i:=0; i < size ; i++{
			currentNode := queue[0]
			queue = queue[1:]
			if i == size -1 {
				currentNode.Next = nil
			}else{
				currentNode.Next = queue[0]
			}
			if currentNode.Left != nil{
				queue = append(queue,currentNode.Left)
			}
			if currentNode.Right != nil{
				queue = append(queue,currentNode.Right)
			}
		}
	}

	return root
}