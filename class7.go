package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 20
	b = 30
)

func main() {
	fmt.Println("Showing custom Package")
	mathlib.Add(a, b)
	mathlib.Sum()
}
