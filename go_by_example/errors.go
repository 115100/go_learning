package main

import (
	"fmt"
)

type myError struct {
	err string
}

func (e *myError) Error() string {
	return e.err
}

func foo() (int, error) {
	return 1, &myError{err: "AHHHHHHHHHHHHHHHHHHHHHHHHHH"}
}

func main() {
	if _, err := foo(); err != nil {
		fmt.Println(err)
	}
}
