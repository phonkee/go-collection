package collection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollection(t *testing.T) {
	t.Run("Test Cycle", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			much   int
			expect Collection[int]
		}{
			{[]int{1, 2, 3}, 4, []int{1, 2, 3, 1}},
			{[]int{}, 10, []int{}},
		}
		for _, item := range data {
			found := item.in.Cycle(item.much)
			if !reflect.DeepEqual(found, item.expect) {
				t.Errorf("Expected %v, found %v", item.expect, found)
			}
		}
	})

	t.Run("Test Each", func(t *testing.T) {
		expect := []int{1, 2, 3}
		found := make([]int, 0)
		Collection[int]([]int{1, 2, 3}).Each(func(i int) {
			found = append(found, i)
		})
		if !reflect.DeepEqual(expect, found) {
			t.Errorf("Expected %v, found %v", expect, found)
		}
	})

	t.Run("Test Filter", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			filter FilterFunc[int]
			expect Collection[int]
		}{
			{[]int{1, 2, 3}, func(i int) bool { return i&1 == 0 }, []int{2}},
		}

		for _, item := range data {
			found := item.in.Filter(item.filter)
			if !reflect.DeepEqual(found, Collection[int](item.expect)) {
				t.Errorf("Expected %v, found %v", item.expect, found)
			}
		}

	})

	t.Run("Test First", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			expect int
			result bool
		}{
			{[]int{1, 2, 3}, 1, true},
			{[]int{}, 0, false},
		}

		for _, item := range data {
			got := item.in.First(func(got int) {
				if !reflect.DeepEqual(item.expect, got) {
					t.Errorf("Expected %v, found %v", item.expect, got)
				}
			})
			if got != item.result {
				t.Error("First failed")
			}

		}
	})

	t.Run("Test Last", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			expect int
			result bool
		}{
			{[]int{1, 2, 3}, 3, true},
			{[]int{}, 0, false},
		}

		for _, item := range data {
			got := item.in.Last(func(got int) {
				if !reflect.DeepEqual(item.expect, got) {
					t.Errorf("Expected %v, found %v", item.expect, got)
				}
			})
			if got != item.result {
				t.Error("Last failed")
			}

		}
	})

	t.Run("Test Take", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			take   int
			expect Collection[int]
		}{
			{[]int{1, 2, 3}, 2, []int{1, 2}},
		}

		for _, item := range data {
			taken := item.in.Take(item.take)
			if !reflect.DeepEqual(taken, item.expect) {
				t.Errorf("Expected %v, found %v", item.expect, taken)
			}
		}

	})

	t.Run("Test Unique", func(t *testing.T) {
		data := []struct {
			in     Collection[int]
			expect Collection[int]
		}{
			{[]int{1, 2, 3, 3, 3, 2}, []int{1, 2, 3}},
		}

		for _, item := range data {
			result := item.in.Unique(func(item int) string {
				return fmt.Sprint(item)
			})
			if !reflect.DeepEqual(result, item.expect) {
				t.Errorf("Expected %v, found %v", item.expect, result)
			}
		}
	})

	t.Run("Test Sort", func(t *testing.T) {
		data := []struct {
			in       Collection[int]
			sortFunc func(i, j int) bool
			expect   Collection[int]
		}{
			{[]int{3, 2, 1, 4}, func(i, j int) bool { return i > j }, []int{4, 3, 2, 1}},
			{[]int{3, 2, 1, 4}, func(i, j int) bool { return i < j }, []int{1, 2, 3, 4}},
			{[]int{3, 2, 1, 4}, SortBackward(func(i, j int) bool { return i > j }), []int{1, 2, 3, 4}},
		}

		for _, item := range data {
			result := item.in.Sort(item.sortFunc)
			if !reflect.DeepEqual(result, item.expect) {
				t.Errorf("Expected %v, found %v", item.expect, result)
			}
		}
	})

	t.Run("Test format", func(t *testing.T) {
		// raw debug format function
		ff := func(in int) string {
			return fmt.Sprint("this value:", in)
		}

		// use raw debug format function
		Collection[int]([]int{1, 2, 3}).Cycle(6).Sub(func(collection Collection[int]) {
			fmt.Printf("Collection is here: %p\n", collection)
		}).Debug("all", ff)

		// predefine format function using Debug
		ff2 := Debug("hello", func(t int) string {
			return fmt.Sprint(t ^ 3)
		})

		// showing multiple values
		Collection[int]([]int{1, 2, 3}).Debug("all", Debug("this is it", func(t int) string {
			return fmt.Sprint(t)
		}), ff2, ff2, ff2, ff2)
	})

	t.Run("Test Complete", func(t *testing.T) {
		var got int
		var expected = 2
		coll := Collection[int]{1, 2, 3}
		coll.
			Sub(func(c Collection[int]) {
				// do something with whole collection independently
			}).
			Each(func(i int) {
				// iterate over each item in collection
				fmt.Printf("square of %d is %d\n", i, i*i)
			}).
			// now debug (print default) with title
			Debug("after").
			// Filter all odd items
			Filter(func(t int) bool {
				return t&1 == 0
			}).
			// custom debug
			Debug("custom", Debug("modulo_2", func(t int) string {
				return fmt.Sprintf("value %v = %v", t, t%2)
			})).
			// do something with last item (if available) - returns bool if called
			Last(func(item int) {
				got = item
			})
		if got != expected {
			t.Error("Last failed")
		}
	})

	t.Run("Test Map", func(t *testing.T) {
		if !reflect.DeepEqual(Collection[int]{1, 2, 3}.Map(func(i int) int {
			return 42
		}), Collection[int]{42, 42, 42}) {
			t.Error("Map failed")
		}
	})

	t.Run("Test Copy", func(t *testing.T) {
		orig := Collection[int]{1, 2, 3}
		if !reflect.DeepEqual(orig, orig.Copy()) {
			t.Error("Copy failed")
		}
		if reflect.DeepEqual(fmt.Sprintf("%p", orig), fmt.Sprintf("%p", orig.Copy())) {
			t.Error("Copy failed")
		}
	})

	t.Run("Test TakeUntil", func(t *testing.T) {
		orig := Collection[int]{1, 2, 3}
		expect := Collection[int]{1, 2}

		result := orig.TakeUntil(func(t int) bool {
			return t <= 2
		})

		if !reflect.DeepEqual(result, expect) {
			t.Error("TakeUntil failed")
		}
	})

	t.Run("Test Chain", func(t *testing.T) {
		first := Collection[int]{1, 2, 3}
		second := Collection[int]{4, 5, 6}

		if !reflect.DeepEqual(first.Chain(second), Collection[int]{1, 2, 3, 4, 5, 6}) {
			t.Error("Chain failed")
		}
	})

	t.Run("Test Shuffle", func(t *testing.T) {
		first := Collection[int]{1, 2, 3}
		shuffled := first.Shuffle()

		if reflect.DeepEqual(first, shuffled) {
			t.Error("Shuffle failed")
		}

		if !reflect.DeepEqual(first, shuffled.Sort(func(t1, t2 int) bool { return t1 < t2 })) {
			t.Error("Shuffle failed")
		}

	})
}
