package collection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("Test FilterAny", func(t *testing.T) {
		data := []struct {
			ff       FilterFunc[int]
			expected bool
		}{
			{FilterAny[int](func(i int) bool { return false }), false},
			{FilterAny[int](func(i int) bool { return false }, func(i int) bool { return true }), true},
			{FilterAny[int](func(i int) bool { return true }), true},
		}

		for i, item := range data {
			result := item.ff(1)
			if result != item.expected {
				t.Errorf("[%v] Expected %v, got %v", i, item.expected, result)
			}
		}
	})

	t.Run("Test FilterAll", func(t *testing.T) {
		data := []struct {
			ff       FilterFunc[int]
			expected bool
		}{
			{FilterAll[int](func(i int) bool { return false }), false},
			{FilterAll[int](func(i int) bool { return false }, func(i int) bool { return true }), false},
			{FilterAll[int](func(i int) bool { return true }), true},
			{FilterAll[int](func(i int) bool { return true }, func(i int) bool { return true }), true},
		}

		for i, item := range data {
			result := item.ff(1)
			if result != item.expected {
				t.Errorf("[%v] Expected %v, got %v", i, item.expected, result)
			}
		}
	})

	t.Run("Test FilterNone", func(t *testing.T) {
		data := []struct {
			ff       FilterFunc[int]
			expected bool
		}{
			{FilterNone[int](func(i int) bool { return false }), true},
			{FilterNone[int](func(i int) bool { return false }, func(i int) bool { return true }), false},
			{FilterNone[int](func(i int) bool { return true }), false},
			{FilterNone[int](func(i int) bool { return false }, func(i int) bool { return false }), true},
		}

		for i, item := range data {
			result := item.ff(1)
			if result != item.expected {
				t.Errorf("[%v] Expected %v, got %v", i, item.expected, result)
			}
		}
	})
}

func TestFilterMut(t *testing.T) {
	t.Run("Test FilterMut - inplace change", func(t *testing.T) {
		for _, data := range []struct {
			in           Collection[int]
			deleteValues []int
			expect       Collection[int]
		}{
			{[]int{1, 2, 3}, []int{1}, []int{2, 3}},
			{[]int{1, 2, 3}, []int{6}, []int{1, 2, 3}},
			{[]int{1, 2, 3}, nil, []int{1, 2, 3}},
			{[]int{1, 2, 3}, []int{77, 99}, []int{1, 2, 3}},
			{[]int{1, 2, 3}, []int{2}, []int{1, 3}},
			{[]int{1, 2, 3, 4, 5, 6}, []int{3, 4, 6}, []int{1, 2, 5}},
		} {
			ptr := fmt.Sprintf("%p", data.in)
			nuData := FilterMut(data.in, func(i int) bool {
				for _, idx := range data.deleteValues {
					if idx == i {
						return false
					}
				}
				return true
			})
			assert.Equal(t, data.expect, nuData)
			ptrNu := fmt.Sprintf("%p", nuData)
			assert.Equal(t, ptr, ptrNu)
		}
	})

}
