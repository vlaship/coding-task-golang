package problem_0020

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
