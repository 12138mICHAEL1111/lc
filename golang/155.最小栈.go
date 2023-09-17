package main

type MinStack struct {
	stack []int
	mStack []int
}


func ConstructorStack() MinStack {
	s := MinStack{
		stack: []int{}, // 输出
		mStack: []int{}, // 用于存放最小值
	}
	return s
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	if len(this.mStack) == 0 || val <= this.mStack[len(this.mStack)-1]{ // 最小值放进mStack里
		this.mStack = append(this.mStack, val)
	}
}


func (this *MinStack) Pop()  {
	if this.stack[len(this.stack)-1] == this.mStack[len(this.mStack)-1] { // 弹出的是最小栈顶端元素
		this.mStack = this.mStack[:len(this.mStack)-1]
	} 
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.mStack[len(this.mStack) - 1]
}