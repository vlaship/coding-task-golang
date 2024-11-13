package problem0001

func twoSum(input []int, n int) []int {
	m := make(map[int]int)
	r := make([]int, 0)

	for i, e := range input {
		k := n - e
		v, exists := m[k]
		if exists {
			r = append(r, v)
			r = append(r, i)
			return r
		}

		m[e] = i
	}
	return r
}
