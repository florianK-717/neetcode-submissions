
func isAnagram(s string, t string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}
	
	for i := 0; i < len(s); i++{
		c := s[i]
		adaptMap(m1, c)
		b := t[i]
		adaptMap(m2, b)
	}

	for k,v1 := range m1 {
		if  v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		} 
	}

	return true

}

func adaptMap(m map[byte]int, c byte) {
	v, ok := m[c]
	if ok {
		m[c] = v+1
	} else {
		m[c] = 1
	} 
}
