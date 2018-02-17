package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// if the seed is not specified, the same number is returned at each program execution
	rand.Seed(time.Now().UnixNano())

	fmt.Println("my favorite number is", rand.Intn(10))
}
