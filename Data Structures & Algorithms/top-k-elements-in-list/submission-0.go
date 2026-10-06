func topKFrequent(nums []int, k int) []int {
	v := make(map[int]int)

	for _,num  := range nums {
		v[num]++
	}

	buckets := make([][]int, len(nums)+1)
    for i,count := range v {
		buckets[count] = append(buckets[count], i)
	}

	result := []int{}
	for count := len(buckets) -1; count >= 0 && len(result) < k;count--{
		for _,num := range buckets[count] {
			result = append(result, num)

			if len(result) == k {
				break
			}
		}
	}
	return result
}