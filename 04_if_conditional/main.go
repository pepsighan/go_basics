package main

import "fmt"

func main() {
	a := true

	if a {
		fmt.Println("a is true")
	}

	if !a {
		// Not gonna run
	} else {
		fmt.Println("!a is false")
	}

	b := false
	if a {
		fmt.Println("a is true")
	} else if b {
		// Not gonna run
	} else {
		// Not gonna run
	}

	// b := 1

	// if b {
	// 	fmt.Println("Does not run")
	// }
}
