package kakuro

import (
	"fmt"
	"github.com/sbalev/gobananas/bitset"
	"io"
)

type gridErr string

func (err gridErr) Error() string {
	return string(err)
}

type Grid struct {
	h, w   int
	cells  [][]Cell
	blocks []*block
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

func (g *Grid) assignCellIds() {
	var id cellId
	for _, r := range g.cells {
		for _, c := range r {
			if vc, ok := c.(*ValueCell); ok {
				vc.id = id
				id++
			}
		}
	}
}

func (g *Grid) makeHBlock(i, j int) *block {
	b := &block{
		clue:  g.cells[i][j].(*ClueCell).HClue,
		cells: make(map[cellId]*blockCell),
	}
	for _, c := range g.cells[i][j+1:] {
		if vc, ok := c.(*ValueCell); ok {
			b.cells[vc.id] = vc.blockCell()
			vc.hBlock = b
		} else {
			break
		}
	}
	b.cleft = make([]*blockCell, len(b.cells))
	return b
}

func (g *Grid) makeVBlock(i, j int) *block {
	b := &block{
		clue:  g.cells[i][j].(*ClueCell).VClue,
		cells: make(map[cellId]*blockCell),
	}
	for _, r := range g.cells[i+1:] {
		if vc, ok := r[j].(*ValueCell); ok {
			b.cells[vc.id] = vc.blockCell()
			vc.vBlock = b
		} else {
			break
		}
	}
	b.cleft = make([]*blockCell, len(b.cells))
	return b
}

func (g *Grid) makeBlocks() error {
	g.assignCellIds()
	g.blocks = make([]*block, 0)
	for i, r := range g.cells {
		for j, c := range r {
			if cc, ok := c.(*ClueCell); ok {
				if cc.HClue != 0 {
					g.blocks = append(g.blocks, g.makeHBlock(i, j))
				}
				if cc.VClue != 0 {
					g.blocks = append(g.blocks, g.makeVBlock(i, j))
				}
			}
		}
	}
	for i, r := range g.cells {
		for j, c := range r {
			if vc, ok := c.(*ValueCell); ok {
				if vc.hBlock == nil {
					return gridErr(fmt.Sprintf(
						"No horisontal clue for cell (%d, %d)", i, j))
				}
				if vc.vBlock == nil {
					return gridErr(fmt.Sprintf(
						"No vertical clue for cell (%d, %d)", i, j))
				}
			}
		}
	}
	return nil
}

func (g *Grid) solveOnce() bool {
	for _, b := range g.blocks {
		b.checkAll()
	}
	res := false
	for _, r := range g.cells {
		for _, c := range r {
			if vc, ok := c.(*ValueCell); ok {
				newVal := vc.hBlock.cells[vc.id].possible.Intersection(
					vc.vBlock.cells[vc.id].possible)
				if vc.Val != newVal {
					vc.Val = newVal
					vc.hBlock.cells[vc.id].possible = newVal
					vc.vBlock.cells[vc.id].possible = newVal
					res = true
				}
			}
		}
	}
	return res
}

func (g *Grid) Solve() error {
	err := g.makeBlocks()
	if err != nil {
		return err
	}
	for g.solveOnce() {
	}
	return nil
}
