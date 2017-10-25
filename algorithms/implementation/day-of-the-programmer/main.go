package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Scanf("%d", &year)

	switch {
	case year < 1918:
		julianDay(year)
	case year > 1918:
		gregorianDay(year)
	default:
		fmt.Println("26.09.1918")
	}
}

func julianDay(year int) {
	if year%4 == 0 {
		fmt.Printf("12.09.%d", year)
	} else {
		fmt.Printf("13.09.%d", year)
	}
}

func gregorianDay(year int) {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("12.09.%d", year)
	} else {
		fmt.Printf("13.09.%d", year)
	}
}
