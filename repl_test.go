package main

import "testing"

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
			input:    "oNe two tHrEE fOur",
			expected: []string{"one", "two", "three", "four"},
		},
		{
			input:    "heLlo woRlD  ?  ",
			expected: []string{"hello", "world", "?"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length %d and expected length %d does not match", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word %s and expected word %s does not match", word, expectedWord)
			}
		}
	}
}
