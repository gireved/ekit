package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			want: []int{7},
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			name: "diff 1",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "src less than dst",
		},
		{
			src:  []int{1, 3, 5, 7, 7},
			dst:  []int{1, 3, 5},
			want: []int{7},
			name: "diff deduplicate",
		},
		{
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{7},
			name: "dst duplicate ele",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DiffSet(tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestDiffSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			want: []int{7},
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			name: "diff 1",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "src less than dst",
		},
		{
			src:  []int{1, 3, 5, 7, 7},
			dst:  []int{1, 3, 5},
			want: []int{7},
			name: "diff deduplicate",
		},
		{
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{7},
			name: "dst duplicate ele",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DiffSetFunc[int](tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}
