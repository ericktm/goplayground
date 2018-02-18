package main

import "fmt"

// when all functions parameters are the same type, you can specify only the last
// x and y are both of int type
func add(x, y int) int {
	return x + x
}

func main() {
	fmt.Println(add(42, 13))
}
