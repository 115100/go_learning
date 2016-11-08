package main

import "fmt"

func var_sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(var_sum(1, 2, 3, 4))

	var_nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(var_sum(var_nums...))
}
