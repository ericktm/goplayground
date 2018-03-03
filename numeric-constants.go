package main

import (
	"fmt"
)

// 0 = 0
// -- 1 = 1
// 10 = 2
// -- 11 = 3
// 100 = 4
// -- 101 = 5
// -- 110 = 6
// -- 111 = 7
// 1000 = 8
// --- 1001 = 9
// --- 1010 = 10
// --- 1011 = 11
// --- 1100 = 12
// --- 1101 = 13
// --- 1110 = 14
// --- 1111 = 15
// 10000 = 16

const (
	// create a huge number by shifting 1 bit left 100 places.
	// in other words, the binary number thats is 1 followed by 100 zeroes.
	Big = 1 << 100

	//" shift it right again 99 places, so we end up with 1<<1, or 2."
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// fmt.Println("big", Big)
	fmt.Println("small", Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
