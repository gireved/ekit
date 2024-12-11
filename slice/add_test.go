package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"my-ekit/internal/errs"
	"testing"
)

func TestAdd(t *testing.T) {
	// Add 主要依赖于 internal/slice.Add 来保证正确性
	testCase := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// my-ekit: 下标超出范围，长度 4,下标 -1
}
