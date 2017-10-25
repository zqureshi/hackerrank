package main

import "fmt"

func main() {
	var count int
	fmt.Scanf("%d", &count)

	var score int
	fmt.Scanf("%d", &score)

	highest, highestCount := score, 0
	lowest, lowestCount := score, 0

	for i := 1; i < count; i++ {
		fmt.Scanf("%d", &score)

		if score > highest {
			highest = score
			highestCount++
		} else if score < lowest {
			lowest = score
			lowestCount++
		}
	}

	fmt.Printf("%d %d\n", highestCount, lowestCount)
}
