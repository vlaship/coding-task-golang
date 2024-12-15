package problem0344

func reverseString(input string) string {
	in := []rune(input)
	l := len(in)
	var res = []rune{}
	for i := range in {
		res = append(res, in[l-i-1])
	}

	return string(res)
}
