package problem_0136

func singleNumberNonOnlyEven(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		_, ex := m[v]
		m[v] = ex
	}

	for k, v := range m {
		if !v {
			return k
		}
	}

	return -1
}
