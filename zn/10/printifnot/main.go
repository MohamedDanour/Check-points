package main

import (
	"fmt"
)

func PrintIfNot(str string) string {
	if len(str) < 3 || len(str) == 0 {
		return "G" + "\n"
	} else {
		return "Invalid Input" + "\n"
	}
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}