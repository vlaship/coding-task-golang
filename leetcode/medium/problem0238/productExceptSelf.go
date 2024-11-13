package problem0238

func productExceptSelf(nums []int) []int {
	n := len(nums)

	answer := make([]int, n)

	leftProduct := 1
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}
