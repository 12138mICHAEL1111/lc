package main

func searchMatrix(matrix [][]int, target int) bool {
	rowLen :=  len(matrix)
	columnLen := len(matrix[0])
	top := 0 
	bottom := rowLen - 1 
	rowIndex := 0 
	if target < matrix[0][0] || target > matrix[rowLen-1][columnLen-1] {
		return false
	}
	for top <= bottom {
		rowIndex = top + (bottom-top)/2
		if target >= matrix[rowIndex][0] && target <= matrix[rowIndex][columnLen-1]{
			break
		}

		if target < matrix[rowIndex][0]  {
			bottom = rowIndex - 1
		}

		if target > matrix[rowIndex][columnLen-1]{
			top = rowIndex + 1
		}
	}

	left:=0
	right:= columnLen-1
	columnIndex := 0
	for left <= right {
		columnIndex = left + (right-left) / 2
		 if target == matrix[rowIndex][columnIndex] {
			return true
		 }

		 if target < matrix[rowIndex][columnIndex]{
			right = columnIndex - 1
		 }

		 if target > matrix[rowIndex][columnIndex]{
			left = columnIndex + 1
		 }
	}

	return false
}

// var searchMatrix = function(matrix, target) {
//     const m = matrix.length, n = matrix[0].length;
//     let low = 0, high = m * n - 1;
//     while (low <= high) {
//         const mid = Math.floor((high - low) / 2) + low; // ******映射到原数组中
//         const x = matrix[Math.floor(mid / n)][mid % n];
//         if (x < target) {
//             low = mid + 1;
//         } else if (x > target) {
//             high = mid - 1;
//         } else {
//             return true;
//         }
//     }
//     return false;
// };
