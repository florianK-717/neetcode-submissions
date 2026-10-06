func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i,num := range nums {
		rest := target - num
		v, ok := m[rest] 
		if ok {
			return []int{v,i}
		} 
		m[num] = i
	}
	return nil
}
