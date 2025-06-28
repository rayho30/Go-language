package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}

func main() {
	var a,b int
	fmt.Println("Enter Number a: ")
	fmt.Scanln(&a)
	fmt.Println("Enter Number b: ")
	fmt.Scanln(&b)
	//a := 10
	//b := 20
	add(a, b)
	//add(5, 7)
}
