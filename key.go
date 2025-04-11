package collection

import "golang.org/x/exp/constraints"

// KeyFunc is a function that extracts key from a given item.
type KeyFunc[T any, K constraints.Ordered] func(T) K
