package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	for _, data := range []struct {
		input    []int
		fn       func(p int) int
		expected []int
	}{
		{
			input:    []int{1, 1, 2, 2, 3, 3, 3, 1, 1, 1, 1, 2, 2, 2, 2},
			fn:       func(p int) int { return p * 2 },
			expected: []int{1, 2, 3},
		},
	} {
		found := Unique(Collection[int](data.input), data.fn)
		exp := Collection[int](data.expected)
		assert.Equal(t, exp, found)
	}
}
