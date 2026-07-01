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

func TestIsPalindrome(test *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Palindrome 121", 121, true},
		{"Not a palindrome -121", -121, false},
		{"Not a palindrome 10", 10, false},
		{"Palindrome 12321", 12321, true},
	}

	for _, tc := range tests {
		test.Run(tc.name, func(t *testing.T) {
			result := IsPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("IsPalindrome(%d) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
