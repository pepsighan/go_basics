package main

import (
	"errors"
	"fmt"
	"log"
)

func errorableFunction() (int, error) {
	// Errors are returned not thrown.
	return 0, errors.New("bad things happen")
}

func main() {
	val, err := errorableFunction()
	if err != nil {
		log.Fatal("Error occurred: ", err)
		// Fatal prints the error and exits the program.
	}
	fmt.Println(val)
}
