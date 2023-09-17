package main

func findMin2(nums []int) int {
	l := 0 
	r := len(nums) - 1 
	for l < r{
		mid := l + (r-l) / 2 
		if nums[mid] > nums[r]{
			l = mid + 1 
		}else if (nums[mid] < nums[r]){
			r = mid
		}else {
			 r--
		}
	}
	return nums[l]
}