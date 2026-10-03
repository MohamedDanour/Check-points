package train

func WeAreUnique(str1, str2 string) int {
	var counter int

	if str1 == "" && str2 == "" {
		return -1
	}

	for i := 0; i <= len(str1)-1; i++ {
		if str1[i] != str2[i] {
			counter += 2
		} else {
			continue
		}
	}
	return counter
}
