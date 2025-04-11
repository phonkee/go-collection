package collection

// Unique only returns collection of items that are unique
//
// This function gives ability to have unique items in a collection.
func Unique[T any, K comparable](c Collection[T], fn func(p T) K) Collection[T] {
	// result collection
	result := make(Collection[T], 0)
	// set of seen items (by key)
	seen := make(map[K]struct{})
	// iterate over the collection and check if an item key is unique
	c.Each(func(t T) {
		// get the key for the item
		key := fn(t)

		// check if the key is already in the seen set
		// if not store the item in the result collection and store the key in the seen set (for perf reasons)
		if _, ok := seen[key]; !ok {
			result = append(result, t)
			seen[key] = struct{}{}
		}
	})
	return result
}
