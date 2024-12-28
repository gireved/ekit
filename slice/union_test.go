package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestUnionSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{1, 2, 3, 4, 5, 6},
			name: "not empty",
		},
		{
			src:  []int{},
			dst:  []int{1, 3},
			want: []int{1, 3},
			name: "src is empty",
		},
		{

			src:  []int{1, 3},
			dst:  []int{},
			want: []int{1, 3},
			name: "dst is empty",
		},
		{
			src:  []int{},
			dst:  []int{},
			want: []int{},
			name: "src and dst are empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := UnionSet[int](tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestUnionSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{1, 2, 3, 4, 5, 6},
			name: "not empty",
		},
		{
			src:  []int{},
			dst:  []int{1, 3},
			want: []int{1, 3},
			name: "src is empty",
		},
		{

			src:  []int{1, 3},
			dst:  []int{},
			want: []int{1, 3},
			name: "dst is empty",
		},
		{
			src:  []int{},
			dst:  []int{},
			want: []int{},
			name: "src and dst are empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := UnionSetFunc[int](tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func ExampleUnionSet() {
	res := UnionSet[int]([]int{1, 3, 4, 5}, []int{1, 4, 7})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 3 4 5 7]
}

func ExampleUnionSetFunc() {
	res := UnionSetFunc[int]([]int{1, 3, 4, 5}, []int{1, 4, 7}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 3 4 5 7]
}
