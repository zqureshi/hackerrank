package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)

	candles := make([]int, length)
	max := 0
	for i := range candles {
		fmt.Scanf("%d", &candles[i])
		if candles[i] > max {
			max = candles[i]
		}
	}

	countMax := 0
	for _, n := range candles {
		if n == max {
			countMax++
		}
	}

	fmt.Println(countMax)
}
