package collection

import "testing"

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
