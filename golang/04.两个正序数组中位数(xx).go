package main

import "fmt"

// x 一般难，第一遍理解会有苦难，但有套路
// xx 较难， 无套路，但能理解加背
// xxx 非常难， 很难理解，只能背

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	if totalLength % 2 == 0 {
		fisrt := float64(find(nums1, nums2, totalLength/2))
		second := float64(find(nums1, nums2, (totalLength/2)+1))
		fmt.Println(nums1,nums2)
		return (fisrt + second) / 2
	} else {
		return float64(find(nums1, nums2, (totalLength/2)+1))
	}
}

func find(nums1, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 == 0 {
		return nums2[k-1]
	}

	if len2 == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	position1 := min(k/2, len1) - 1
	position2 := min(k/2, len2) - 1
	num1 := nums1[position1]
	num2 := nums2[position2]

	if num1 < num2 {
		if k/2 >= len1 {
			nums1 = []int{}
		} else {
			nums1 = nums1[k/2 : len1]
		}
		return find(nums1, nums2, k-min(k/2, len1))
	} else {
		if k/2 >= len2 {
			nums2 = []int{}
		} else {
			nums2 = nums2[k/2 : len2]
		}
		return find(nums1,nums2, k-min(k/2, len2))
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}