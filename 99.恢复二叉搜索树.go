package main

import "fmt"

//先中序遍历得到list，查找list里顺序错误的数字，再交换数字
func recoverTree(root *TreeNode) {
	var list []int
	inOrder(root, &list)
	fmt.Print(list)
	x, y := findWrongPosition(list)
	swap(root, x, y, 2)
}

func inOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, list)
	*list = append(*list, root.Val)
	inOrder(root.Right, list)
}

func findWrongPosition(list []int) (int, int) {
	x, y := -1, -1
	count := 1
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			if count == 1 {
				x = i
				y = i + 1
				count++
			} else {
				y = i + 1
			}
		}
	}
	return list[x], list[y]
}

func swap(root *TreeNode, x, y, count int) {
	if root == nil || count == 0 {
		return
	}
	if root.Val == x {
		root.Val = y
		count--
	} else if root.Val == y {
		root.Val = x
		count--
	}
	swap(root.Left, x, y, count)
	swap(root.Right, x, y, count)
}
