package main

import (
	"fmt"
)

// this function will return two integers, x and y
// they are named. when call return, the current values will be returned

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	// 9,44 is converted to 10
	// 7,55 is converted do 7
}
