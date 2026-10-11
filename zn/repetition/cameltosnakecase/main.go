package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{
	if len(str) == 0 {
		return ""
	}

	// If the string is NOT camelCase, return the string unchanged. eg:

	// last letter capital
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return 
	} 

	// two following letters capital
	for index, char := range str {
		if index != 0 && (s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'A' && s[index-1] <= 'Z') {
			return s
		}
	} 

	// nn letters
	for _, char := range s {
		if !((charachter >= 'a' && charachter <= 'z') || (charachter >= 'A' && charachter <= 'Z')) {
			return s
		}
	}

	// its camelcase

	


}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

