package medium

func convertNaive(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	matrix := make([][]string, numRows)
	row, direction := 0, 1

	for _, char := range s {
		matrix[row] = append(matrix[row], string(char))
		if row == numRows-1 {
			direction = -1
		}
		if row == 0 {
			direction = 1
		}
		row += direction
	}

	var result string
	for _, sub := range matrix {
		for _, char := range sub {
			result += char
		}
	}

	return result
}
