package main

import (
	"fmt"
	"strings"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)

	for i := 1; i <= length; i++ {
		fmt.Println(strings.Repeat(" ", length-i) + strings.Repeat("#", i))
	}
}
