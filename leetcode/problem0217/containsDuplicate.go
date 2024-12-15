package problem0217

func containsDuplicate(nums []int) bool {
	var m = make(map[int]struct{})
	empty := struct{}{}

	for _, num := range nums {
		_, exists := m[num]
		if exists {
			return true
		}
		m[num] = empty
	}
	return false
}
