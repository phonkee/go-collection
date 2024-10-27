package collection

// Unique only returns collection of items that are unique
//
// This function gives ability to have unique items in collection.
func Unique[T any, U comparable](c Collection[T], fn func(p T) U) Collection[T] {
	result := make(Collection[T], 0, len(c))
	seen := make(map[U]struct{})
	c.Each(func(t T) {
		key := fn(t)
		if _, ok := seen[key]; !ok {
			result = append(result, t)
			seen[key] = struct{}{}
		}
	})
	return result
}
