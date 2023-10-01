// go run 2_vars.go
package main

import "fmt"

// In this script we'll see how to declare functions
// - a function is declared with the 'func' keyword.
// the function name can be written in camelCase.
// - function names are usually lowercase, but inside a package
// they are public if they start with an uppercase letter
// - the function parameters are declared with the respective type
// - the function return type is declared after the parameters and before the curly braces, or empty if void
// e.g. this function returns an int 'func function(args...) int { ... }'
// checkout https://go.dev/ref/spec#Function_types for advanced function declarations

// docs for this function starts here,
// this function takes a string and shows it in the console with the fmt package,
// it doesn't return anything
func my_function(param string) {
	fmt.Println(param)
}

func main() {
	// call a function like this
	my_function("Hello World")

	// functions are first class citizens in golang
	// this means they represent a value and can be passed around
	// e.g. you can assign a function to a variable
	var f func(string) = my_function
	f("Hello from f")

	// you can also declare anonymous functions to encapsulate variables form the execution environment
	// this is called a closure, and this allows you to decorate other functions
	var a int = 10
	var g func(string) = func(param string) {
		fmt.Println(param, "with var", a)
		f("Hello from f from g")
	}
	g("Hello from g")

	// TODO: there are other example patterns you could do
}
