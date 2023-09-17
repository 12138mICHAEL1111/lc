package main

func search2(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] { // 如果左指针等于mid，左指针+1，（背），如果不写这行代码，10110左边不是有序数组
			left++
			continue
		}

		if nums[left] <= nums[mid] { // 左边为有序数组
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边为有序数组
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
