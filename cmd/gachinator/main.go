package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/stefanitsky/gachinator"
)

var (
	currentLine    []byte
	gachinatedLine []byte
)

func main() {
	lang := flag.String("lang", "ru", "select language config")
	flag.Parse()

	lc, err := gachinator.FindLangConfig(*lang)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		currentLine = sc.Bytes()
		gachinatedLine = gachinator.Gachinate(currentLine, *lc)
		fmt.Fprint(os.Stdout, string(gachinatedLine))
	}

	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}
}
