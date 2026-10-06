
func isAnagram(s string, t string) bool {
	m := make(map[byte]int)


	if len(s) != len(t) {
		return false
	}
	
	for i := 0; i < len(s); i++{
		m[s[i]]++
		m[t[i]]--
	}

	for _,v := range m {
		if  v != 0{
			return false
		} 
	}

	return true

}


