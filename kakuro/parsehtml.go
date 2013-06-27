package kakuro

import (
	"code.google.com/p/go.net/html"
	"github.com/sbalev/gobananas/bitset"
	"github.com/sbalev/gobananas/htmlutil"
	"net/http"
	"strconv"
	"strings"
)

func FetchGrid(url string) (*Grid, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	n, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	for _, tag := range []string{"html", "body", "table"} {
		if c := htmlutil.GetElChildren(n, tag); len(c) == 0 {
			return nil, gridErr(tag + " not found")
		} else {
			n = c[0]
		}
	}
	if cl, _ := htmlutil.GetAttr(n, "class"); cl != "kakurotablegrey" {
		return nil, gridErr("Wrong table class")
	}
	tab, err := htmlutil.ParseTable(n)
	if len(tab) < 2 {
		return nil, gridErr("Grid too short")
	}
	colCount := len(tab[0])
	if colCount < 2 {
		return nil, gridErr("Grid too narrow")
	}
	for _, r := range tab {
		if len(r) != colCount {
			return nil, gridErr("Rows of different lengths")
		}
	}
	g := &Grid{h: len(tab), w: colCount, cells: make([][]Cell, len(tab))}
	for i, r := range tab {
		g.cells[i] = make([]Cell, g.w)
		for j, c := range r {
			if x, err := parseHTMLCell(c); err != nil {
				return nil, err
			} else {
				g.cells[i][j] = x
			}
		}
	}
	return g, nil
}

func parseHTMLCell(n *html.Node) (Cell, error) {
	switch v, _ := htmlutil.GetAttr(n, "class"); v {
	case "whitegrey":
		return &ValueCell{Val: bitset.Interval(1, 9)}, nil
	case "grey":
		return &ClueCell{}, nil
	case "infocellgrey":
		if v, h, err := getClues(n); err != nil {
			return nil, err
		} else {
			return &ClueCell{VClue: v, HClue: h}, nil
		}
	default:
		return nil, gridErr("Unknown or missing cell class")
	}
}

func getClues(n *html.Node) (uint8, uint8, error) {
	if c := htmlutil.GetElChildren(n, "table"); len(c) == 0 {
		return 0, 0, gridErr("table not found")
	} else {
		n = c[0]
	}
	if cl, _ := htmlutil.GetAttr(n, "class"); cl != "infotablegrey" {
		return 0, 0, gridErr("Wrong table class")
	}
	tab, err := htmlutil.ParseTable(n)
	if err != nil {
		return 0, 0, err
	}
	if len(tab) != 2 || len(tab[0]) != 2 || len(tab[1]) != 2 {
		return 0, 0, gridErr("Wrong table size")
	}
	v, err := getClue(tab[1][0])
	if err != nil {
		return 0, 0, err
	}
	h, err := getClue(tab[0][1])
	if err != nil {
		return 0, 0, nil
	}
	return v, h, nil
}

func getClue(n *html.Node) (uint8, error) {
	n = n.FirstChild
	if n == nil {
		return NoClue, nil
	}
	if n.Type != html.TextNode {
		return 0, gridErr("Clue is not text element")
	}
	if c, err := strconv.Atoi(strings.TrimSpace(n.Data)); err != nil {
		return NoClue, nil
	} else {
		return uint8(c), nil
	}
}
