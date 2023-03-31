package medium

/*
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
