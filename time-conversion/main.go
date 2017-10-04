package main

import (
	"fmt"
	"strings"
)

func main() {
	var hh, mm, ss int
	var period string
	fmt.Scanf("%d:%d:%d%s", &hh, &mm, &ss, &period)

	period = strings.ToLower(period)

	if period == "pm" && hh != 12 {
		hh += 12
	} else if period == "am" && hh == 12 {
		hh = 0
	}

	fmt.Printf("%02d:%02d:%02d\n", hh, mm, ss)
}
