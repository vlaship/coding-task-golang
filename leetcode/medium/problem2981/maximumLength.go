package problem2981

func maximumLength(s string) int {
	maxLength := -1
	n := len(s)

	// Iterate over each character in the string
	for i := 0; i < n; {
		char := s[i]
		length := 0

		// Count the length of the special substring
		for i < n && s[i] == char {
			length++
			i++
		}

		// Check all possible lengths that occur at least thrice
		for l := 1; l <= length; l++ {
			if countOccurrences(s, char, l) >= 3 {
				maxLength = max(maxLength, l)
			}
		}
	}

	return maxLength
}

func countOccurrences(s string, char byte, length int) int {
	count := 0
	n := len(s)

	for i := 0; i <= n-length; i++ {
		if isSpecial(s[i:i+length], char) {
			count++
		}
	}

	return count
}

func isSpecial(subs string, char byte) bool {
	for i := 0; i < len(subs); i++ {
		if subs[i] != char {
			return false
		}
	}
	return true
}
