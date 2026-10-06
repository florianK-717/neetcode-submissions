func hasDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		_, ok := m[v] 
		if ok {
			return ok
		} else {
			m[v] = !ok
		}
	}
	return false 
}
