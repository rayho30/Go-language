package main

import (
	"fmt"
)

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("You are okay to marry.")
	} else if age < 18 {
		fmt.Println("You are not okay to marry.")
	} else {
		fmt.Println("You are okay to love.")
	}
}
