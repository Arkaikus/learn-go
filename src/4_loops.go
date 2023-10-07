package main

import "fmt"

// In this scripts we will review the different types of loops in golang
// basically golang only uses a for loop, but it has different ways to use it

func basic() {
	// the basic for loop, similar to C/Java/JS etc
	// for init; condition; post { ... }
	// init: executed before the first iteration
	// condition: evaluated before every iteration
	// post: executed at the end of every iteration
	// for i := 0; i < 10; i++ { ... }
	// this defines a i variable with value 0, and it will loop while i < 10 while incrementing i by 1
	for i := 0; i < 10; i++ {
		fmt.Println("basic", i)
	}
}

func while() {
	// in the for loop syntax the init and post statements are optional
	// this is makes it equivalent to a while loop
	// for condition { ... }
	i := 0
	for i < 10 {
		fmt.Println("while", i)
		i++
	}
}

func while_true() {
	// if you omit the condition it will loop forever, equivalent to a while(true) loop
	// the empty condition is equivalent to condition=true
	// for { ... }
	i := 0
	for {
		fmt.Println("while_true", i)
		i++
		if i > 10 {
			break // you can break out of the loop with the break keyword
		}
	}
}

func foreach() {
	// golang doesn't have a foreach loop, but you can use the range keyword (like python)
	// range will iterate over arrays, slices, maps, strings, and channels (more on this later)
	// for index, value := range array { ... }
	// for index, value := range slice { ... }
	// for key, value := range map { ... }
	// for index, value := range string { ... }
	items := []string{"a", "b", "c"}
	for index, value := range items {
		fmt.Println("foreach_range_slice", index, value)
	}

	itemap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range itemap {
		fmt.Println("foreach_range_map", key, value)
	}

	// you can ommit the index or the value using the wildcard _ character
	// for _, value := range array { ... }
	// for index, _ := range array { ... }
	// for key, _ := range map { ... }
	// for _, value := range string { ... }
	for index, _ := range items {
		fmt.Println("foreach_range_value", index)
	}
	for _, value := range items {
		fmt.Println("foreach_range_value", value)
	}
}

func main() {
	basic()
	while()
	while_true()
	foreach()
}
