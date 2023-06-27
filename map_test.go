package collection

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	in := []string{"1", "2", "3"}
	expect := []int{1, 2, 3}

	result := Map(in, func(s string) int {
		val, _ := strconv.ParseInt(s, 10, 64)
		return int(val)
	}).Into()

	assert.Equal(t, expect, result)
}
