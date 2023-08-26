package main

func numIslands(grid [][]byte) int {
	var dfs func(int,int)
	rowNo := len(grid)
	columnNo := len(grid[0])
	dfs = func(i,j int){
		if (i < 0 || i >= rowNo || j < 0 || j >=columnNo || grid[i][j] != '1'){
			return
		}

		grid[i][j] = '2'
		
		dfs(i+1,j)
		dfs(i-1,j)
		dfs(i,j+1)
		dfs(i,j-1)
	}

	num := 0

	for  i:=0 ;i<rowNo;i++{
		for j:= 0 ; j< columnNo ; j++ {
			if grid[i][j] == '1'{
				dfs(i,j)
				num ++
			}
		}
	}
	
	return num
}