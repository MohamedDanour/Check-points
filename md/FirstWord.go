package train

func FirstWord(s string) string {

	newStr := ""

	for index, char := range s {
		if char != ' ' {
			newStr += string(s[index])
		} else {
			break
		}
	}
	return newStr + "\n"
}
