package problem_2554

func maxCount(banned []int, n int, maxSum int) int {
	b := make([]bool, n+1)
	for _, v := range banned {
		if v <= n {
			b[v] = true
		}
	}

	sum, count := 0, 0
	for i := 1; i <= n; i++ {
		if b[i] {
			continue
		}
		if sum+i > maxSum {
			return count
		}
		sum += i
		count++
	}

	return count
}
