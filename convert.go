package collection

// Convert from one collection to another
func Convert[T any, U any](collection Collection[T], fn func(T) U) Collection[U] {
	result := make(Collection[U], len(collection))
	for index, item := range collection {
		result[index] = fn(item)
	}
	return result
}
