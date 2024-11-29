package problem0003

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	start := 0
	curr := 0
	m := 0
	set := make(map[byte]struct{})

	for curr < len(s) {
		l := s[curr]
		_, ok := set[l]
		if ok {
			start++
			curr = start
			c := len(set)
			if c > m {
				m = c
			}
			clear(set)
			continue
		}

		set[l] = struct{}{}
		curr++
	}

	c := len(set)
	if c > m {
		m = c
	}

	return m
}
