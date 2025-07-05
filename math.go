package mathlib

import "fmt"

func Add(x int, y int) int {
	result := x + y
	fmt.Println("Sum from Add:", result)
	return result
}

func Sum() {
	Add(4, 7) // No arguments expected when you call Sum from main
}
