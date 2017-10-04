package main

import "fmt"

func main() {
	alice_triplets := read_triplets()
	bob_triplets := read_triplets()

	var alice, bob = 0, 0
	for i := 0; i < 3; i++ {
		if alice_triplets[i] > bob_triplets[i] {
			alice++
		} else if alice_triplets[i] < bob_triplets[i] {
			bob++
		}
	}

	fmt.Println(alice, bob)
}

func read_triplets() [3]int {
	var t [3]int
	fmt.Scanf("%d %d %d", &t[0], &t[1], &t[2])

	return t
}