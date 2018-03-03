package main

import "fmt"

func main() {
	defer fmt.Println("word")

	fmt.Println("hello")
}
