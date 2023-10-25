package main

import (
	"fmt"
	"testing"
)

func Test_ReverseInteger(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{-11, -11},
		{-123, -321},
		{321, 123},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			result := ReverseInteger(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseInteger(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
