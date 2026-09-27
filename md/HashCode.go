package train

func HashCode(s string) string {
	result := []byte(s)

	
	for i := 0; i < len(s); i++ {

		newChar := (int(s[i]) + len(s)) % 127
		if newChar < 32 || newChar == 127 {
			newChar += 33

		}

		result[i] = byte(newChar)

	}
	return string(result)
}

// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.
