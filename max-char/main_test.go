package main

import (
	"testing"
)

func TestMaxChar(t *testing.T) {
	tests := []struct {
		input    string
		expected rune
	}{
		{"aabbccc", 'c'},
		{"abdcaa", 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := MaxChar(tt.input)
			if result != tt.expected {
				t.Errorf("MaxChar(%s) = %c; want %c", tt.input, result, tt.expected)
			}
		})
	}
}
