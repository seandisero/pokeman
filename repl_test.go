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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		// more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected a slice length of %d, got %d", len(c.expected), len(actual))
			t.Fail()
		}
		// if they don't match, use the t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s got %s", expectedWord, word)
				t.Fail()
			}
		}
	}
}
