package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{ // Test 1
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{ // Test 2
			input:    "HELLO World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			continue
		}
		for i := range got {
			if got[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			}
		}
	}
}
