package problem0005

/*
5. Longest Palindromic Substring

https://leetcode.com/problems/longest-palindromic-substring/

Given a string s, return the longest palindromic substring  in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters.
*/
func longestPalindrome(s string) string {
	var max string

	for i := range s {
		p1 := getPal(s, i, i)
		p2 := getPal(s, i, i+1)
		if len(p1) > len(max) {
			max = p1
		}
		if len(p2) > len(max) {
			max = p2
		}
	}

	return max
}

func getPal(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
