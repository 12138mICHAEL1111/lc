package main

func searchMatrix2(matrix [][]int, target int) bool {
    var dfs func(int,int) bool 
	rl := len(matrix)
	cl := len(matrix[0])
	dfs = func(row, col int) bool {
		if row == rl || col == -1 {
			return false
		}

		if target ==  matrix[row][col] {
			return true
		}else if target < matrix[row][col] {
			return dfs(row,col - 1)
		}else{
			return dfs(row+1 ,col)
		}
	}
	return dfs(0,cl-1)
}

