package collection

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Collection describes a collection of elements.
type Collection[T any] []T

// Chain other collections and return new combined collections
func (c Collection[T]) Chain(others ...Collection[T]) Collection[T] {
	size := len(c)
	for _, other := range others {
		size += len(other)
	}
	result := make(Collection[T], 0, size)
	result = append(result, c...)
	for _, other := range others {
		result = append(result, other...)
	}
	return result
}

// Copy current collection
func (c Collection[T]) Copy() Collection[T] {
	result := make(Collection[T], len(c))
	copy(result, c)
	return result
}

// Cycle returns collection by cycling over collection until the result size is reached.
func (c Collection[T]) Cycle(size int) Collection[T] {
	if c.Len() == 0 {
		return Collection[T]{}
	}
	result := make(Collection[T], 0, size)
	for {
		for _, item := range c {
			if result.Len() == size {
				return result
			}
			result = append(result, item)
		}
	}
}

// Debug prints data to stdout, if formatFuncs are not provided, item is printed as whole
func (c Collection[T]) Debug(title string, formatFuncs ...func(T) string) Collection[T] {
	Section(fmt.Sprintf("Debug: %v", title), func() {
		c.Enumerate(func(index int, item T) {
			var repr string
			if len(formatFuncs) > 0 {
				data := make([]string, 0)
				for _, ff := range formatFuncs {
					if what := strings.TrimSpace(ff(item)); what != "" {
						data = append(data, what)
					}
				}
				repr = strings.Join(data, ", ")
			} else {
				repr = fmt.Sprintf("%#v", item)
			}
			fmt.Printf("Item[%v] = %v\n", index, repr)
		})
	})
	return c
}

// Each calls given function for each item in collection
func (c Collection[T]) Each(fn func(T)) Collection[T] {
	return c.Enumerate(func(_ int, t T) {
		fn(t)
	})
}

// Enumerate calls given closure with each element and index
func (c Collection[T]) Enumerate(fn func(int, T)) Collection[T] {
	for index, item := range c {
		fn(index, item)
	}
	return c
}

// Filter Collection items  by given filter function
func (c Collection[T]) Filter(fn FilterFunc[T]) Collection[T] {
	result := make(Collection[T], 0, c.Len())
	c.Each(func(t T) {
		if fn(t) {
			result = append(result, t)
		}
	})
	return result
}

// First calls method on first element in the list, if not applied it returns false
func (c Collection[T]) First(fn func(p T)) bool {
	if len(c) == 0 {
		return false
	}
	fn(c[0])
	return true
}

// Last calls method on last element in the list, if no elements returns false
func (c Collection[T]) Last(fn func(p T)) bool {
	if len(c) == 0 {
		return false
	}
	fn(c[len(c)-1])
	return true
}

// Len returns length of the list
func (c Collection[T]) Len() int {
	return len(c)
}

// Map calls given function for each element in the list and returns new element
func (c Collection[T]) Map(mapFunc func(T) T) Collection[T] {
	result := make(Collection[T], len(c))
	c.Enumerate(func(index int, t T) {
		result[index] = mapFunc(t)
	})
	return result
}

// Shuffle shuffles the collection
func (c Collection[T]) Shuffle() Collection[T] {
	// copy collection first
	result := c.Copy()

	// do the work
	rand.Shuffle(result.Len(), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

// Sort collection and return it
func (c Collection[T]) Sort(sortFunc func(t1, t2 T) bool) Collection[T] {
	result := c.Copy()

	sort.Slice(result, func(i, j int) bool {
		return sortFunc(result[i], result[j])
	})

	return result
}

// Sub is sub routine that will be run with current collection
func (c Collection[T]) Sub(fn func(Collection[T])) Collection[T] {
	fn(c)
	return c
}

// Take takes only given amount of elements from the list, it can be less
func (c Collection[T]) Take(n int) Collection[T] {
	result := make(Collection[T], 0, n)
	result = append(result, c[:n]...)
	return result
}

// TakeUntil is called until the given function returns false.
func (c Collection[T]) TakeUntil(fn func(p T) bool) Collection[T] {
	result := make(Collection[T], 0)
	for _, t := range c {
		if fn(t) {
			result = append(result, t)
			continue
		}
		break
	}
	return result
}

// Unique only returns collection of items that are unique
func (c Collection[T]) Unique(fn func(p T) string) Collection[T] {
	result := make(Collection[T], 0, len(c))
	seen := make(map[string]struct{})
	c.Each(func(t T) {
		key := fn(t)
		if _, ok := seen[key]; !ok {
			result = append(result, t)
			seen[key] = struct{}{}
		}
	})
	return result
}
