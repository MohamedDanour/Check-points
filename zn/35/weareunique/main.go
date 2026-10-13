/*
If both strings are empty return -1. 
Initialise a counter at 0. 
Then use a for loop looping char by char through str1. First check if we have ALREADY SEEN this char earlier in str1, by looping through all the chars before it. If we seen it, we skip it and go to the next char, so repeated chars only count as 1.

If it is a new char, loop through str2 COMPARING it with every char there. If we find in str2, eg it is shared, do nothing. If don't find it in str2, it is NOT shared, so increase the counter with 1.

Then do the exact same thing the other way around, so loop through str2, skip the chars already seen in str2, and increase the counter with 1 for every char that is NOT in str1.

Finally return the counter. If nothing was unique, the counter stays 0, so I return 0.

*/



package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	counter := 0
	for index := 0; index < len(str1); index++ { // s[i] gives the character (a byte) at position i.
		
		seen := false
		for j := 0; j < index; j++ {
			if str1[j] == str1[index] {
				seen = true
			}
		}

		found := false
		for k := 0; k < len(str2); k++ {
			if str2[k] == str1[index] {
				found = true
			}
		}
		if !seen && !found {
			counter++
		}
	}
	return counter
}