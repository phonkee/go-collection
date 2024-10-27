package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		for _, data := range []struct {
			input     []int
			splitFunc func(p int) int
			expected  map[int][]int
		}{
			{
				input:     []int{1},
				splitFunc: func(p int) int { return p },
				expected:  map[int][]int{1: {1}},
			},
			{
				input:     []int{1, 1, 2, 2, 3, 3, 3},
				splitFunc: func(p int) int { return p },
				expected: map[int][]int{
					1: {1, 1},
					2: {2, 2},
					3: {3, 3, 3},
				},
			},
			{
				input:     []int{1, 2, 1, 2, 3, 2, 3},
				splitFunc: func(p int) int { return p },
				expected: map[int][]int{
					1: {1, 1},
					2: {2, 2, 2},
					3: {3, 3},
				},
			},
		} {
			exp := make(map[int]Collection[int])
			for k, v := range data.expected {
				exp[k] = v
			}
			assert.Equal(t, exp, Split(data.input, data.splitFunc))
		}
	})
}

func TestSplitMut(t *testing.T) {

	t.Run("Test unsorted", func(t *testing.T) {
		for _, item := range []struct {
			input     []int
			splitFunc func(p int) int
			expected  map[int][]int
		}{
			{
				input:     []int{1},
				splitFunc: func(p int) int { return p },
				expected:  map[int][]int{1: {1}},
			},
			{
				input:     []int{1, 1, 2, 2, 3, 3, 3},
				splitFunc: func(p int) int { return p },
				expected: map[int][]int{
					1: {1, 1},
					2: {2, 2},
					3: {3, 3, 3},
				},
			},
			{
				input:     []int{1, 2, 1, 2, 3, 2, 3},
				splitFunc: func(p int) int { return p },
				expected: map[int][]int{
					1: {1, 1},
					2: {2, 2, 2},
					3: {3, 3},
				},
			},
		} {
			exp := make(map[int]Collection[int])
			for k, v := range item.expected {
				exp[k] = v
			}

			assert.Equal(t, exp, SplitMut(item.input, item.splitFunc, false))
		}
	})

	t.Run("Test sorted", func(t *testing.T) {
		for _, item := range []struct {
			input     []int
			splitFunc func(p int) int
			expected  map[int][]int
		}{
			{
				input:     []int{1},
				splitFunc: func(p int) int { return p },
				expected:  map[int][]int{1: {1}},
			},
			{
				input:     []int{1, 1, 2, 2, 3, 3, 3},
				splitFunc: func(p int) int { return p },
				expected: map[int][]int{
					1: {1, 1},
					2: {2, 2},
					3: {3, 3, 3},
				},
			},
		} {
			exp := make(map[int]Collection[int])
			for k, v := range item.expected {
				exp[k] = v
			}

			assert.Equal(t, exp, SplitMut(item.input, item.splitFunc, true))
		}

	})

}
