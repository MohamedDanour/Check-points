package main

import (
	"fmt"
	"piscine"
)

func RepeatAlpha(s string) string {
	// alphabetical index 
	word := ""

	for _, char := range s {

		if char >= 'a' && char <= 'z' {
			position := char - 'a' + 1
			for i := 0 ; i < int(position); i++ {
				word += string(char)
			}
			
		} else if char >= 'A' && char <= 'Z' {
			position := char - 'A' + 1
			for i := 0 ; i < int(position); i++ {
				word += string(char)
			}

		} else {
			word += char 
		}
	}
	return word
}


func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}