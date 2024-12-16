package problem_1346

func checkIfExist(arr []int) bool {

	m := make(map[int]int)

	for i, v := range arr {
		if j, ok := m[v*2]; ok && i != j {
			return true
		}
		if v%2 == 0 {
			if j, ok := m[v/2]; ok && i != j {
				return true
			}
		}
		m[v] = i
	}

	return false
}
