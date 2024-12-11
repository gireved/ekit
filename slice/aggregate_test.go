package slice

import (
	"github.com/stretchr/testify/assert"
	my_ekit "my-ekit"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{1, 2, 3},
			want:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		Max[int](nil)
	})
	assert.Panics(t, func() {
		Max[int]([]int{})
	})

	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{3},
			want:  3,
		},
		{
			name:  "values",
			input: []Integer{3, 1, 2},
			want:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		Min[int](nil)
	})
	assert.Panics(t, func() {
		Min[int]([]int{})
	})

	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []Integer{},
		},
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{1, 2, 3},
			want:  6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

// testSumTypes 只是用来测试一下满足 Sum 方法约束的所有类型
func testSumTypes[T my_ekit.RealNumber](t *testing.T) {
	res := Sum[T]([]T{1, 2, 3})
	assert.Equal(t, T(6), res)
}

// testMinTypes 只是用来测试一下满足 Min 方法约束的所有类型
func testMinTypes[T my_ekit.RealNumber](t *testing.T) {
	res := Min[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

// testMaxTypes 只是用来测试一下满足 Max 方法约束的所有类型
func testMaxTypes[T my_ekit.RealNumber](t *testing.T) {
	res := Max[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

type Integer int
