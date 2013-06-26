package kakuro

import (
	"github.com/sbalev/gobananas/bitset"
)

type blockCell struct {
	id                cellId
	possible, checked bitset.BitSet
}

type block struct {
	clue  uint8
	cells map[cellId]*blockCell
	cleft []*blockCell
}

func (b *block) findFeas(i, sum int, used bitset.BitSet) bool {
	if i == len(b.cleft) {
		return sum == 0
	}
	if sum <= 0 {
		return false
	}
	for _, e := range b.cleft[i].possible.Minus(used).Range() {
		if b.findFeas(i+1, sum-int(e), used.Add(e)) {
			b.cleft[i].checked = b.cleft[i].checked.Add(e)
			return true
		}
	}
	return false
}

func (b *block) check(cid cellId, v bitset.Element) bool {
	sum := int(b.clue) - int(v)
	used := bitset.Set(v)
	b.cleft = b.cleft[:0]
	for id, c := range b.cells {
		if id != cid {
			switch c.possible.Size() {
			case 0:
				return false
			case 1:
				e := c.possible.Min()
				if used.Contains(e) {
					return false
				}
				sum -= int(e)
				used = used.Add(e)
			default:
				b.cleft = append(b.cleft, c)
			}
		}
	}
	return b.findFeas(0, sum, used)
}

func (b *block) checkAll() {
	for _, c := range b.cells {
		c.checked = bitset.EmptySet
	}
	for id, c := range b.cells {
		for e := bitset.Element(1); e <= 9; e++ {
			// read from input channel here
			if c.possible.Minus(c.checked).Contains(e) {
				if b.check(id, e) {
					c.checked = c.checked.Add(e)
				} else {
					c.possible = c.possible.Remove(e)
					// write to output channel here
				}
			}
		}
	}
}
