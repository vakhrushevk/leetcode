package twosum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Test case 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Test case 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Test case 3", []int{3, 3}, 6, []int{0, 1}},
		{"Test case 4", []int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{"Test case 5", []int{1, 1, 1, 1, 1}, 6, nil},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			require.Equal(t, tc.want, result)
		})
	}
}
