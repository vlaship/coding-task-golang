package problem_0205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mS := make(map[byte]byte)
	mT := make(map[byte]byte)
	for i := range s {
		bS := s[i]
		bT := t[i]

		if b, ok := mS[bS]; !ok {
			if _, ok := mT[bT]; ok {
				return false
			}
			mT[bT] = bS
			mS[bS] = bT
		} else if b != bT {
			return false
		}
	}

	return true
}
