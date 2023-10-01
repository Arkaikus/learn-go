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
}
