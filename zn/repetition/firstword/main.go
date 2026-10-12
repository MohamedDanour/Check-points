package main

import "fmt"

func FirstWord(s string) string {
	
	// new_string can only stop being empty if something gets added, but nothing can be added until it stops being empty.
	word := ""
	for _, char := range s {
		if char == ' ' {
			if word != "" {
				break
			}
			continue
		}
		word += string(char)
	}
	return word + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}