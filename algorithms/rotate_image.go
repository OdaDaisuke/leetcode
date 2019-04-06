package algorithms

func rotate(matrix [][]int) {
	sideLen := len(matrix)
	var newMatrix = make([][]int, sideLen)
	for i, _ := range newMatrix {
		newMatrix[i] = make([]int, sideLen)
	}

	for rowIdx, row := range matrix {
		for colIdx, _ := range row {
			newMatrix[colIdx][sideLen - rowIdx - 1] = matrix[rowIdx][colIdx]
		}
	}

	for rowIdx, row := range newMatrix {
		matrix[rowIdx] = row
	}

}

// https://leetcode.com/submissions/detail/220418559/