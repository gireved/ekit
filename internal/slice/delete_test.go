package slice

import (
	"github.com/stretchr/testify/assert"
	"my-ekit/internal/errs"
	"testing"
)

func TestDelete(t *testing.T) {
	testCase := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{23, 12},
			index:     0,
			wantSlice: []int{12},
			wantVal:   23,
		},
		{
			name:      "index middle",
			slice:     []int{23, 12, 56},
			index:     1,
			wantSlice: []int{23, 56},
			wantVal:   12,
		},
		{
			name:    "index out of range",
			slice:   []int{23, 12, 56},
			index:   4,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
		{
			name:    "index less than 0",
			slice:   []int{23, 12, 45, 65},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(4, -1),
		},
		{
			name:      "index last",
			slice:     []int{23, 12, 45, 65, 100, 90},
			index:     5,
			wantSlice: []int{23, 12, 45, 65, 100},
			wantVal:   90,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
