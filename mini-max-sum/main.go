package main

import (
	"fmt"
	"sort"
)

func main() {
	var n [5]int
	fmt.Scanf("%d %d %d %d %d", &n[0], &n[1], &n[2], &n[3], &n[4])

	sort.Ints(n[:])

	fmt.Println(n[0]+n[1]+n[2]+n[3], n[1]+n[2]+n[3]+n[4])
}
