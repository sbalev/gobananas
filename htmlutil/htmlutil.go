package htmlutil

//This package contains several useful functions completing the
//go.net/html package

import "code.google.com/p/go.net/html"

type parseErr string

func (err parseErr) Error() string {
	return string(err)
}

// IsEl returns true if n is of type ElementNode and
// it's Data field is equal to data parameter
func IsEl(n *html.Node, data string) bool {
	return n != nil && n.Type == html.ElementNode && n.Data == data
}

// GetElChildren returns a slice containing the children of n
// of type ElementNode and with Data field equal to data parameter
func GetElChildren(n *html.Node, data string) []*html.Node {
	if n == nil {
		return nil
	}
	i := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsEl(c, data) {
			i++
		}
	}
	children := make([]*html.Node, i)
	i = 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsEl(c, data) {
			children[i] = c
			i++
		}
	}
	return children
}

// ParseTable returns a two-dimensional slice containing
// the "td" children of a table node
func ParseTable(n *html.Node) ([][]*html.Node, error) {
	if !IsEl(n, "table") {
		return nil, parseErr("Not a table")
	}
	if b := GetElChildren(n, "tbody"); len(b) != 1 {
		return nil, parseErr("None or multiple tbody")
	} else {
		n = b[0]
	}
	rows := GetElChildren(n, "tr")
	tab := make([][]*html.Node, len(rows))
	for i, row := range rows {
		tab[i] = GetElChildren(row, "td")
	}
	return tab, nil
}
