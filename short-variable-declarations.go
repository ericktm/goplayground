package main

import "fmt"

// cannot assign value to variable outside a function Scope
// var z := 3
// e := "name"

func main() {
	var i, j int = 1, 2
	k := 3
	// implicit types (bool,bool, string)
	c, python, java := true, false, "no!!!!"

	fmt.Println(i, j, k, c, python, java)
}
