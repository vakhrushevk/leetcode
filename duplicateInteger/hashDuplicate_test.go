package duplicateinteger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{[]int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{[]int{1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{[]int{1, 2, 3, 4, 5, 6, 9, 8, 9, 10}, true},
		{[]int{-10, 2, 3, 4, 5, 10, 22}, false},
		{[]int{}, false},
	}

	for _, test := range tests {
		t.Run("test case", func(t *testing.T) {
			result := HashDuplicate(test.input)
			require.Equal(t, test.expected, result)
		})
	}

}
