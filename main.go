package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var columnWidth int // How long each line should be

func handle(line string) {
	// Implement this
	fmt.Println(line)
}

func main() {
	width, _ := strconv.ParseInt(os.Args[1], 10, 8)
	columnWidth = int(width)

	input := bufio.NewReader(os.Stdin)

	for line, err := input.ReadString('\n'); err == nil; line, err = input.ReadString('\n') {
		handle(line)
	}
}
