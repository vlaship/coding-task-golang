package problem0120

import "math"

/*
120. Triangle

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
func minimumTotalV2(triangle [][]int) int {
	n := len(triangle)

	var cached = make([][]int, 2)
	for i := 0; i < 2; i++ {
		cached[i] = make([]int, n)
		for j := range cached[i] {
			cached[i][j] = math.MaxInt
		}
	}

	cached[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		for j, num := range triangle[i] {
			cached[i%2][j] = minInt(cached[(i-1)%2][j], cached[(i-1)%2][maxInt(j-1, 0)]) + num
		}
	}
	res := math.MaxInt
	for _, num := range cached[(n-1)%2] {
		res = minInt(res, num)
	}
	return res
}
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
