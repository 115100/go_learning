package main

import "fmt"

func main() {
	c1 := make(chan string, 3)

	c1 <- "j1"
	c1 <- "j2"
	c1 <- "j3"
	close(c1)

	for elem := range c1 {
		fmt.Println(elem)
	}
}
