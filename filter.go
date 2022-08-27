package collection

// FilterFunc is a definition of filter function to be passed to Filter method
type FilterFunc[T any] func(T) bool

// FilterAll evaluates filters, and if first returns false, it returns false
func FilterAll[T any](filters ...FilterFunc[T]) FilterFunc[T] {
	return func(i T) bool {
		for _, f := range filters {
			if !f(i) {
				return false
			}
		}
		return true
	}
}

// FilterAny evaluates filters, and if first returns true, it returns true
func FilterAny[T any](filters ...FilterFunc[T]) FilterFunc[T] {
	return func(i T) bool {
		for _, f := range filters {
			if f(i) {
				return true
			}
		}
		return false
	}
}

// FilterNone evaluates filters, and if first returns true, it returns false
func FilterNone[T any](filters ...FilterFunc[T]) FilterFunc[T] {
	return func(i T) bool {
		for _, f := range filters {
			if f(i) {
				return false
			}
		}
		return true
	}
}
