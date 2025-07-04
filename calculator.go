package main

import (
    "fmt"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) float64 {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return a / b
}

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)

    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)

    var result float64

    switch operator {
    case "+":
        result = add(num1, num2)
    case "-":
        result = subtract(num1, num2)
    case "*":
        result = multiply(num1, num2)
    case "/":
        result = divide(num1, num2)
    default:
        fmt.Println("Invalid operator")
        return
    }

    fmt.Printf("Result: %.2f\n", result)
}
