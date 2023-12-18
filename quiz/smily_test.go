package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmilyFace(t *testing.T) {
	type TestCases struct {
		Input  []string
		Expect int
	}

	testCases := []TestCases{
		{
			Input:  []string{":)", ";(", ";}", ":-D"},
			Expect: 2,
		},
		{
			Input:  []string{";D", ":-(", ":-)", ";~)"},
			Expect: 3,
		},
		{
			Input:  []string{";]", ":[", ";*", ":$", ";-D"},
			Expect: 1,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Expect, CountSmilyFace(tc.Input))
	}
}
