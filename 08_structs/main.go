package main

import "fmt"

// Fruit is a sample struct.
type Fruit struct {
	name  string
	count int
}

func main() {
	a := Fruit{name: "Apple", count: 12}
	fmt.Println(a)
}
