package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 1)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}

		text := string(p[:n])

		fmt.Println(strings.ToUpper(text))
	}

	for {
		time.Sleep(5 * time.Duration(time.Millisecond))

		// fmt.Printf("time: %s \n", time.Now().String())
	}
}
