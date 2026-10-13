/*
Pseudocode: 

1. First we need to check if the string is empty, if so we return and empty string back.

2. Second we need to check if the Word IS NOT camelcase, on three conditions and return the same thing eg the original string:
• capital at the end if the letter at position len(s)-1 is in the range 'A'–'Z'
• two capitals in a row using a for loop here if positions i and i-1 are both capital, and also we start checking this at index 1 bcs we cannot start checking from 0 and then 0-1 index does not exist.
• Non letters if not between 'a'–'z' or 'A'–'Z' starting at index 0.

3. Third means that If it is not empty and the three conditions does not apply it is a camelcase string (either lower or upper does not matter). We will initialise an empty string and a for loop looping char by char. BECAUSE a underscore cannot come before the first letter even if it is capitalized we need to check that the char is not the first index and then that it is capitalized we will append as ”_”+ capitalized letter to the string, other chars added as they are. Last we return the string modified as snakecase.
*/



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

func CamelToSnakeCase(s string) string {
    if len(s) == 0 {
        return "" // return empty string
    }

    // capital at the end
    if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
        return s // return the string 
    }

    // not first index and two capitals in a row
    for index := range s {
        if index != 0 && (s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'A' && s[index-1] <= 'Z') { // not first and both capital
            return s // return the string 
        }
    }

    // non letters
    for _, character := range s {
        if !((character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')) { // NOT a letter
            return s // return the string 
        }
    }

    // IT IS camelCase
    camelToSnake := ""
    for index, character := range s {
        if index != 0 && (character >= 'A' && character <= 'Z') {
            camelToSnake += "_" + string(character)
        } else {
            camelToSnake += string(character)
        }
    }
    return camelToSnake // return snakecase
}