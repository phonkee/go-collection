package collection

import (
	"fmt"
	"strings"
)

const (
	PrependedDividerCount = 10
	SectionHeaderWidth    = 80
	DividerChar           = "="
)

// Section formats section to stdout, inner function should only use fmt to print anything to stdout
func Section(name string, in func()) {
	header := strings.Repeat(DividerChar, PrependedDividerCount) + " " + name + " " + strings.Repeat(DividerChar, SectionHeaderWidth-len(name)+1)
	footer := strings.Repeat(DividerChar, PrependedDividerCount) + "/" + name + " " + strings.Repeat(DividerChar, SectionHeaderWidth-len(name)+1)
	fmt.Println(header[:80])
	in()
	fmt.Println(footer[:80])
}
