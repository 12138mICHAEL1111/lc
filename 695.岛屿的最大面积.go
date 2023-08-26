package main
func maxAreaOfIsland(grid [][]int) int {
	var dfs func(int,int) int
	dfs = func(i, j int) int {
		if i<0 || i>= len(grid) || j<0 || j>= len(grid[0]){
			return 0
		}

		if grid[i][j] != 1{
			return 0
		}
	
		grid[i][j] = 2

		return 1 + dfs (i-1,j) + dfs (i+1,j) + dfs (i,j-1) + dfs (i,j+1)
	}

	maxArea := 0 

	for i:=0;i<len(grid);i++ {
		for j := 0;j<len(grid[0]) ; j ++{
			if grid[i][j] == 1{
				area := dfs(i,j)
				if area > maxArea{
					maxArea = area
				}
			}
		}
	}

	return maxArea
}