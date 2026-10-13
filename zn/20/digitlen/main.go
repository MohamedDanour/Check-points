package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	if !(base >= 2 && base <= 36) {
		return -1
	} 

	if n < 0 {
		n = -n
	}

	num_of_divs := 0

	for n > 1 {
		n = n / base
		num_of_divs += 1
	} 

	return num_of_divs

}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

/*
ex. Start with n = 100.
n/base
100 / 10 = 10
10 / 10 = 1
1 / 10 = 0
*/

