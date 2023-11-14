// Package for replicating Unix command `wc`
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	count := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lines := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	words := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	characters := flag.Bool("m", false, "The number of characters in each input file is written to the standard output.")
	flag.Parse()

	filename := flag.Arg(0)

	input, err := getInput(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *count {
		bytes := countBytes(input)

		fmt.Println(bytes, filename)
	} else if *lines {
		lineCount := countLines(input)

		fmt.Println(lineCount, filename)
	} else if *words {
		words := countWords(input)

		fmt.Println(words, filename)
	} else if *characters {
		runes := countRune(input)

		fmt.Println(runes, filename)
	} else {
		wordCount := countWords(input)
		lineCount := countLines(input)
		bytes := countRune(input)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(lineCount, wordCount, bytes, filename)
	}
}

func countBytes(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	numberOfLines := 0

	for scanner.Scan() {
		numberOfLines++
	}

	return
}

func countRune(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	count = 0

	for scanner.Scan() {
		count++
	}

	return
}

func countLines(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)

	count = 0
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		count++
	}

	return
}

func countWords(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)

	numberOfWords := 0

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}

func getInput(filename string) (fileContent io.Reader, err error) {
	input := os.Stdin
	fi, err := input.Stat()

	if err != nil {
		fmt.Println("file.Stat()", err)
	}

	size := fi.Size()

	if size > 0 {
		return input, nil
	} else {
		contents, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		return contents, nil
	}
}
