// transforms a number represented as anint in a number represented as a string

package main

import "fmt"

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}


func Itoa(n int) string {

	integer_string := ""

	if n == 0 {
		return "0"
	}

	is_negative := false

	if n < 0 {
		is_negative = true
		n = -n // make n positive so loop runs later
	}

	for n > 0 {
		integer_string = string('0'+rune(n%10)) + integer_string
		n = n / 10
	}

	if is_negative {
		integer_string = "-" + integer_string // add the sign once, at the front AFTER LOOP 
	}

	return integer_string
}


