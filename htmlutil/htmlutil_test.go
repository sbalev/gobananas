package htmlutil

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"strconv"
	"strings"
)

func ExampleParseTable() {
	s := `
<table>
  <tr><td>11</td><td>12</td></tr>
  <tr><td>21</td><td>22</td></tr>
</table>
`
	n, _ := html.Parse(strings.NewReader(s))
	n = GetElChildren(n, "html")[0]
	n = GetElChildren(n, "body")[0]
	n = GetElChildren(n, "table")[0]
	tab, _ := ParseTable(n)
	itab := make([][]int, len(tab))
	for i, r := range tab {
		itab[i] = make([]int, len(r))
		for j, c := range r {
			itab[i][j], _ = strconv.Atoi(c.FirstChild.Data)
		}
	}
	fmt.Println(itab)
	// Output: [[11 12] [21 22]]
}
