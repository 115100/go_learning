package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "z", "d", "c"}

	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Ints are sorted:", s)
}
