package main

import "sort"

func permuteUnique(nums []int) [][]int {
	// 排列
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	visited := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			result = append(result, append([]int(nil), path...))
			return
		}

		for k, v := range nums {
			if visited[k] || (k > 0 && v == nums[k-1] && !visited[k-1]) {
				continue
			}
			path = append(path, v)
			visited[k] = true
			dfs()
			visited[k] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return result
}