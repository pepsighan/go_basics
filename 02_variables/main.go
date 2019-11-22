package main

import "fmt"

func main() {
	const num int = 3

	var flt float64 = 3.2

	var truthy bool = true

	var char rune = 'a'

	var str string = "a"

	var arr []string = []string{"a", "b", "c"}

	fmt.Println(num)
	fmt.Println(flt)
	fmt.Println(truthy)
	fmt.Println(char)
	fmt.Println(str)
	fmt.Println(arr)

	zeroValue()
}

// Unassigned variables have zero values.
// Most of the primitive types have non-nil zero values and all reference types have nil as its zero value.
func zeroValue() {
	var num int
	var flt float64
	var falsy bool
	var char rune
	var str string
	var arr []string
	var ref *string

	fmt.Println(num == 0)
	fmt.Println(flt == 0)
	fmt.Println(falsy == false)
	fmt.Println(char == 0x00)
	fmt.Println(str == "")
	fmt.Println(arr == nil)
	fmt.Println(ref == nil)
}
