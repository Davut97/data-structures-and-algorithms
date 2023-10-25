package main

import "testing"

type test []struct {
	input    string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	tests := test{{"hello", false}, {"aba", true}, {"aba ", false}}

	for _, test := range tests {
		isPalindrome := isPalindrome(test.input)
		if isPalindrome != test.expected {
			t.Errorf("ReverseString(%s) = %t; expected %t", test.input, isPalindrome, test.expected)
		}

	}
}
