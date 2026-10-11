package main

import "github.com/01-edu/z01"

func CheckNumber(arg string) bool {
	aNumber := false
	for _, char := range arg {
		if char >= '0' && char <= '9' {
			aNumber = true
			return aNumber
		} else {
			continue
		}
	}
	return aNumber
}

func main() {
	if CheckNumber("Hello") {
		z01.PrintRune('t')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
	} else {
		z01.PrintRune('f')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('s')
		z01.PrintRune('e')
	}
	z01.PrintRune('\n')

	if CheckNumber("Hello1") {
		z01.PrintRune('t')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
	} else {
		z01.PrintRune('f')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('s')
		z01.PrintRune('e')
	}
	z01.PrintRune('\n')
}