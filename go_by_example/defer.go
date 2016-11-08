package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("/tmp/asdfasjdfhjsk")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("something")
	fmt.Println("something")
	fmt.Println("something")
	fmt.Println("something")
}
