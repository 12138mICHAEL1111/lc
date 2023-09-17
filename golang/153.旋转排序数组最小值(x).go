package main

func findMin(nums []int) int {
	left,right := 0, len(nums)-1

	if len(nums) == 1 {
		return nums[0]
	}

	for left < right {
		if right-left == 1 { // l,r相邻，最小值存在于l，r之间，例如 0，1
			return min(nums[right],nums[left])
		}

		mid := left + (right-left) / 2

		if nums[left] < nums[mid] && nums[mid] < nums[right] {
			return nums[left]
		}

		if nums[mid] > nums [left] { // mid比左边大，最小值就在右侧 ，4 5 6 7 8 9 10 0 1 2 3
			left = mid
		}else  {  // 不然最小值就在左侧 ，7，0，1，2，3，4，5，6 （规律，死记硬背）
			right = mid
		}
	}

	return -1
}