package main

import "fmt"

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("sync")

	go f("async")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
