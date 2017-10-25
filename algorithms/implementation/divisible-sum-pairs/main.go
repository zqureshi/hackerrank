package main

import (
	"fmt"
)

func main() {
	var count, factor int
	fmt.Scanf("%d %d", &count, &factor)

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	pairs := 0
	for i, second := range numbers {
		for _, first := range numbers[i+1:] {
			if (first+second)%factor == 0 {
				pairs++
			}
		}
	}

	fmt.Println(pairs)
}
