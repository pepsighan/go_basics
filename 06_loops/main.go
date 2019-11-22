package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		i++
		fmt.Println(i)

		if i > 10 {
			break
		}
	}

	obj := []string{"a", "b", "c"}
	// obj := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	for k, v := range obj {
		fmt.Println(k, " -> ", v)
	}
}
