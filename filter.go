package collection

// Filter Collection items  by given filter function
func (c Collection[T]) Filter(fn FilterFunc[T]) Collection[T] {
	var result Collection[T]
	c.Each(func(t T) {
		if !fn(t) {
			return
		}
		result = append(result, t)
	})
	return result
}

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
