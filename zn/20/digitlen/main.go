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

	count := 0

	for n > 0 {
		n = n / base
		count += 1
	}
	return count 
	
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

