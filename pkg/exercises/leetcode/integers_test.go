package leetcode

import "testing"

func TestReverse(test *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Reverse of 123", 123, 321},
		{"Reverse of -123", -123, -321},
		{"Reverse of 120", 120, 21},
		{"Reverse of 0", 0, 0},
		{"Reverse of 1534236469 (overflow case)", 1534236469, 0},
	}

	for _, tc := range tests {
		test.Run(tc.name, func(t *testing.T) {
			result := Reverse(tc.input)
			if result != tc.expected {
				t.Errorf("Reverse(%d) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}
