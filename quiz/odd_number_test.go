package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddNumber(t *testing.T) {
	type TestCases struct {
		Input  []int
		Expect int
	}

	testCases := []TestCases{
		{
			Input:  []int{7},
			Expect: 7,
		},
		{
			Input:  []int{0},
			Expect: 0,
		},
		{
			Input:  []int{1, 1, 2},
			Expect: 2,
		},
		{
			Input:  []int{0, 1, 0, 1, 0},
			Expect: 0,
		},
		{
			Input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			Expect: 4,
		},
		{
			Input:  []int{1, 1},
			Expect: -1,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Expect, FindOddNumber(tc.Input))
	}
}
