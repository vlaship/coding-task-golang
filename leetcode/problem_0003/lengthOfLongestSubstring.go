package problem_0003

func lengthOfLongestSubstring(s string) int {
	start, maxLen := 0, 0
	set := make(map[byte]struct{})

	for curr := 0; curr < len(s); curr++ {
		for _, ok := set[s[curr]]; ok; _, ok = set[s[curr]] {
			delete(set, s[start])
			start++
		}
		set[s[curr]] = struct{}{}
		curLen := curr - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
