package collection

// GroupBy elements of the collection by the result of the function
//
// Deprecated: use Partition instead
func GroupBy[T any, K comparable](collection Collection[T], fn func(T) K) map[K]Collection[T] {
	return Partition(collection, fn)
}

// Partition elements of the collection by the result of the function
func Partition[T any, K comparable](collection Collection[T], fn func(T) K) map[K]Collection[T] {
	// maintain the order and return
	grouped := make(map[K]Collection[T])
	for _, item := range collection {
		key := fn(item)
		if _, ok := grouped[key]; !ok {
			grouped[key] = make(Collection[T], 0)
		}
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}
