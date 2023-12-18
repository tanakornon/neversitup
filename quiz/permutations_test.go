package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	type TestCases struct {
		Input  string
		Expect []string
	}

	testCases := []TestCases{
		{
			Input:  "a",
			Expect: []string{"a"},
		},
		{
			Input:  "ab",
			Expect: []string{"ab", "ba"},
		},
		{
			Input:  "abc",
			Expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			Input:  "aabb",
			Expect: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Expect, Permutations(tc.Input))
	}
}
