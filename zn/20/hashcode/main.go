package main

import (
	"fmt"
)

func HashCode(dec string) string {
	string_size := len(dec)
	hashed_string := ""

	for _, char := range dec {
		hash := (int(char) + string_size) % 127
		if hash < 32 || hash == 127 {
			hash += 33
		}
		hashed_string += string(rune(hash))
	}
	return hashed_string
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}