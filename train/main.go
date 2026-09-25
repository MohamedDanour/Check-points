package main

import (
	"fmt"
	"train"
)

func main() {
	fmt.Println(train.Gcd(42, 10))
	fmt.Println(train.Gcd(42, 12))
	fmt.Println(train.Gcd(14, 77))
	fmt.Println(train.Gcd(17, 3))
}
