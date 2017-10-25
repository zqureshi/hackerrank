package main

import "fmt"

func main() {
	var count int
	fmt.Scanf("%d", &count)

	squares := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &squares[i])
	}

	var day, month int
	fmt.Scanf("%d %d", &day, &month)

	sum := 0
	for _, n := range squares[:month] {
		sum += n
	}

	totalWays := 0
	if sum == day {
		totalWays++
	}

	for i, n := range squares[month:] {
		sum = sum - squares[i] + n

		if sum == day {
			totalWays++
		}
	}

	fmt.Println(totalWays)
}
