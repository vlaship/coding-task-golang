package problem0005

func longestPalindrome(s string) string {
	var m string

	for i := range s {
		p1 := getPal(s, i, i)
		p2 := getPal(s, i, i+1)
		if len(p1) > len(m) {
			m = p1
		}
		if len(p2) > len(m) {
			m = p2
		}
	}

	return m
}

func getPal(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
