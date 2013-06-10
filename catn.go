package main

import (
	"bufio"
	"fmt"
	"os"
)

// Prints a file to the standard output numbering each line
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>\n", os.Args[0])
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	for sc, line := bufio.NewScanner(file), 1; sc.Scan(); line++ {
		fmt.Printf("%3d: %s\n", line, sc.Text())
	}
}
