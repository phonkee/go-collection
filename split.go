package collection

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Split splits a collection by given function.
// This function creates new map and leaves the original collection untouched.
func Split[T any, U constraints.Ordered](c Collection[T], fn func(T) U) map[U]Collection[T] {
	result := make(map[U]Collection[T])

	// iterate and split
	for _, item := range c {
		key := fn(item)
		value, ok := result[key]
		if !ok {
			result[key] = []T{item}
		} else {
			result[key] = append(value, item)
		}
	}

	return result
}

// SplitMut splits a collection by given function.
//
// This function will return a map where key is a result of given function and value is a slice of elements that have the same key.
// It automatically sorts the original collection by given function, or if you provide a sorted flag, it will skip sorting.
//
// Warning! This function will mutate the collection and returns a map. You will lose the original collection.
// Warning! When you set sorted to true, you must be sure that collection is sorted by given function, otherwise the result will be incorrect.
func SplitMut[T any, U constraints.Ordered](c Collection[T], fn KeyFunc[T, U], sorted ...bool) map[U]Collection[T] {
	var isSorted bool
	if len(sorted) > 0 {
		isSorted = sorted[0]
	}
	// if the collection is not sorted, we need to do it now
	// Warning! sort alters the original collection (that's why this function has Mut in its name)
	if !isSorted {
		sort.Slice(c, func(i, j int) bool {
			return fn(c[i]) < fn(c[j])
		})
	}

	// now we need to go over the collection and collect ranges of the same key
	// we are able to do this because the collection is sorted
	ranges := make(map[U]splitMutRange)
	for i, item := range c {
		// get the key for given item
		key := fn(item)

		// get range for a given key, if it does not exist, create a new one
		val, ok := ranges[key]
		if !ok {
			// create a new range with start at current index and length 1
			val = splitMutRange{i, 1}
		} else {
			// increase length of the range
			val.length++
		}

		// store the range back to the map
		ranges[key] = val
	}

	// now we need to convert ranges to actual slices
	result := make(map[U]Collection[T], len(ranges))

	// now do slices to a sorted collection
	for key, val := range ranges {
		result[key] = c[val.start : val.start+val.length]
	}

	return result
}

// splitMutRange is a struct that holds start and length of a range
// it is used in SplitMut function.
type splitMutRange struct {
	start  int
	length int
}
