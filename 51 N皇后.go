package main

// 基本思路: 放一个皇后，把对应的列，两条对角线的值都设为true，放下一个皇后的时候检查符不符合

func solveNQueens(n int) [][]string {
	queens := make([]int,n) //第n项代表第n行queen的位置
	diagonal1Map := make(map[int]bool) //从左到右的斜线
	diagonal2Map := make(map[int]bool)	//右到左
	column := make(map[int]bool)
	result := [][]string{}
	var dfs func(int)
	dfs = func(row int) {
		if row == n{
			oneSolution := generateSolution(&queens,n)
			result = append(result, oneSolution)
			return
		}

		for i := 0; i<n ; i++{ //遍历列
			if column[i]  { //当前列已经有皇后
				continue
			}

			diagonal1Pos := i +row // 一条斜线上的这个值都相等
			if diagonal1Map[diagonal1Pos]{
				continue
			}

			diagonal2Pos := i - row 
			if diagonal2Map[diagonal2Pos]{
				continue
			}

			column[i] = true
			diagonal1Map[diagonal1Pos] = true
			diagonal2Map[diagonal2Pos] = true
			queens[row] = i
			dfs(row + 1)  // 下一行
			queens[row] = -1
			column[i] = false
			diagonal1Map[diagonal1Pos] = false
			diagonal2Map[diagonal2Pos] = false
		}	

	}
	dfs(0)
	return result
}

func generateSolution(queens *[]int,n int)[]string{
	oneSolution := []string{}
	for i := 0; i < n; i++{ // 行
		oneRow := ""
		for j := 0; j < n ; j++{ //列
			if j == (*queens)[i] {
				oneRow = oneRow + "Q"
			}else{
				oneRow = oneRow + "."
			}
		}
		oneSolution = append(oneSolution, oneRow)
	}
	return oneSolution
}