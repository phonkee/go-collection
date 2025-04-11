package collection

// Unique only returns collection of items that are unique
//
// This function gives ability to have unique items in collection.
func Unique[T any, U comparable](c Collection[T], fn func(p T) U, opts ...*UniqueOptions) Collection[T] {
	var o *UniqueOptions
	if len(opts) > 0 && opts[0] != nil {
		o = opts[0]
	} else {
		o = defaultUniqueOptions()
	}
	_ = o
	// result collection
	result := make(Collection[T], 0)
	// set of seen items (by key)
	seen := make(map[U]struct{})
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

// defaultUniqueOptions returns default options for Unique function if not provided
func defaultUniqueOptions() *UniqueOptions {
	return &UniqueOptions{}
}

// UniqueOptions is a struct that contains options for the Unique function
type UniqueOptions struct {
}
