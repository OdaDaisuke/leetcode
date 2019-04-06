package algorithms

func generate(numRows int) [][]int {
	var res [][]int

	for row := 1; row <= numRows; row++ {
		var columns []int
		for column := 1; column <= row; column++ {
			if row > 1 && len(res[row - 2]) == 1 {
				columns = append(columns, 1)
				continue
			}
			// 端っこのカラムなら無条件で1
			if column == 1 || column == row {
				columns = append(columns, 1)
				continue
			}
			prevF := res[row - 2][column - 1]
			prevS := res[row - 2][column - 2]
			columns = append(columns, prevF + prevS)
		}

		res = append(res, columns)
	}

	return res
}

// https://leetcode.com/submissions/detail/220407395/
