package main

import "testing"

func TestRepl(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Three days of grace",
			expected: []string{"three", "days", "of", "grace"},
		},
		{
			input:    "Lainey is a CUTIE",
			expected: []string{"lainey", "is", "a", "cutie"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %v did not match expected word %v", word, expectedWord)
			}
		}
	}

}
