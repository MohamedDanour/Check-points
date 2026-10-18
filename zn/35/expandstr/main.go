package main 

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 || expandstr(os.Args[1]) == "" { // expandstr(os.Args[1]) == "" checks no words at all, e.g. "     "
		return // display nothing
	} 

	fmt.Println(expandstr(os.Args[1])) // Println adds the '\n' 
}

func expandstr(s string) string {
	new_string := ""
	needs_spaces := false

	for _, char := range s {
		if char >= '!' && char <= '~' { // visible char
			if needs_spaces && len(new_string) != 0 { // skip spaces in beginning (leading spaces)
				new_string += "   " 
			}
			needs_spaces = false
			new_string += string(char)
		} else { 
			needs_spaces = true
		}
	}
	return new_string
}


// NOTES
// A character is visible if its ASCII code is between 33 (!) and 126 (~)
// A character is not visible if its ASCII code is 32 (space) or lower, or 127 (DEL).

/*
Same as cleanstr logic. Only two things differ:
The test is flipped. Before, you asked “is it a space or tab?” Now you ask “is it visible?” If it’s visible, it belongs to a word. Anything else counts as a separator.
You add three spaces instead of one.
*/
