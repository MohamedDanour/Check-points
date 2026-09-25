package train

func LastWord(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	last := i // now we have the word without any spaces i traced it
	for i >= 0 && s[i] != ' ' {
		i--
	} // this supposed to count the characters and after tracing it is 11
	start := i + 1 // supposed to 0 i is zero
	newStr := ""
	for start <= last {
		newStr += string(s[start])
		start++
	}
	return newStr + "\n"
}
