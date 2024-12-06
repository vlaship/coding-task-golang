package problem2554

func maxCount(banned []int, n int, maxSum int) int {
	set := make(map[int]struct{})
	for _, v := range banned {
		set[v] = struct{}{}
	}

	sum, count := 0, 0
	for i := 1; i <= n; i++ {
		if _, ok := set[i]; ok {
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
