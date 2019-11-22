package main

import "fmt"

func main() {
	a := 1      // value
	b := "name" // value

	c := &a // Reference that points to a
	d := b  // Copied value

	*c = 5
	fmt.Println(a)  // Prints 5
	fmt.Println(*c) // Prints 5

	d = "another"
	fmt.Println(b) // Prints name
	fmt.Println(d) // Prints another
}
