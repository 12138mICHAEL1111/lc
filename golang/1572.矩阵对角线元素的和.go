package main

func diagonalSum(mat [][]int) int {
	length := len(mat)
	var rightDiaSum int
	var leftDiaSum int
	initCol := length-1
    for i:=0 ;i < length; i++{
		rightDiaSum = rightDiaSum + mat[i][i]
		leftDiaSum = leftDiaSum + mat[i][initCol]
		initCol --
	}
	sum := leftDiaSum + rightDiaSum
	if length%2 != 0 {
		sum = sum - mat[length/2][length/2]
	}
	return sum
}