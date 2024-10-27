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

// FilterMut filters data inplace and returns a new collection from this altered data.
//
// Warning: The Original collection should not be used after this operation.
func FilterMut[T any](c Collection[T], fn FilterFunc[T]) Collection[T] {
	// indices is channel where we will push indices of items that should be removed
	indices := make(chan int)

	// in a separate goroutine we will iterate over a collection and send indices of items that should be removed
	go func() {
		// close the channel when we are done
		defer close(indices)

		// iterate over a collection and check if item should be removed by calling filter function
		for i, t := range c {
			if !fn(t) {
				indices <- i
			}
		}
	}()

	// the algorithm iterates over the channel of indices that need to be removed.
	// for each index received, it moves the pointer in the collection until it reaches the index.
	// it then increments the deleted count and continues.
	// at each step, items that are not to be deleted are moved left by the deleted count.
	// this approach ensures an almost O(n) time complexity.
	deleted := 0
	current := 0
	allCount := len(c)
	deletedItems := make([]T, 0)
outer:
	// iterate over indices to be removed
	for idx := range indices {
		// move items to the left and move on
		for current < idx {
			c[current-deleted] = c[current]
			current += 1
		}
		if current == idx {
			deletedItems = append(deletedItems, c[current])
			deleted += 1
			current += 1
			continue outer
		}
	}

	// finish moving the rest
	for current < allCount {
		c[current-deleted] = c[current]
		current += 1
	}

	// move deleted items to the end (even when we throw them out)
	for i, deletedItem := range deletedItems {
		c[allCount-deleted+i] = deletedItem
	}

	return c[:allCount-deleted]
}
