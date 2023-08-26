package main

import "math"
//和第9题一样
func reverse(x int) int {
	reversedNum := 0
	for (x != 0){
		reversedNum = reversedNum* 10 + x % 10
		x = x/10
	}
	if reversedNum > math.MaxInt32 || reversedNum < math.MinInt32 {
		return 0
	}
	return reversedNum
}