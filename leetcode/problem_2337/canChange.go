package problem_2337

func canChange(start string, target string) bool {
	strt, trgt := parse(start), parse(target)

	if len(strt) != len(trgt) {
		return false
	}

	for i, s := range strt {
		t := trgt[i]
		if s.direction != t.direction {
			return false
		}
		if s.direction == "L" && s.index < t.index {
			return false
		}
		if s.direction == "R" && s.index > t.index {
			return false
		}
	}

	return true
}

type move struct {
	direction string
	index     int
}

func parse(str string) []move {
	var res []move

	for i, c := range str {
		if c == 'L' {
			res = append(res, move{"L", i})
		} else if c == 'R' {
			res = append(res, move{"R", i})
		}
	}

	return res
}
