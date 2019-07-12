package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution(N int) (total int) {
	current := 0

	for _, value := range strings.Split(AsBinary(N), "") {
		if value == "0" {
			current += 1
		} else {
			if current > total {
				total = current
			}

			current = 0
		}
	}

	return
}

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
