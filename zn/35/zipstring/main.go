package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}


func ZipString(s string) string {
	if s == "" {
		return ""
	}

	result := ""
	counter := 1 // start from 1 bcs a single letter gives 1

	for index := 1; index <= len(s); index++ {

		// keep counting while the letter is the same
		if index < len(s) && s[index] == s[index-1] { 
			counter++
			continue
		}

		// convert counter (int) to text (string), stor it in "number"
		number := "" 
		converting_conter := counter // store it again
		for converting_conter > 0 {
			number = string(rune('0' + converting_conter % 10)) + number 
			converting_conter = converting_conter / 10                         
		}

		// write the group of int + string to result variable
		result += number + string(s[index-1]) 
		counter = 1 // start fresh FROM 1
	}

	return result
}


/*
Pseudocode: 
Input: a string. 
Output: a string but the occurance of the charachter consequently (a number) and the charachter itself one time (string)
Each group of identical characters directly next to each other becomes: count + character.
Spaces and symbols are treated exactly like letters and get also counted

Edge cases:
Only chars directly after each other are are counted together "aba" gives 1a1b1a, not 2a1b. So in "aba", the two a's are NOT counted together.

If the string is empty, return "". Initialize variables with string result = "" and counter = 1. I´ll use a foor loop, looping from the 2nd character to one step past the end: 
If the character is the same as the previous one we increase 1. Otherwise it means new character or end of string. I will also have to convert THE COUNTER to text manually TO USE IT IN THE RETURN STRING (% 10 and / 10). 
I´LL Add that number and the just the previous character to result, then Reset counter to 1.
Finally Return the result.
*/