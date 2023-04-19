package easy

/*
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/
func isValidParentheses(s string) bool {
	if len(s)%2 > 0 {
		return false
	}

	p := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []rune{}

	for _, el := range s {
		l := len(stack)
		_, ok := p[el]
		if !ok {
			stack = append(stack, el)
		} else if l == 0 || stack[l-1] != p[el] {
			return false
		} else {
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
