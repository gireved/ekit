package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reverse[int](tt.src)
			assert.ElementsMatch(t, tt.want, res)
			assert.NotSame(t, &tt.src, &res)
		})
	}
}

func TestReverseStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	tests := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			name: "normal test",
			src:  []testStruct{{1, []int{1, 2, 3}}, {3, []int{2, 3, 4}}, {6, []int{4, 6, 9}}},
			want: []testStruct{{6, []int{4, 6, 9}}, {3, []int{2, 3, 4}}, {1, []int{1, 2, 3}}},
		},
		{
			name: "length of src is 0",
			src:  []testStruct{},
			want: []testStruct{},
		},
		{
			src:  nil,
			want: []testStruct{},
			name: "length of src is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reverse[testStruct](tt.src)
			assert.ElementsMatch(t, tt.want, res)
			assert.NotSame(t, &tt.src, &res)
		})
	}
}

func TestReverseSelfInt(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf[int](tt.src)
			assert.ElementsMatch(t, tt.want, tt.src)
		})
	}
}

func TestReverseSelfStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	tests := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			name: "normal test",
			src:  []testStruct{{1, []int{1, 2, 3}}, {3, []int{2, 3, 4}}, {6, []int{4, 6, 9}}},
			want: []testStruct{{6, []int{4, 6, 9}}, {3, []int{2, 3, 4}}, {1, []int{1, 2, 3}}},
		},
		{
			name: "length of src is 0",
			src:  []testStruct{},
			want: []testStruct{},
		},
		{
			src:  nil,
			want: []testStruct{},
			name: "length of src is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf[testStruct](tt.src)
			assert.ElementsMatch(t, tt.want, tt.src)
		})
	}
}

func ExampleReverse() {
	res := Reverse[int]([]int{1, 3, 2, 2, 4})
	fmt.Println(res)
	res2 := Reverse[string]([]string{"a", "b", "c", "d", "e"})
	fmt.Println(res2)
	// Output:
	// [4 2 2 3 1]
	// [e d c b a]
}

func ExampleReverseSelf() {
	src := []int{1, 3, 2, 2, 4}
	ReverseSelf[int](src)
	fmt.Println(src)
	src2 := []string{"a", "b", "c", "d", "e"}
	ReverseSelf[string](src2)
	fmt.Println(src2)
	// Output:
	// [4 2 2 3 1]
	// [e d c b a]
}
