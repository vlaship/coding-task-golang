package problem0001

func problem001(input []int) []int {
	var sorted = make([]int, len(input))

	r := 0
	l := len(input) - 1

	for i := l; i >= 0; i-- {
		left := input[l] * input[l]
		right := input[r] * input[r]

		if left > right {
			sorted[i] = left
			l--
		} else {
			sorted[i] = right
			r++
		}
	}

	return sorted
}
