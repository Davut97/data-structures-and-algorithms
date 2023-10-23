package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World! ", " !dlroW ,olleH"},
		{"12345", "54321"},
		{"", ""}, // Edge case: empty string
	}

	// Iterate through the test cases
	for _, test := range tests {
		reversed := reverseString(test.input)
		if reversed != test.expected {
			t.Errorf("ReverseString(%s) = %s; expected %s", test.input, reversed, test.expected)
		}
	}
}
