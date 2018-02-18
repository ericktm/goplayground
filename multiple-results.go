package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	// the parentesis can be ommited
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
