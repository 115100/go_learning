package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend. Stop working!")
	default:
		fmt.Println("Get back to work!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() >= 18:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good afternoon.")
	}
}
