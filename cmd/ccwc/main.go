// Package for replicating Unix command `wc`
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	count := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lines := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	words := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	characters := flag.Bool("m", false, "The number of characters in each input file is written to the standard output.")
	flag.Parse()

	filename := flag.Arg(0)

	if *count {
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(len(contents))

	} else if *lines {
		contents, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		scanner := bufio.NewScanner(contents)

		numberOfLines := 0

		for scanner.Scan() {
			numberOfLines++
		}

		fmt.Println(numberOfLines)
	} else if *words {
		contents, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		scanner := bufio.NewScanner(contents)

		numberOfWords := 0

		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)
			numberOfWords += len(words)
		}

		fmt.Println(numberOfWords)
	} else if *characters {
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(utf8.RuneCount(contents))
	}
}
