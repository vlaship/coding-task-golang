package problem0120

import "math"

func minimumTotalV2(triangle [][]int) int {
	n := len(triangle)

	var cached = make([][]int, 2)
	for i := 0; i < 2; i++ {
		cached[i] = make([]int, n)
		for j := range cached[i] {
			cached[i][j] = math.MaxInt
		}
	}

	cached[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		for j, num := range triangle[i] {
			cached[i%2][j] = minInt(cached[(i-1)%2][j], cached[(i-1)%2][maxInt(j-1, 0)]) + num
		}
	}
	res := math.MaxInt
	for _, num := range cached[(n-1)%2] {
		res = minInt(res, num)
	}
	return res
}
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
