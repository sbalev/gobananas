package kakuro

import (
	"fmt"
	"github.com/sbalev/gobananas/bitset"
)

const NoClue uint8 = 0

type Cell interface {
	fmt.Stringer
	prettyString(k int) string
}

type ClueCell struct {
	VClue, HClue uint8
}

type ValueCell struct {
	Val bitset.BitSet
}
