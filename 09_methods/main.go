package main

import "fmt"

// Fruit is a sample struct.
type Fruit struct {
	name  string
	count int
}

func (f Fruit) print() {
	fmt.Println("name =", f.name, " count =", f.count)
}

func main() {
	a := Fruit{name: "Apple", count: 12}
	a.print()
}
