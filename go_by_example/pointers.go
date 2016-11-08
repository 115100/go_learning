package main

import "fmt"

func noPointer(i int) {
	i = 5
}

func Pointer(i *int) {
	*i = 5
}

func main() {
	i := 1
	fmt.Println("Address of i:", &i)

	// Call without pointer
	noPointer(i)
	fmt.Println(i)

	// Call with pointer
	Pointer(&i)
	fmt.Println(i)
}
