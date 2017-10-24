package main

import "fmt"

func main() {
	var length, num int
	fmt.Scanf("%d", &length)

	plus, minus, zero := 0, 0, 0
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &num)

		if num < 0 {
			minus++
		} else if num > 0 {
			plus++
		} else {
			zero++
		}
	}

	fmt.Println(safeDivide(plus, length))
	fmt.Println(safeDivide(minus, length))
	fmt.Println(safeDivide(zero, length))
}

func safeDivide(num int, den int) float64 {
	if den == 0 {
		return 0
	} else {
		return float64(num) / float64(den)
	}
}
