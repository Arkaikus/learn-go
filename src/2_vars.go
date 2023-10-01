// go run 2_vars.go
package main

import "fmt"

func main() {
	// In this script there are different ways to declare variables, inspect them and print them
	// constant a with type int
	const a int = 10 // cannot assign 'a=11' later

	// assignable variable - declaration and assignment
	var b = 3.14

	// declaration and assignment (short variable declaration clause)
	c := 1e13

	// multiple declaration and assignment in one line
	d, e := 1, 2

	// you can print format and inspect variables this way
	// %s for strings types
	// %d for known int types, supports int only
	// %v for unknown/any type
	// %T to print type
	fmt.Printf("a: %d %v %T\n", a, a, a)
	fmt.Printf("b: %d %v %T\n", b, b, b) // shows %!d(float64=3.14) warn for unsupported type
	fmt.Printf("c: %d %v %T\n", c, c, c) // shows %!d(float64=1e+13) warn for unsupported type
	fmt.Printf("d: %d %v %T\n", d, d, d)
	fmt.Printf("e: %d %v %T\n", e, e, e)

	// checkout https://go.dev/ref/spec#Assignment_statements

	// There are different types of variables in golang
	var integer int = 42
	var int8var int8 = 127                   // can hold -128 to 127
	var int16var int16 = 32767               // can hold -32768 to 32767
	var int32var int32 = 2147483647          // can hold -2147483648 to 2147483647
	var int64var int64 = 9223372036854775807 // can hold -9223372036854775808 to 9223372036854775807

	// Floating-point data types
	var float32var float32 = 3.14          // can hold 1.18*10^-38 to 3.4*10^38
	var float64var float64 = 3.14159265359 // can hold 2.23*10^-308 to 1.80*10^308

	// String data type
	var stringvar string = "Hello, World!"

	// Boolean data type
	var boolvar bool = true

	// Character data type (not a primitive type in Go)
	var runevar rune = 'A'

	// Byte data type (alias for uint8)
	var bytevar byte = 65

	// Complex data types
	// go can handle complex numbers inherently
	var complex64var complex64 = 3.0 + 4.0i
	var complex128var complex128 = 1.0 + 2.0i

	// Pointer data types
	// & is the address of operator
	// * is the dereference operator, returns the value in the address
	var intPointer *int = &integer

	// Array and Slice data types
	// arrays are fixed length, slices are dynamic
	// you can add a new element to a slice with append(slice, element)
	var intArray [3]int = [3]int{1, 2, 3}
	var intSlice []int = []int{4, 5, 6}

	// Maps data type
	// maps are like dictionaries in python
	// maps are unordered
	var intMap map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(integer, int8var, int16var, int32var, int64var)
	fmt.Println(float32var, float64var)
	fmt.Println(stringvar)
	fmt.Println(boolvar)
	fmt.Println(runevar)
	fmt.Println(bytevar)
	fmt.Println(complex64var, complex128var)
	fmt.Println(*intPointer) // access the value in the &integer address
	fmt.Println(intArray, intSlice)
	fmt.Println(intMap)
}
