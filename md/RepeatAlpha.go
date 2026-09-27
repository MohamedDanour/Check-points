package train

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			pos := (c - 'a') + 1
			for i := 0; i < int(pos); i++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			pos := (c - 'A') + 1
			for i := 0; i < int(pos); i++ {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}
	return result
}