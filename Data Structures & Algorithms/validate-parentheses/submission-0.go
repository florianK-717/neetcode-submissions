func isValid(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}

	if len(s)%2 != 0 {
		return false
	}

	for _, char := range s {
		switch char {
			case '(', '[', '{': 
			 	stack = append(stack, char)
			default: 
				if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
    
}
