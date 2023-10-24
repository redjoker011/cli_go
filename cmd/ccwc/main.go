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
		bytes, err := readFile(filename)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(len(bytes), filename)

	} else if *lines {
		lineCount, err := countLines(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(lineCount, filename)
	} else if *words {
		wordCount, err := countWords(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(wordCount, filename)
	} else if *characters {
		bytes, err := readFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(utf8.RuneCount(bytes), filename)
	} else {
		lineCount, err := countLines(filename)
		wordCount, err := countWords(filename)
		bytes, err := readFile(filename)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(lineCount, wordCount, len(bytes), filename)
	}
}

func readFile(filename string) (bytes []byte, err error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func countLines(filename string) (count int, err error) {
	contents, err := os.Open(filename)
	defer contents.Close()

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(contents)

	numberOfLines := 0

	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines, nil
}

func countWords(filename string) (count int, err error) {
	contents, err := os.Open(filename)
	defer contents.Close()
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(contents)

	numberOfWords := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		numberOfWords += len(words)
	}

	return numberOfWords, nil
}
