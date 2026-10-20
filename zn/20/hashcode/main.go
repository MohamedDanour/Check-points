package main

import (
	"fmt"
)

func HashCode(dec string) string {
	string_size := rune(len(dec))
	new_string := ""

	for _, char := range dec {
		char = (char + string_size) % 127
		if char < 32 { // unprintable
			char += 33
		}
		new_string += string(char)
	}
	return new_string
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}