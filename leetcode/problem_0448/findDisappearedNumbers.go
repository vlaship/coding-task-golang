package problem_0448

func findDisappearedNumbers(nums []int) []int {
	var r = make([]int, len(nums))
	var res = make([]int, 0)

	for _, value := range nums {
		r[value-1] = 1
	}

	for i, v := range r {
		if v == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
