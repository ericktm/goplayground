package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	// s = {1,2,3,4}
	if s == nil {
		fmt.Println("nil!!")
	}
}
