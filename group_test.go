package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupBy(t *testing.T) {

	type TestStruct struct {
		value int
	}

	c := New[TestStruct](
		TestStruct{value: 1},
		TestStruct{value: 1},
		TestStruct{value: 2},
		TestStruct{value: 2},
		TestStruct{value: 3},
		TestStruct{value: 3},
	)

	g := GroupBy(c, func(item TestStruct) int {
		return item.value
	})

	assert.Equal(t, 3, len(g))
	assert.Equal(t, 2, len(g[1]))
	assert.Equal(t, 2, len(g[2]))
	assert.Equal(t, 2, len(g[3]))
}
