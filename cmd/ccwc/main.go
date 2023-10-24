// Package for replicating Unix command `wc`
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	count := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	flag.Parse()

	filename := flag.Arg(0)

	if *count {
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(len(contents))
	} else {
		fmt.Println("Do nothing")
	}
}
