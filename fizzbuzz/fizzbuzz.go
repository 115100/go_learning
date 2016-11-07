package main

import (
	"fmt"
)

func main() {
	i := 1

	for i <= 50 {
		switch {
		case i%15 == 0:
			fmt.Println("Fizzbuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

		i += 1
	}
}
