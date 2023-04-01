package medium

/*
https://leetcode.com/problems/triangle/

Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally,
if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
2
3 4
6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Example 2:
Input: triangle = [[-10]]
Output: -10

Constraints:
1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?
*/
func minimumTotal(triangle [][]int) int {
	var cached = [][]int{}

	// 0 row
	cached = append(cached, make([]int, 1))
	cached[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		cached = append(cached, make([]int, len(triangle[i])))

		cached[i][0] = cached[i-1][0] + triangle[i][0]
		cached[i][len(triangle[i])-1] = cached[i-1][len(cached[i-1])-1] + triangle[i][len(triangle[i])-1]

		for j := len(triangle[i]) - 2; j > 0; j-- {
			cj1 := cached[i-1][j-1]
			cj2 := cached[i-1][j]

			var cj int
			if cj1 > cj2 {
				cj = cj2
			} else {
				cj = cj1
			}

			cached[i][j] = cj + triangle[i][j]
		}
	}

	last := len(cached) - 1
	res := cached[last][0]
	for i := 1; i < len(cached[last]); i++ {
		if cached[last][i] < res {
			res = cached[last][i]
		}
	}

	return res
}
