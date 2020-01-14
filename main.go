package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Argument not a valid number")
		os.Exit(1)
	}

	fmt.Printf("Number of sequential zeros: %d\n", Solution(num))
}

