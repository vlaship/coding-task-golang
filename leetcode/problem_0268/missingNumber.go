package problem_0268

func missingNumber(nums []int) int {
	var r = make([]int, len(nums)+1)

	for _, value := range nums {
		r[value] = 1
	}

	for i, v := range r {
		if v == 0 {
			return i
		}
	}

	return -1
}
