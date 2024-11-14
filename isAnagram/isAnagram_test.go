package isanagram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		input        string
		inputReverse string
		expected     bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"listen", "silent", true},
		{"rat", "tar", true},
		{"evil", "vile", true},
		{"god", "dog", true},
		{"night", "thing", true},
		{"binary", "brainy", true},
		{"abcd", "dcba", true},
		{"civic", "vicic", true},
		{"angel", "glean", true},
		{"ditwy", "tidy", false},
		{"aabbcc", "abcabc", true},
		{"racecar", "carrace", true},
		{"jar", "raj", true},
		{"a", "b", false},
		{"abc", "ab", false},
		{"abc", "def", false},
		{"abc", "abcd", false},
		{"apple", "pale", false},
		{"abc", "acb", true},
		{"hello", "ehllo", true},
	}
	for _, test := range tests {
		t.Run(test.input+":"+test.inputReverse, func(t *testing.T) {
			result := IsAnagram(test.input, test.inputReverse)
			require.Equal(t, test.expected, result)
		})

	}

}
