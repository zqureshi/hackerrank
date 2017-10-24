package main

import "fmt"

func main() {
	var k1_start, k1_speed int
	fmt.Scanf("%d %d", &k1_start, &k1_speed)

	var k2_start, k2_speed int
	fmt.Scanf("%d %d", &k2_start, &k2_speed)

	// Assuming k2 is always on the right, so swap places if that isn't true
	if k1_start > k2_start {
		k1_start, k1_speed, k2_start, k2_speed = k2_start, k2_speed, k1_start, k1_speed
	}

	if k1_start == k2_start {
		fmt.Println("YES")
		return
	}

	if k2_speed >= k1_speed {
		fmt.Println("NO")
		return
	}

	for k1_start < k2_start {
		k1_start += k1_speed
		k2_start += k2_speed
	}

	if k1_start == k2_start {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
