package kakuro

import (
	"github.com/sbalev/gobananas/bitset"
	"io"
)

type gridErr string

func (err gridErr) Error() string {
	return string(err)
}

type Grid struct {
	h, w  int
	cells [][]Cell
}

func (g *Grid) H() int {
	return g.h
}

func (g *Grid) W() int {
	return g.w
}

func (g *Grid) Cell(i, j int) Cell {
	return g.cells[i][j]
}

func (g *Grid) SetCell(i, j int, c Cell) {
	g.cells[i][j] = c
}

func NewGrid(h, w int) *Grid {
	if h < 2 || w < 2 {
		panic("A grid must be at least 2x2")
	}
	g := &Grid{cells: make([][]Cell, h), h: h, w: w}
	g.cells[0] = make([]Cell, w)
	for j := range g.cells[0] {
		g.cells[0][j] = &ClueCell{}
	}
	for i := 1; i < h; i++ {
		g.cells[i] = make([]Cell, w)
		g.cells[i][0] = &ClueCell{}
		for j := 1; j < w; j++ {
			g.cells[i][j] = &ValueCell{Val: bitset.Interval(1, 9)}
		}
	}
	return g
}

func ReadGrid(rd io.Reader) (*Grid, error) {
	cells, err := parseGrid(rd)
	if err != nil {
		return nil, err
	}
	if len(cells) < 2 {
		return nil, gridErr("grid too short")
	}
	return &Grid{h: len(cells), w: len(cells[0]), cells: cells}, nil
}
