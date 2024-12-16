package problem_0006

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 || len(s) < 2 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	row := 0
	dw := false

	for _, c := range s {
		rows[row].WriteRune(c)
		if row == 0 || row == numRows-1 {
			dw = !dw
		}
		if dw {
			row++
		} else {
			row--
		}
	}

	res := strings.Builder{}
	for _, r := range rows {
		res.WriteString(r.String())
	}

	return res.String()
}
