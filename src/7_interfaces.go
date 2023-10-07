package main

import "fmt"

// In this script we'll review how to approach OOP inheritance with interfaces in golang

// - golang doesn't have classes, but it has structs
// - golang favors composition over inheritance
// - an interface defines a set of abstract methods, a struct cannot define a method unless you implement it
// - you could say a struct implements the interface if it has all the methods defined by the interface
// - interfaces can be used as a parent-like class and allows us to approach polymorphism
//    - this means we can pass any struct that implements the interface to a function that expects the interface

// for now we'll focus on this example
// lets create an interface for animals
type Animal interface {
	walk() string
}

// Duck doesn't define the walk method, on declaration
type Duck struct {
	name string
}

// but it does implement the walk method, thus it implements the Animal interface
func (p Duck) walk() string {
	return "I walk like a duck 🦆"
}

// Same for Dog
type Dog struct {
	name string
}

// but it does implement the walk method, thus it implements the Animal interface
func (s Dog) walk() string {
	return "I walk like a dog 🐕"
}

// lets create a function that takes an Animal interface
func lets_walk(a Animal) string {
	// this will call the walk method of the Animal interface, it doesn't matter if it's a Duck or a Dog
	return a.walk()
}

func main() {
	var a Animal = Duck{name: "Duck"} // a is an Animal interface, but it's a Duck instance underneath
	var b Animal = Dog{name: "Dog"}   // b is an Animal interface, but it's a Dog instance underneath

	fmt.Println(lets_walk(a))
	fmt.Println(lets_walk(b))
}
