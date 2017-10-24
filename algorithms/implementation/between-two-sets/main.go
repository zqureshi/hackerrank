package main

import (
	"fmt"
)

func main() {
	var size_a, size_b int
	fmt.Scanf("%d %d", &size_a, &size_b)

	a := make([]int, size_a)
	for i := 0; i < size_a; i++ {
		fmt.Scanf("%d", &a[i])
	}

	b := make([]int, size_b)
	for i := 0; i < size_b; i++ {
		fmt.Scanf("%d", &b[i])
	}

	min_b := b[0]
	for _, n := range b {
		if n < min_b {
			min_b = n
		}
	}

	count_between := 0
Outer:
	for i := 1; i <= min_b; i++ {
		for _, n := range b {
			if n%i != 0 {
				continue Outer
			}
		}

		for _, n := range a {
			if i%n != 0 {
				continue Outer
			}
		}

		count_between++
	}

	fmt.Println(count_between)
}
