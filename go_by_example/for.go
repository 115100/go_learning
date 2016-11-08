package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// C-style init/cond/after
	for j := 9000; j <= 9010; j++ {
		fmt.Println(j)
	}

	// Infinite loop
	for {
		fmt.Println("loop")
		break
	}
}
