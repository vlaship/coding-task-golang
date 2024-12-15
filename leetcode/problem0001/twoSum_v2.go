package problem0001

/*
Find all pairs in an array of int type, whose sum is equal to the specified number n

Example
Input: arr: [1,3,3,4,5,2], n: 6
Output: [1,5], [3,3], [2,4]
*/
func twoSumV2(input []int, n int) [][]int {
	m := make(map[int]int)
	r := make([][]int, 0)

	for _, e := range input {
		k := n - e
		v, exists := m[k]
		if exists {
			r = append(r, []int{k, v})
			continue
		}

		m[e] = k
	}
	return r
}
