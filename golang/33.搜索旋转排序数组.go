package main

// mid的左侧或右侧一定是有序数组
// 如果target在有序数组中，更新left或right包含有序数组
// 如果在另一侧，更新left或right包含另一侧
func search(nums []int, target int) int {
	left :=0 
	right := len(nums) - 1 
	mid := 0
	for left <= right {
		mid = left + (right - left) / 2 
		if nums[mid] == target {
			return mid 
		}

		if nums[left] <= nums[mid] { // 左边为有序数组
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			}else{
				left = mid + 1
			}
		}else{ // 右边为有序数组
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}
	return -1
}