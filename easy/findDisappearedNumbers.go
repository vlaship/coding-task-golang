package easy

/*
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:
Input: nums = [1,1]
Output: [2]
*/
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
