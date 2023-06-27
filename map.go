package collection

// Map calls given function for each element in the list and returns collection of new elements
// Map that returns different type cannot be implemented as method, so it is implemented as function
func Map[T any, U any](collection Collection[T], fn func(T) U) Collection[U] {
	result := make(Collection[U], len(collection))
	for index, item := range collection {
		result[index] = fn(item)
	}
	return result
}
