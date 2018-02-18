package main

import (
	"fmt"
)

// https://golang.org/ref/spec#The_zero_value

// declaring three variables of type boolean
// if not defined, They assume the value false (Zero Value = not true)
var c, python, java bool

func main() {
	// c = true
	// if not defined, assume value Zero
	var i int
	fmt.Println(i, c, python, java)
}
