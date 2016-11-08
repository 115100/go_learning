package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("No message received.")
	}

	// Non-blocking send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent:", msg)
	default:
		fmt.Println("No message sent.")
	}

	// Multiple case non-blocking receives
	select {
	case msg := <-messages:
		fmt.Println(msg)
	case signal := <-signals:
		fmt.Println(signal)
	default:
		fmt.Println("No activity received.")
	}
}
