/*
Pseudocode: "Divides" means there's no remainder when dividing by a number. Use modulus (%). No division by 0!

If an input is 0, return 0. If both numbers are equal, that same number is the GCD. (This happens first.)
Initialize two empty lists.
Loop through numbers starting from 1, increasing one step at a time, up to and including the number.
If our number divides evenly by the looped number, append the looped number to the list.
Do this for both numbers.
First, bring out the numbers that appear in both lists.
Then, the largest one of those is the GCD.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint { // uint means non-negative integer (0 included)
	if a == 0 || b == 0 {
		return 0
	} else if a == b {
		return a // does not matter which one
	}

	// numbers that divide
	var a_list []uint // empty slices need a type, and variable declaration this way
	var b_list []uint

	for index := uint(1); index <= a; index++ {
		if a%index == 0 {
			a_list = append(a_list, index)
		}
	}

	for index := uint(1); index <= b; index++ {
		if b%index == 0 {
			b_list = append(b_list, index)
		}
	}

	// bring out common numbers
	var common []uint
	for _, a_number := range a_list {
		for _, b_number := range b_list {
			if a_number == b_number {
				common = append(common, a_number)
			}
		}
	}

	// bring out the largest common one
	var largest uint
	for _, number := range common {
		if number > largest {
			largest = number
		}
	}
	return largest
}