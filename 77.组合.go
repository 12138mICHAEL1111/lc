package main

func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}
	var dfs func(int)
	dfs = func(startIndex int) {
		if len(path) == k{
			result = append(result, append([]int(nil),path...))
			return
		}

		for i := startIndex; i<=n ; i++{
			path = append(path, i)
			dfs(i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return result
}