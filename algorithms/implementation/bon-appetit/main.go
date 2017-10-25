package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	var item int
	totalAnna := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &item)

		if k != i {
			totalAnna += item
		}
	}
	// Always guaranteed to be integer
	totalAnna /= 2

	var charged int
	fmt.Scanf("%d", &charged)

	if totalAnna == charged {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged - totalAnna)
	}
}
