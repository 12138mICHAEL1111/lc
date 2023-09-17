package main

import "fmt"

// x 一般难，第一遍理解会有困难，但有套路
// xx 较难， 无套路，但能理解加背
// xxx 非常难， 很难理解，只能背

// 总体思路：
// 假设有13个数，那么第7个最小的数就是中位数，我们就要找到第7个数
// 7/2 = 3 ， 那么第一个数组或第二个数组里的前三个数一定是小于第七个数
// 比较两个数组的第三个数，哪个小，就把前三个数移除，因为这三个数一定小于第七个数
// 用7-3=4，就代表我们要找到第四个最小的数，重复以上步骤直到要找第1小的数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	if totalLength%2 == 0 {
		fisrt := float64(find(nums1, nums2, totalLength/2))
		second := float64(find(nums1, nums2, (totalLength/2)+1)) // 用第七个数和第八个数的平均数
		fmt.Println(nums1, nums2)
		return (fisrt + second) / 2
	} else {
		return float64(find(nums1, nums2, (totalLength/2)+1))
	}
}

func find(nums1, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)

	// 如果一个数组没有元素了，直接返回另一个数组k-1
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
		// 如果k/2超出了数组长度，直接用数组最后一个值比较
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
		return find(nums1, nums2, k-min(k/2, len2))
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}
