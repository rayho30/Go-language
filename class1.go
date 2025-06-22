package main

import (
	"fmt"
)

func main() {
	// Integer types
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = -2147483648
	var i64 int64 = 9223372036854775807

	// Unsigned integers
	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	// Floating point numbers
	var f32 float32 = 3.14
	var f64 float64 = 2.7182818284

	// Complex numbers
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	// Boolean
	var b bool = true

	// String
	var s string = "Hello, Go!"

	// Byte (alias for uint8)
	var by byte = 'A'

	// Rune (alias for int32, used for characters)
	var r rune = '☀'

	// Print all values
	fmt.Println("Integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned Integers:", ui, ui8, ui16, ui32, ui64)
	fmt.Println("Floats:", f32, f64)
	fmt.Println("Complex numbers:", c64, c128)
	fmt.Println("Boolean:", b)
	fmt.Println("String:", s)
	fmt.Println("Byte:", by)
	fmt.Println("Rune:", r)
}
