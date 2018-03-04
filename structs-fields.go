package main

import (
	"fmt"
)

type Vertext struct {
	X int
	Y int
}

func main() {
	v := Vertext{1, 2}

	v.X = 4

	fmt.Println(v.X)
}
