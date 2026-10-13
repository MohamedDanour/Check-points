package main 

import (
	"os"
	"fmt"
)

func searchreplace(word string, letter rune, rep_letter rune) string {

	exists := false 
	for _, char := range word {
		if char == letter {
			exists = true
		} 
	}

	if !exists {
		return word 
	}

	new_word := ""
	for _, char := range word {
		if char == letter {
			new_word += string(rep_letter)
		} else {
			new_word += string(char)
		}
	}
	return new_word
}

func main() {
	if len(os.Args) != 4 { // 3 args + progam name
		return 
	}

	// argument
	word := os.Args[1]
	letter := []rune(os.Args[2])[0]
	rep_letter := []rune(os.Args[3])[0]

	fmt.Println(searchreplace(word, letter, rep_letter))
}