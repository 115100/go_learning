package main

import (
	"fmt"
)

func main() {
	i := 1

	for i <= 50 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

		i += 1
	}
}
