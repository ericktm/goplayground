package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var k float64 = float64(x + y)
	var z uint = uint(f)
	fmt.Println(x, y, z, f, k)
	fmt.Printf("%T %T %T %T\n", x, y, z, f)
}
