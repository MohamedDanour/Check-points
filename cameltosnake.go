package train

func CamelToSnakeCase(s string) string {
	NewStr := ""

	if (s[len(s)-1] <= 'Z' && s[len(s)-1] >= 'A') || (s[len(s)-1] >= '0' && s[len(s)-1] <= '9') {
		return s
	}
	for index, char := range s {
		if (s[index] >= 'A' && s[index] <= 'Z') && (s[index+1] >= 'A' && s[index+1] <= 'Z') {
			return s
		} else if index > 0 &&
			(s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'a' && s[index-1] <= 'z') {
			NewStr += "_"
		}
		NewStr += string(char)
	}
	return NewStr
}
