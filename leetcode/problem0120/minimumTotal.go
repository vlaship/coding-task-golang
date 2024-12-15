package problem0120

func minimumTotal(triangle [][]int) int {
	var cached = [][]int{}

	// 0 row
	cached = append(cached, make([]int, 1))
	cached[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		cached = append(cached, make([]int, len(triangle[i])))

		cached[i][0] = cached[i-1][0] + triangle[i][0]
		cached[i][len(triangle[i])-1] = cached[i-1][len(cached[i-1])-1] + triangle[i][len(triangle[i])-1]

		for j := len(triangle[i]) - 2; j > 0; j-- {
			cj1 := cached[i-1][j-1]
			cj2 := cached[i-1][j]

			var cj int
			if cj1 > cj2 {
				cj = cj2
			} else {
				cj = cj1
			}

			cached[i][j] = cj + triangle[i][j]
		}
	}

	last := len(cached) - 1
	res := cached[last][0]
	for i := 1; i < len(cached[last]); i++ {
		if cached[last][i] < res {
			res = cached[last][i]
		}
	}

	return res
}
