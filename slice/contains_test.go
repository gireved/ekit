package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 7, 8, 9},
			dst:  9,
			name: "dst exist",
		},
		{
			want: false,
			src:  []int{1, 4, 7, 8, 9},
			dst:  3,
			name: "dst not exist",
		},
		{
			want: false,
			src:  []int{},
			dst:  3,
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Contains[int](tt.src, tt.dst)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 7, 8, 9},
			dst:  9,
			name: "dst exist",
		},
		{
			want: false,
			src:  []int{1, 4, 7, 8, 9},
			dst:  3,
			name: "dst not exist",
		},
		{
			want: false,
			src:  []int{},
			dst:  3,
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  3,
			name: "src nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ContainsFunc[int](tt.src, func(src int) bool {
				return src == tt.dst
			})
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestContainAny(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			name: "exist two ele",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ContainAny(tt.src, tt.dst)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAll[int](test.src, test.dst))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc[int](test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func ExampleContains() {
	res := Contains[int]([]int{1, 2, 3}, 3)
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsFunc() {
	res := ContainsFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 3
	})
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsAll() {
	res := ContainsAll[int]([]int{1, 2, 3}, []int{3, 1})
	fmt.Println(res)
	res = ContainsAll[int]([]int{1, 2, 3}, []int{3, 1, 4})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAllFunc() {
	res := ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1, 4}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainAny() {
	res := ContainAny[int]([]int{1, 2, 3}, []int{3, 6})
	fmt.Println(res)
	res = ContainAny[int]([]int{1, 2, 3}, []int{4, 5, 9})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainAnyFunc() {
	res := ContainAnyFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{4, 7, 6}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}
