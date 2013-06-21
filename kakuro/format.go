package kakuro

import (
	"fmt"
	"github.com/sbalev/gobananas/bitset"
)

// String

func (cell *ClueCell) String() string {
	return clueString(cell.VClue) + "\\" + clueString(cell.HClue)
}

func (cell *ValueCell) String() string {
	if cell.Val.Size() == 1 {
		return fmt.Sprintf("  %d  ", cell.Val.Min())
	}
	return "     "
}

func (g *Grid) String() string {
	s := ""
	for _, r := range g.cells {
		s += "|"
		for _, c := range r {
			s += c.String() + "|"
		}
		s += "\n"
	}
	return s
}

// Pretty String

func (g *Grid) PrettyString() string {
	s := "  |"
	sep := "--+"
	for j := 0; j < g.w; j++ {
		sep += "---+"
		s += fmt.Sprintf("%3d|", j)
	}
	sep += "\n"
	s += "\n" + sep
	for i, r := range g.cells {
		for k := 0; k < 3; k++ {
			if k == 1 {
				s += fmt.Sprintf("%2d|", i)
			} else {
				s += "  |"
			}
			for _, c := range r {
				s += c.prettyString(k) + "|"
			}
			s += "\n"
		}
		s += sep
	}
	return s
}

// helpers
func clueString(clue uint8) string {
	if clue == NoClue {
		return "##"
	}
	return fmt.Sprintf("%02d", clue)
}

func (cell *ClueCell) prettyString(k int) string {
	var s string
	switch k {
	case 0:
		s = "\\" + clueString(cell.HClue)
	case 1:
		if cell.VClue == NoClue {
			s = "#"
		} else {
			s = " "
		}
		s += "\\"
		if cell.HClue == NoClue {
			s += "#"
		} else {
			s += " "
		}
	case 2:
		s = clueString(cell.VClue) + "\\"
	}
	return s
}

func (cell *ValueCell) prettyString(k int) string {
	s := ""
	for e := 3*k + 1; e <= 3*k+3; e++ {
		if cell.Val.Contains(bitset.Element(e)) {
			s += fmt.Sprintf("%d", e)
		} else {
			s += " "
		}
	}
	return s
}
