package problem2109

import (
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var res strings.Builder
	prev := 0
	for _, idx := range spaces {
		res.WriteString(s[prev:idx])
		res.WriteByte(' ')
		prev = idx
	}
	res.WriteString(s[prev:])
	return res.String()
}
