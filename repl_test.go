package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hi this is a entrance",
			expected: []string{"hi", "this", "is", "a", "entrance"},
		},
		{
			input:    "hi there",
			expected: []string{"hi", "there"},
		},

		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(c.expected) != len(actual) {
			t.Errorf("Output not of correct length, should be %d not %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Word: %s does not match expected %s", word, expectedWord)
			}
		}
	}
}
