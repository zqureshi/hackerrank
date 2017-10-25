package main

import (
	"fmt"
)

func main() {
	var sockCount int
	fmt.Scanf("%d", &sockCount)

	var sockType int
	counts := make([]int, 101)
	for i := 0; i < sockCount; i++ {
		fmt.Scanf("%d", &sockType)
		counts[sockType]++
	}

	pairs := 0
	for _, c := range counts {
		pairs += c / 2
	}

	fmt.Println(pairs)
}
