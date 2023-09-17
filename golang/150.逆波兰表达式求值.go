package main

import "strconv"

//遍历字符串，如果是数字就放入栈中，是操作符就弹出两个数字，计算时，第二个在前，再把结果放入栈中
func evalRPN(tokens []string) int {
	stack := []int{}
	for _,v := range tokens{
		switch v{
		case "+" :
			a,b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a,b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a,b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a,b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			i,_:=strconv.Atoi(v)
			stack = append(stack,i)
		}
	}
	return stack[0]
}