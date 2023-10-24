// Package for replicating Unix command `wc`
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	count := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lines := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
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

		number_of_lines := 0

		for scanner.Scan() {
			number_of_lines++
		}

		fmt.Println(number_of_lines)
	}
}
