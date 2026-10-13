
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var buffer bytes.Buffer

	for _, str := range strs {
		buffer.WriteString(strconv.Itoa(len(str)))
		buffer.WriteByte('#')
		buffer.WriteString(str)
	}

	return buffer.String()
}

func (s *Solution) Decode(encoded string) []string {
	var decodedStrs []string

	for len(encoded) > 0 {
		separator := strings.IndexByte(encoded, '#')
		if separator == -1 {
			return decodedStrs
		}

		length, err := strconv.Atoi(encoded[:separator])
		if err != nil {
			return decodedStrs
		}

		start := separator + 1
		end := start + length

		if end > len(encoded) {
			return decodedStrs
		}

		decodedStrs = append(decodedStrs, encoded[start:end])
		encoded = encoded[end:]
	}

	return decodedStrs
}
