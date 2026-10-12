package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	stringLenght := len(str)
	
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	} else if stringLenght % 2 == 0 {         //even 
		middle := stringLenght / 2
		return str[:middle]
	} else { // odd
		middle := stringLenght / 2 - 1
		return str[:middle]
	}
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

