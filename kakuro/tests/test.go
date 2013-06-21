package main

import (
	"fmt"
	"github.com/sbalev/gobananas/kakuro"
	"os"
)

func main() {
	g := kakuro.NewGrid(6, 6)

	g.Cell(0, 2).(*kakuro.ClueCell).VClue = 30
	g.Cell(0, 3).(*kakuro.ClueCell).VClue = 16
	g.Cell(2, 0).(*kakuro.ClueCell).HClue = 25
	g.Cell(3, 0).(*kakuro.ClueCell).HClue = 11

	g.SetCell(1, 1, &kakuro.ClueCell{4, 13})
	g.SetCell(1, 4, &kakuro.ClueCell{30, kakuro.NoClue})
	g.SetCell(1, 5, &kakuro.ClueCell{})

	g.SetCell(2, 5, &kakuro.ClueCell{3, kakuro.NoClue})

	g.SetCell(3, 3, &kakuro.ClueCell{4, 10})

	g.SetCell(4, 1, &kakuro.ClueCell{kakuro.NoClue, 20})

	g.SetCell(5, 1, &kakuro.ClueCell{})
	g.SetCell(5, 2, &kakuro.ClueCell{kakuro.NoClue, 8})
	g.SetCell(5, 5, &kakuro.ClueCell{})

	fmt.Println(g.PrettyString())

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	g2, err := kakuro.ReadGrid(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(g)
	fmt.Println(g2)
}
