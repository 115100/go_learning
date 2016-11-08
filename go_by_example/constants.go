package main

import (
	"fmt"
	"math"
)

const s string = "The quick brown fox jumps over the lazy dog."

func main() {
	fmt.Println(s)

	const n = 9001

	const d = (900100 / n) * 1e9
	fmt.Println(d)

	// Explicit cast to a Go type
	fmt.Println(int64(d))

	fmt.Println(math.Sin(d))
}
