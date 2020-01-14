package main

import (
	"fmt"
	"strings"
)

func Solution(N int) (total int) {
	current := 0

	split := strings.Split(AsBinary(N), "")
	fmt.Printf("%v\n", split)

	for _, value := range split {
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
