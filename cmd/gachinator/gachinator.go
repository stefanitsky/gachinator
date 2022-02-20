package main

import (
	"bufio"
	"fmt"
	"os"

	"stefanitsky.me/gachinator"
)

var (
	currentLine    []byte
	gachinatedLine []byte
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		currentLine = sc.Bytes()
		gachinatedLine = gachinator.Gachinate(currentLine)
		fmt.Fprint(os.Stdout, string(gachinatedLine))
	}

	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}
}
