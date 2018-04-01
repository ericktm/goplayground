package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := make(map[string]int)

	for _, v := range strings.Fields(s) {
		words[v] = words[v] + 1
	}

	fmt.Println(words)

	return words
}

func main() {
	wc.Test(WordCount)
}
