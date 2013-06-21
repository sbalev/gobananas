package kakuro

import (
	"bufio"
	"fmt"
	"github.com/sbalev/gobananas/bitset"
	"io"
	"strconv"
	"strings"
)

const errClue uint8 = 100

type parseErr int

func (err parseErr) Error() string {
	return fmt.Sprintf("parse error in line %d", int(err))
}

func parseClue(s string) uint8 {
	if s == "##" {
		return NoClue
	}
	if clue, err := strconv.Atoi(s); err != nil || clue < 1 {
		return errClue
	} else {
		return uint8(clue)
	}
}

func parseCell(s string) Cell {
	if len(s) != 5 {
		return nil
	}
	if s == "     " {
		return &ValueCell{Val: bitset.Interval(1, 9)}
	}
	if s[2] == '\\' {
		if v, h := parseClue(s[:2]), parseClue(s[3:]); v == errClue || h == errClue {
			return nil
		} else {
			return &ClueCell{VClue: v, HClue: h}
		}
	}
	e, err := strconv.Atoi(s)
	if err != nil || e < 1 || 9 < e {
		return nil
	}
	return &ValueCell{Val: bitset.EmptySet.Add(bitset.Element(e))}
}

func parseRow(s string) []Cell {
	tokens := strings.Split(s, "|")
	l := len(tokens)
	if l < 4 || tokens[0] != "" || tokens[l-1] != "" {
		return nil
	}
	tokens, l = tokens[1:l-1], l-2
	row := make([]Cell, l)
	for j, t := range tokens {
		if row[j] = parseCell(t); row[j] == nil {
			return nil
		}
	}
	return row
}

func parseGrid(rd io.Reader) ([][]Cell, error) {
	cells := make([][]Cell, 0)
	colCount := 0
	for sc, lineNo := bufio.NewScanner(rd), parseErr(1); sc.Scan(); lineNo++ {
		line := sc.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		row := parseRow(line)
		if row == nil {
			return nil, lineNo
		}
		if colCount == 0 {
			colCount = len(row)
		}
		if len(row) != colCount {
			return nil, lineNo
		}
		cells = append(cells, row)
	}
	return cells, nil
}
