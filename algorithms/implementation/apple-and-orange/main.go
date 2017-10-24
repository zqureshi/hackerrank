package main

import "fmt"

func main() {
	var house_start, house_end int
	fmt.Scanf("%d %d", &house_start, &house_end)

	var apple_tree, orange_tree int
	fmt.Scanf("%d %d", &apple_tree, &orange_tree)

	var count_apples, count_oranges int
	fmt.Scanf("%d %d", &count_apples, &count_oranges)

	var apple_location int
	apples_on_house := 0
	for i := 0; i < count_apples; i++ {
		fmt.Scanf("%d", &apple_location)

		apple_location += apple_tree
		if apple_location >= house_start && apple_location <= house_end {
			apples_on_house++
		}
	}

	var orange_location int
	oranges_on_house := 0
	for i := 0; i < count_oranges; i++ {
		fmt.Scanf("%d", &orange_location)

		orange_location += orange_tree
		if orange_location >= house_start && orange_location <= house_end {
			oranges_on_house++
		}
	}

	fmt.Println(apples_on_house)
	fmt.Println(oranges_on_house)
}
