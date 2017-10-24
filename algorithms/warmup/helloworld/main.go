package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fmt.Println(string(line))
}
