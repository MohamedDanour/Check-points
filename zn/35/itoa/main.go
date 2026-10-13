// transforms a number represented as anint in a number represented as a string




package main

import (
	"fmt"
)

func Itoa(n int) string {
	// noll check 
	if n == 0 {
		return "0"
	}
	
	// negative check
	isNegative := false 
	if n < 0 {
		isNegative = true 
		n = -n
	}

	integer_string := ""

	// convert the integer into the charachter version so we can add it to a string : '0' (eg. 48) + the integer value, is to convert to ascii
	for n > 0 {
		integer_string = string('0' + rune(n%10)) + integer_string 
		n = n / 10
	}

	if isNegative {
		integer_string = "-" + integer_string
	}

	return integer_string
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}



/*
package main 

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, rune := range s {
		z01.PrintRune(rune)
	}
}
*/





