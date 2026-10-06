func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	out := make([]int, 2, 2)

	for i,num := range nums {
		rest := target - num
		v, ok := m[rest] 
		if ok {
			out[0] = v
			out[1] = i
			return out
		} else {
			m[num] = i
		}
	}
	return out
}
