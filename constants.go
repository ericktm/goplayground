package main

import (
	"fmt"
)

const Pi = 3.14

func main() {

	// constants cannot be declared using :=
	const World = "EERRT"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
