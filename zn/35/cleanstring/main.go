// IMPORTANT: 
// A function gets its input from whoever calls it in code. 
// A program gets its input from the command line.
// Everything after go run . is an argument

// ALMOST ALMOST LIKE FIRSTWORD, BUT DONT BREAK WHEN A WORD IS completed and collected

package main 

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println() // ny rad
        return
    }

    fmt.Println(CleanString(os.Args[1])) // kör funktionen CleanString
}

func CleanString(s string) string {
	new_string := ""
	need_space := false //mellanslag/tab efter ett ord?

	for _, char := range s {
		if char == ' ' || char == '\t' {
			if len(new_string) > 0 {
				need_space = true
			}
		} else {
			if need_space {
				new_string += " "
				need_space = false
			}
			new_string += string(char)
		}
	}
	return new_string
}