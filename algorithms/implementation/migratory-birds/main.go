package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scanf("%d", &count)

	counts := []int{-1, 0, 0, 0, 0, 0}
	var birdType int
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &birdType)
		counts[birdType]++
	}

	idx, max := 0, 0
	for id, count := range counts {
		if count > max {
			idx, max = id, count
		}
	}

	fmt.Println(idx)
}
