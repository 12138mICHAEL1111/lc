package main

import (
	"math"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	mostLeftHeight := 0
	mostRightHeight := 0
	node := root
	for node.Left != nil {
		mostLeftHeight++
		node = node.Left
	}

	node = root
	for node.Right != nil {
		mostRightHeight++
		node = node.Right
	}

	if mostLeftHeight == mostRightHeight {
		return  int(math.Pow(2.0,float64(mostLeftHeight+1))) - 1
	}else{
		return 1 + countNodes(root.Left) + countNodes((root.Right))
	}
}