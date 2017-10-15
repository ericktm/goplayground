package main

import (
	"fmt"
	"api/model"
	"api/extra"
)

func main() {
	b := model.Tag{"123", "2345"}
	//rel := Tag
	fmt.Println("starting", b.ToString)

	extra.Balance(123)
}
