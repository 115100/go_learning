package main

import "fmt"

func add(op1, op2 int) int {
	return op1 + op2
}

func main() {
	x := 1
	y := 2

	z := add(x, y)

	fmt.Println(z)
}
