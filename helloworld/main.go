package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Hello, World.")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fmt.Println(string(line))
}
