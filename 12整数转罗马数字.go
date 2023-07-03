package main

import (
	"strings"
)

//先预设好每个int对应的字符，再遍历valueSymbols，
//如果数字比当前便利的int大或等于，就一直减去int，直到比他小
//把对应的字符储存起来
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var sb strings.Builder
	for _, v := range valueSymbols {
		for num >= v.value {
			sb.WriteString(v.symbol)
			num = num - v.value
		}
		if num == 0 {
			break
		}
	}
	return sb.String()
}
