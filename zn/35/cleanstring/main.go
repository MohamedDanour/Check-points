package main 

import (
    "fmt"
    "os"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
		fmt.Println() // en ny rad 
		return
	}

	fmt.Print(cleanstr(os.Args[1])) // kör programmet med andra argumentet 
}

func cleanstr(s string) string {
	new_word := ""
	needs_space := false

	for _, char := range s {
	
		if char == ' ' || char == '\t' {
			needs_space = true
		} else { 
			if needs_space && len(new_word) != 0 { // handles start of the string 
				new_word += " "
			}
			needs_space = false // reset on every letter
			new_word += string(char) // we add the space only after checkling if needs_space is true so it wont gets added as if it were a letter 
		}	
	}
	return new_word + "\n"
}


// IMPORTANT: 
// A function gets its input from whoever calls it in code 
// A program gets its input from the command line


// NOTES 
// Should the if needs_space check run on every character, or only inside else?
// Inside else, should it come before or after you add the letter?