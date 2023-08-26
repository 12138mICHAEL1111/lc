package main

import (
	"math"
)

func trap(height []int) int {
	i := 0
	j := len(height) - 1
	volumn := 0
	leftMax := height[0]
	rightMax := height[len(height)-1]
	for i < j {
		if leftMax < rightMax { // 小于或小于等于都行,如果leftMax小于rightMax，那么i-j（不包括i，包括j）这个区间一定存在一个高度比leftMax大或者相等，因为包括j,至少会有j那么高
			volumn = volumn + (leftMax - height[i]) //所以算当前格子的体积直接用leftMax减当前高度就行，一个格子的体积永远按照左右两侧短的那一边来计算
			i++                                     //右侧不可能都比左侧矮，因为在leftMax < rightMax这个条件下,就算未知区间全部比leftmax小，那么至少还有rightMax的高度
			leftMax = int(math.Max(float64(leftMax), float64(height[i])))
		} else {
			volumn = volumn + (rightMax - height[j])
			j--
			rightMax = int(math.Max(float64(rightMax), float64(height[j])))
		}
	}
	return volumn
}
