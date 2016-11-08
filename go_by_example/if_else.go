package main

import "fmt"

func main() {

	if true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	if true {
		fmt.Println("True")
	}

	if true {
		fmt.Println("True")
	} else if false {
		fmt.Println("False")
	} else {
		fmt.Println("Something else")
	}
}
