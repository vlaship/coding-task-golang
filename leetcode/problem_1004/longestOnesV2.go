package problem_1004

func longestOnesV2(nums []int, k int) int {
	res := 0
	zeros := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			for zeros >= k {
				if nums[left] == 0 {
					zeros--
				}
				left++
			}
			zeros++
		}
		res = maxV2(res, right-left+1)
	}
	return res
}

func maxV2(max, new int) int {
	if new > max {
		return new
	}
	return max
}
