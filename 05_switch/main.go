package main

import "fmt"

func main() {
	a := "apple"

	switch a {
	case "apple":
		fmt.Println("It is an apple")
		// fallthrough
	case "ball":
		fmt.Println("It is a ball")
	default:
		fmt.Println("It is something else")
	}
}
