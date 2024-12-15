package problem0136

func singleNumber(nums []int) int {
	// only even solution
	var ans int //since XOR with 0 returns same number

	for _, num := range nums {
		ans ^= num // ans = (ans) XOR (array element)
	}

	return ans
}
