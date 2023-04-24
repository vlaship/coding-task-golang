package problem0001

/*
Given an array of integers sorted in non-decreasing order.
Return an array of the squares of each number, also in sorted non-decreasing order.

Example
Input: [-4,-1,0,3,4,10]
Output: [0,1,9,16,16,100]
*/
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
