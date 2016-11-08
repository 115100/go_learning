package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, _ := range nums {
		fmt.Println("index:", i)
	}

	kvs := map[string]int{"a": 1, "b": 2}
	for k, v := range kvs {
		fmt.Printf("%s: %d\n", k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
