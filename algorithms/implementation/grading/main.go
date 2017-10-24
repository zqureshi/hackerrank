package main

import "fmt"

func main() {
	var count int
	fmt.Scanf("%d", &count)

	var grade int
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &grade)
		fmt.Println(round(grade))
	}
}

func round(grade int) int {
	if grade < 38 {
		return grade
	}

	if grade%5 > 2 {
		return grade + (5 - (grade % 5))
	}

	return grade
}
