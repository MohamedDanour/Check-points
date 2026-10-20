package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}




func CamelToSnakeCase(s string) string{
	if len(s) == 0 {
		return ""
	} 
	
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}
	
	for index, _ := range s {
			if index != 0 && (s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'A' && s[index-1] <= 'Z') {
				return s
			}
		}

    for _, character := range s {
        if !((character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')) { // NOT a letter
            return s 
        }
    }

	// string is camelcase
	snake_case := ""
		for index, char := range s {
			if index != 0 && (char >= 'A' && char <= 'Z') {
				snake_case += "_" +  string(char+32)
			} else {
				snake_case += string(char)
			}	
		}
	return snake_case
}

