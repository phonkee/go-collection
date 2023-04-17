package collection

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	data := []struct {
		in     []int
		fn     func(int) uint32
		expect Collection[uint32]
	}{
		{
			[]int{1, 2, 3},
			func(i int) uint32 {
				return uint32(i)
			},
			[]uint32{1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			func(i int) uint32 {
				return uint32(i * i)
			},
			[]uint32{1, 4, 9},
		},
	}

	for _, item := range data {
		got := Convert(item.in, item.fn)
		if !reflect.DeepEqual(got, item.expect) {
			t.Errorf("Convert(%v) != %v", item.in, item.expect)
		}

	}

}
