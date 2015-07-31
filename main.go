package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var columnWidth uint64 // How long each line should be

var line string

func handle(source string) {
	var words = strings.Split(regexp.MustCompile("\n").ReplaceAllString(source, ""), " ")

	for _, word := range words {
		if length := len(line + word + " "); length > int(columnWidth) {
			fmt.Println(line)
			line = word + " "
		} else if length == int(columnWidth) {
			fmt.Println(line + word)
			line = ""
		} else {
			line += word + " "
		}
	}

	if source == "\n" {
		fmt.Println(strings.TrimSuffix(line, " "))
		line = ""
	}
}

func main() {
	columnWidth, _ = strconv.ParseUint(os.Args[1], 10, 8)

	input := bufio.NewReader(os.Stdin)

	for line, err := input.ReadString('\n'); err == nil; line, err = input.ReadString('\n') {
		handle(line)
	}
}
