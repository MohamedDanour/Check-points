// problem: how do we know w are in the beginning of a word?
// answer: character before current char was a space, or char is the very first char of the string.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}


func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for index, character := range s {
		if (index == 0 || s[index-1] == ' ') && character >= 'a' && character <= 'z' {
			return false
		}
	}
	return true
}
