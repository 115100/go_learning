package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "The"
	s[1] = "chocolate"
	s[2] = "teapot"
	fmt.Println("s:", s)
	fmt.Println("Teapot:", s[2])

	s = append(s, "is")
	s = append(s, "very", "useless")
	fmt.Println("Full sentence:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c full sentence:", c)

	l := s[1:3]
	fmt.Println("Sliced sentence:", l)

	l = s[:3]
	fmt.Println("Less sliced sentence:", l)

	l = s[1:]
	fmt.Println("Tedious additional slice:", l)

	// 2x3 slice
	twoD := make([][]int, 2)
	for i := 0; i < 2; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
