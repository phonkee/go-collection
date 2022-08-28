package collection

import (
	"fmt"
	"strings"
)

const (
	PrependedDividerCount = 20
	SectionHeaderWidth    = 120
	DividerChar           = "="
)

// Section formats section to stdout, inner function should only use fmt to print anything to stdout
func Section(name string, in func()) {
	header := strings.Repeat(DividerChar, PrependedDividerCount) + " " + name + " " + strings.Repeat(DividerChar, SectionHeaderWidth)
	footer := strings.Repeat(DividerChar, PrependedDividerCount) + "/" + name + " " + strings.Repeat(DividerChar, SectionHeaderWidth)
	fmt.Println(header[:SectionHeaderWidth])
	in()
	fmt.Println(footer[:SectionHeaderWidth])
}
