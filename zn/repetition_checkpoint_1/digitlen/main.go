package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

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