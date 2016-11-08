package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	lineCount := flag.Bool("l", false, "Count lines.")
	flag.Parse()

	if !(*lineCount) {
		fmt.Println("This program only counts lines.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}

	fmt.Println(lines)
}
