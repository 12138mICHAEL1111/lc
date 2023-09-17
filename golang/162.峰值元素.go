package main

import "math"

// 二分查找，总往大的那一侧找
// 因为往小的一侧就是下坡，不一定能再上坡
// 往大的就一定能到坡顶
func findPeakElement(nums []int) int {
	var get func(int) int
	get = func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt32
		}
		return nums[i]
	}
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}
		if get(mid+1) > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return 0
}



