package main

import (
	"fmt"
)

func main() {
	var order int
	fmt.Scanf("%d", &order)

	var matrix [100][100]int
	for i := 0; i < order; i++ {
		for j := 0; j < order; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	sum := 0
	for i := 0; i < order; i++ {
		sum += matrix[i][i] - matrix[i][order-i-1]
	}

	if sum < 0 {
		sum = -sum
	}

	fmt.Println(sum)
}
