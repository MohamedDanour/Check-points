package main

import (
	"fmt"
)

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	} else {
		height := h * 2
		width := w * 2
		perimeter := height + width 
		return perimeter
	} 
	
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}