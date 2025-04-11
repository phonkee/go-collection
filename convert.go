package collection

import "golang.org/x/exp/constraints"

// Convert from one collection to another
func Convert[T any, U constraints.Ordered](collection Collection[T], fn func(T) U) Collection[U] {
	result := make(Collection[U], len(collection))
	for index, item := range collection {
		result[index] = fn(item)
	}
	return result
}
