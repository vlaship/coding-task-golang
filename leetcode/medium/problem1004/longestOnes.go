package problem1004

/*
1004. Max Consecutive Ones III

https://leetcode.com/problems/max-consecutive-ones-iii/

Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Example 1:
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Constraints:
1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length
*/
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
		m = max(m, left, idx)
	}
	return m
}

func max(max, left, idx int) int {
	d := idx - left
	if d > max {
		return d
	}
	return max
}
