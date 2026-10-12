package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	// last letter capital -> not camelCase
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}

	// two capitals in a row -> not camelCase
	for index := range s {
		if index != 0 && (s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'A' && s[index-1] <= 'Z') {
			return s
		}
	}

	// any non-letter -> not camelCase
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return s
		}
	}

	// it's camelCase: put _ before each capital (except the first char)
	camelToSnake := ""
	for index, char := range s {
		if index != 0 && (char >= 'A' && char <= 'Z') {
			camelToSnake += "_"
		}
		camelToSnake += string(char)
	}
	return camelToSnake
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}