package main

type BSTIterator struct {
	index int
	inOrderList []int
}


func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		index:0,
		inOrderList: []int{},
	}
	var inOrder func(*TreeNode)
	inOrder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inOrder(tn.Left)
		iterator.inOrderList = append(iterator.inOrderList,tn.Val)
		inOrder(tn.Right)
	}
	inOrder(root)
	return iterator
}


func (this *BSTIterator) Next() int {
	r := this.inOrderList[this.index]
	this.index ++
	return r
}


func (this *BSTIterator) HasNext() bool {
	if this.index >= len(this.inOrderList){
		return false
	}
	return true
}

