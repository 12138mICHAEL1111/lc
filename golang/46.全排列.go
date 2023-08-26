package main

func permute(nums []int) [][]int {
	result := [][]int{}
	visited := make(map[int]bool)
	path := []int{}
	var dfs func()
	dfs = func() {
		if len(path) == len(nums){
			result = append(result, append([]int(nil),path...))
			return
		}

		for _,v := range nums{
			if (visited[v]){
				continue
			}
			visited[v] = true
			path = append(path, v)
			dfs()
			path = path[:len(path)-1]
			visited[v] = false
		}

	}
	dfs()
	return result
}