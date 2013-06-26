package main

import (
	"fmt"
	"github.com/sbalev/gobananas/kakuro"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	g, err := kakuro.ReadGrid(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = g.Solve()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(g)
}
