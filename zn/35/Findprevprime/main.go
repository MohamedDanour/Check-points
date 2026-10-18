// inferior means less than

package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}


func FindPrevPrime(nb int) int {
	for number := nb; number >= 2; number-- {
		isPrime := true
		for divisor := 2; divisor*divisor <= number; divisor++ {
			if number%divisor == 0 {
				isPrime = false // number divides evenly, so it's not prime
				break
			}
		}
		if isPrime {
			return number
		}
	}
	return 0 // no prime found (nb is less than 2)
}