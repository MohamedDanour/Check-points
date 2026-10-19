package main

import (
	"fmt"
)

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n" 
	} 

	new_string := ""
	

	for index := 2; index < len(str); index += 3 {
		new_string += string(str[index])
	}
	
	return new_string + "\n"
}

