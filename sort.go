package collection

// SortBackward sorts collection in descending order.
func SortBackward[T any](sortFunc func(t1, t2 T) bool) func(t1, t2 T) bool {
	return func(t1, t2 T) bool {
		return !sortFunc(t1, t2)
	}
}
