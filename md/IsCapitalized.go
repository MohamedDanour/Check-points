package train

func IsCapitalized(s string) bool {
	for index, _ := range s {
		if (s[0] >= 'A' || s[0] <= 'Z') && (s[index] >= ' ' && s[index+1] >= 'A' && s[0] <= 'Z') {
			return true
		} else if s[index] == ' ' && s[index+1] <= 'a' || s[index+1] >= 'z' {
			return false
		}

	}
	return false
}
