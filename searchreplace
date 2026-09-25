package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	chars := []rune(os.Args[1])
	found := false

	for index, char := range chars {
		if string(char) == os.Args[2] {
			chars[index] = []rune(os.Args[3])[0]
			found = true
		}
	}

	if !found {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println(string(chars))
	}
}
