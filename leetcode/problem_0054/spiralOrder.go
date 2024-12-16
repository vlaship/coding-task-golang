package problem_0054

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	result := make([]int, 0)
	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1

	for left <= right && top <= bottom {
		// Traverse right
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// Traverse down
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// Traverse left
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		// Traverse up
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
