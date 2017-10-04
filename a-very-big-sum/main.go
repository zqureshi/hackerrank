package main

import "fmt"

func main() {
	var length, num int
	fmt.Scanf("%d", &length)

	sum := 0
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &num)
		sum += num
	}

	fmt.Println(sum)
}
