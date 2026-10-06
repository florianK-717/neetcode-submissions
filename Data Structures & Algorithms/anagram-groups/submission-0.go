func groupAnagrams(strs []string) [][]string {
	out := make([][]string, 0, len(strs))
	ctr := make(map[[26]int][]string)

	for _,v := range strs {
		var charCounter [26]int
		for _,c := range v {
			charCounter[c-'a']++ 
		}
		ctr[charCounter] = append(ctr[charCounter], v)
	}
	for _,v := range ctr {
		out = append(out, v)
	}
	return out
}
