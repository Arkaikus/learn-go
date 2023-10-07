package main

import "fmt"

// In this script we'll review how to approach OOP with structs in golang
// - golang doesn't have classes, but it has structs
// - structs are declared with the 'type' keyword
// - structs should define fields
// - structs can have methods
// - structs can have structs as fields (composition)
// - structs can implement interfaces (more on this later)
// for now we'll focus on this example

type person struct {
	// by default all fields are public
	// a public/private scope can be achieved at the level of packages
	// inside the package you can lowercase the first letter of the type name for private
	// and uppercase the first letter of the type name for public
	// https://stackoverflow.com/a/22148186
	name string
	age  int
}

// methods for the person struct are declared like this
// func (t T) methodName(args...) (T') { ... } where (t T) refers to the type, and T' is the return type or empty for void
func (p person) greet() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age)
}

// if you want to implement a setter method you should use pointers to access the self values in memory
// if not go will create a copy of the struct instance and the changes will be applied to the copy
// e.g func (self person) setAge(age int) { // (self person) would be a copy of the instance, different memory address
func (self *person) setAge(age int) { // (self *person) is a pointer to the instance, same memory address
	// if you compare it with python then it will be like 'def setAge(self: person, age): self.age = age'
	self.age = age
}

// you could also override the default print format with
// this means fmt.Println(p) will run p.String() underneath
func (p person) String() string {
	return "Hello, my name is " + p.name + " and I am " + fmt.Sprint(p.age)
}

func main() {
	p := person{name: "Alice", age: 30}
	p.greet()      // Hello, my name is Alice and I am 30
	fmt.Println(p) // Hello, my name is Alice and I am 30

	// if setAge is not declared with a pointer this wont work as expected
	// p.setAge(31) // p is copied, and p.age=31 is applied to the copy
	// p.greet() // this will still print "Hello, my name is Alice and I am 30"

	// if setAge is declared with a pointer this will work as expected
	p.setAge(31) // p will be passed down as a pointer, and p.age=31 is applied to itself
	p.greet()    // thus this will print "Hello, my name is Alice and I am 31"
}
