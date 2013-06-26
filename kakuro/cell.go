package kakuro

import (
	"fmt"
	"github.com/sbalev/gobananas/bitset"
)

type cellId uint8

const NoClue uint8 = 0

type Cell interface {
	fmt.Stringer
	prettyString(k int) string
}

type ClueCell struct {
	VClue, HClue uint8
}

type ValueCell struct {
	Val            bitset.BitSet
	id             cellId
	hBlock, vBlock *block
}

func (cell *ValueCell) blockCell() *blockCell {
	return &blockCell{
		id:       cell.id,
		possible: cell.Val,
	}
}
