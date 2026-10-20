package main

import (
    "fmt"
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
    if s == "" {
		return "\n"
	}

	first_letter := false
	first_word := ""

	for index, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			first_letter = true 
		}

		if first_letter {
			first_word += string(char)
			if index == len(s)-1 || s[index+1] == ' '  {
				break
			}	
		}
	}
	return first_word + "\n"

	

}