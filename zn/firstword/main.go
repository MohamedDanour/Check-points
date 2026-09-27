/* 
Pseudocode: 
Takes a string and returns the strings first word and a newline. A word starts at the first NON-SPACE CHARACHTER and ends exactly before the next space, OR at the end of the string. 
I will start by initiating an empty string variable and a for loop, looping through the string char by char. 

In the loop we start checking for non space chars to find the beginnnng of a word. 
OBS! do not check for letters bcs "..." in the test output counts as part of a word. 

If the char is a space charachter and the word is still empty, we skip it and just continue looping (bcs spaces can be in the beginning of a word as in the test output)
If the char is a space charachter and the word is not empty (eg. we started filling a word, we need to know when we are going to stop) so hitting a space now means we stop. 
If it is a non-space charachter we append it to the string variable 

Lastly we return the word followed by a newline


logs: 
run go mod init firstword (eg folder name)
*/

package main

import (
    "fmt"
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
    word := ""

	for _, charachter := range s {
		if charachter == ' ' && len(word) == 0 { // space charachter and the word is still empty
			continue
		} else if charachter == ' ' && len(word) > 0 { // space charachter and the word is not empty
			break 
		} else {
			word += string(charachter)
		}
	}
	return word + "\n"
}
