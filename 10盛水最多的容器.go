package main

import "math"

// 总是移动短的一侧的木板， 因为是按短的那一侧来计算高度，
// 如果移动长的，那么结果要么不变，要么变小，因为每次移动底永远是-1
// 例如 2 ，5 ，6
// 移动2 那么长就变成5了 。移动6， 长还是2，而底会减少，那么面积一定少
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	area := 0.0
	for i < j {
		if height[i] < height[j] {
			area = math.Max(area, float64((j-i)*height[i]))
			i++
		} else {
			area = math.Max(area, float64((j-i)*height[j]))
			j--
		}
	}
	return int(area)
}
