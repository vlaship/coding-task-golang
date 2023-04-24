package problem0344

/*
344. Reverse String

https://leetcode.com/problems/reverse-string/

Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

Constraints:
1 <= s.length <= 105
s[i] is a printable ascii character.
*/
func reverseString(input string) string {
	in := []rune(input)
	l := len(in)
	var res = []rune{}
	for i := range in {
		res = append(res, in[l-i-1])
	}

	return string(res)
}
