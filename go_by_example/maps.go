package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	fmt.Println("map:", m)

	v1 := m["a"]
	fmt.Println("m[\"a\"]:", v1)

	fmt.Println(len(m))

	delete(m, "a")
	fmt.Println("map:", m)

	_, prs := m["asdf"]
	if !prs {
		fmt.Println("Key does not exist!")
	}

	n := map[int]int{1: 1, 2: 2}
	fmt.Println("intmap:", n)
}
