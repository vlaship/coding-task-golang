package problem1004

func longestOnes(nums []int, k int) int {
	m := 0
	zeros := k
	left := 0
	idx := 0
	for idx < len(nums) {
		if nums[idx] == 0 {
			if zeros > 0 {
				zeros--
			} else {
				if nums[left] == 0 {
					zeros++
				}
				left++
				continue
			}
		}
		idx++
		m = longest(m, left, idx)
	}
	return m
}

func longest(max, left, idx int) int {
	d := idx - left
	if d > max {
		return d
	}
	return max
}
