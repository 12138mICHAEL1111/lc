package main

import "strings"

func longestCommonPrefix(strs []string) string {
	firstStr := strs[0] //先选一个字符串作为基准
	index := 1
	prefix := ""
	for index <= len(firstStr) { // 截取这个字符串的前缀
		prefix = firstStr[0:index]
		for _, v := range strs { //再遍历数组
			if !strings.HasPrefix(v, prefix) { //如果数组里有一个字符串没有前缀
				return firstStr[0 : index-1] //返回
			}
		}
		index++
	}
	return prefix
}
