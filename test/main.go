package main

import (
	"fmt"
	"strconv"
)


func ZipString(s string) string{
	dict := make(map[string]int)
	sol := ""
	for _,c := range s{
		dict[string(c)]+=1
	}
	for i:=0; i<len(s);i++{
		sol+=string(s[i])
		sol+=strconv.Itoa(dict[string(s[i])] )
	}
	return sol
}


func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}