package problem2825

func canMakeSubsequence(str1 string, str2 string) bool {
	l, r := 0, 0

	for r < len(str2) && l < len(str1) {
		u := str1[l] + 1
		if u == 123 {
			u = 97
		}
		if str1[l] == str2[r] || u == str2[r] {
			r++
		}
		l++
	}

	return r == len(str2)
}
