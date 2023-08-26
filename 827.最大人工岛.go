package main

// https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
func largestIsland(grid [][]int) int {
	var dfs func(int, int, int) int
	dfs = func(i, j, index int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return 0
		}

		if grid[i][j] != 1 {
			return 0
		}

		grid[i][j] = index

		return 1 + dfs(i-1, j, index) + dfs(i+1, j, index) + dfs(i, j-1, index) + dfs(i, j+1, index)
	}
	areaMap := make(map[int]int)
	index := 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j, index)
				areaMap[index] = area
				index++
			}
		}
	}
	maxArea := areaMap[2]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				nearIsland := []int{}
				// 遍历上下左右四个格子，如果是岛屿，就添加到nearIsland
				if i-1 >= 0 {
					if grid[i-1][j] != 0 && !Contains(&nearIsland,grid[i-1][j]){
						nearIsland = append(nearIsland, grid[i-1][j])
					}
				}
				if i+1 < len(grid) {
					if grid[i+1][j] != 0 && !Contains(&nearIsland,grid[i+1][j] ) {
						nearIsland = append(nearIsland, grid[i+1][j])
					}
				}
				if j-1 >= 0 {
					if grid[i][j-1] != 0 && !Contains(&nearIsland,grid[i][j-1]){
						nearIsland = append(nearIsland, grid[i][j-1])
					}
				}
				if j+1 < len(grid) {
					if grid[i][j+1] != 0 && !Contains(&nearIsland,grid[i][j+1]){
						nearIsland = append(nearIsland, grid[i][j+1])
					}
				}
				// fmt.Println(nearIsland)
				area := 1
				for _, v := range nearIsland {
					area = area + areaMap[v]
				}
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	// fmt.Println(grid)
	return maxArea
}

func Contains(slice *[]int, item int) bool {
    for _, value := range *slice {
        if value == item {
            return true
        }
    }
    return false
}

func RemoveDuplicates(input []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, item := range input {
		if _, exists := seen[item]; !exists {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}