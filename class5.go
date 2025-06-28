package main

import "fmt"

// getnumber returns sum and product of two integers
func getnumber(num1 int, num2 int) (int, int) {
    sum := num1 + num2
    mul := num1 * num2
    return sum, mul
}

func main() {
    var a, b int

    fmt.Print("Enter first number: ")
    fmt.Scanln(&a)

    fmt.Print("Enter second number: ")
    fmt.Scanln(&b)

    sum, product := getnumber(a, b)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
