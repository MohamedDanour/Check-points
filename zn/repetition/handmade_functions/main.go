// no printing whole strings 
// no converting Integers to Ascii 


package main 

import "github.com/01-edu/z01"

func printStr(s string) string {
	for _, rune := range s {
		z01.PrintRune(rune)
	}
}

func main() {
	printStr("hello\n")
}