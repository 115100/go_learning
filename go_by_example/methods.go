package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	square := rect{height: 10, width: 10}

	fmt.Println("Area:", square.area())
	fmt.Println("Perimeter:", square.perim())
}
