package main

import (
	"fmt"
	"strings"
	"unicode"
)

func calculate(s string) int {
	fmt.Print(convertToRPN(s))
	// return evalRPN(convertToRPN(s))
	return 0
}

// 转换为逆波兰式
// 操作符栈，和输出列表
// 左括号，压入栈中
// 右括号，弹出栈里的操作符，直至遇到左括号
// 操作符， 弹出栈里的操作符，添加到输出列表，直到栈为空/优先级较低的操作符/左括号，然后压入当前操作符到栈中
// 数字，直接添加到输出列表
// 遍历完毕后，弹出弹出栈里的操作符到输出列表
func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

// * / 优先you级比+ -高
func priority(s string) int {
	if s == "+" || s == "-" {
		return 1
	} else if s == "*" || s == "/"  {
		return 2
	}else{
		return 0 // 括号
	}
}

func convertToRPN(s string) []string {
	s = strings.ReplaceAll(s, " ", "") // 移除所有空格
	result := []string{}
	stack := []string{}
	sList := []rune(s)
	for k:= 0 ; k<len(sList); k++{
		str := string(sList[k])

		if unicode.IsDigit(sList[k]){
			for k < len(sList) - 1 && unicode.IsDigit(sList[k+1]) { // 连续数字
				str = str + string(sList[k+1])
				k++
			}
			result = append(result, str)
		}
		if isOperator(str) {
			if str == "-" && (k == 0 || string(sList[k-1]) == "(" || isOperator(string(sList[k-1]))) { // 如果前一个字符是左括号或者前一个字符是操作符,就是负数，例如1- -1
				result = append(result,"0") // 添加0
			}else{
				for len(stack) > 0 && priority(stack[len(stack)-1]) >= priority(str) { // 遇到优先级高的就停止
					result = append(result, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
			}
			stack = append(stack, str)
		}
		if str == "("{
			stack = append(stack, "(")
		}
		if str == ")" {
			for stack[len(stack)-1] != "("{
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // 弹出左括号
		}
	}

	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result
}