package main
func islandPerimeter(grid [][]int) int {
	var dfs func(int,int) int
	dfs = func(i, j int) int {
		// 超出边界
		if (i<0 || i >=len(grid) || j <0 || j>=len(grid[0])){
			return 1
		}

		// 海洋
		if grid[i][j] == 0 {
			return 1
		}

		//已经遍历过，或者是陆地
		if grid[i][j] != 1 {
			return 0
		}

		grid [i][j] = 2

		return dfs(i+1,j) + dfs(i-1,j) + dfs(i,j-1) +dfs(i,j+1)
	}

	for i:=0;i<len(grid);i++{
		for j:=0; j<len(grid[0]); j++ {
			if ( grid[i][j] == 1){
				return dfs(i,j)
			}
		}
	}
	return 0
}