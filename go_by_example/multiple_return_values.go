package main

import "fmt"

func foo() (string, string) {
	return "bar", "baz"
}

func main() {
	x, y := foo()

	fmt.Println(x)
	fmt.Println(y)
}
