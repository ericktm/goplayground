package main

import (
	"fmt"
)

func fibonacci() func() int {

	prev := 1
	current := 0

	return func() int {
		prev, current = current, prev+current
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
