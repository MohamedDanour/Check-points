// no printing whole strings 
// no converting Integers to Ascii 

package main 

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, rune := range s {
		z01.PrintRune(rune)
	}
}


// can’t index an int
func printInt(n int) {

	integer_string := ""

	// convert the integer into the charachter version so we can add it to a string : '0' (eg. 48) + the integer value, is to convert to ascii
	for n > 0 {
		integer_string = string('0' + rune(n%10)) + integer_string 
		n = n / 10
	}

	for _, rune := range integer_string {
		z01.PrintRune(rune)
	} 
}
func main() {
	printStr("hello\n")
	printInt(42)
}