// go run hw.go
package main

import "fmt"

func main() {
	// golang was created by google in 2009, the creator is Rob Pike
	// golang is a language with a C-like syntax
	// it is compiled and statically typed, this means it talks directly to the machine
	// and you have to specify the type of the variables, variables cannot change type while running
	// you can build an executable with 'go build'
	// or run it directly with 'go run' i.e. 'go run 1_hello_world.go'
	// golang supports concurrency and parallelism naturally with goroutines

	// the first executed function is main and underneath it's a goroutine
	// golang uses the builtin fmt package for printing and formatting
	// checkout https://pkg.go.dev/fmt
	fmt.Println("Hello World")

	// Advanced tip: you can checkout the language specificaiton and grammar here https://go.dev/ref/spec
}
