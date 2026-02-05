package leetcode

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want *string
	}{
		{"Example 1", "swiss", strPtr("w")},
		{"Example 2", "aabbcc", nil},
		{"Example 3", "Hello", strPtr("h")},
		{"Example 4", "abacddbec", strPtr("e")},
		{"Example 5", "aAbBABac", strPtr("c")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstUniqueChar(tt.s)
			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Errorf("FirstUniqueChar() = %v, want %v", got, tt.want)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("FirstUniqueChar() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Example 1", "abcabcbb", 3},
		{"Example 2", "bbbbb", 1},
		{"Example 3", "pwwkew", 3},
		{"Example 4", "", 0},
		{"Example 5", "au", 2},
		{"Example 6", "dvdf", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Single char", "a", 1},
		{"All same", "aaaa", 1},
		{"Two chars", "ab", 2},
		{"Reversed duplicate", "abba", 2},
		{"Long unique", "abcdefghij", 10},
		{"Special chars", "a!@#$%^&*()", 11},
		{"Spaces", "a b c a", 3}, // " b c" o "a b" o "c a" = 3
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"Example 1", "babad", "bab"},
		{"Example 2", "cbbd", "bb"},
		{"Example 3", "a", "a"},
		{"Example 4", "ac", "a"},
		{"Example 5", "racecar", "racecar"},
		{"Example 6", "abb", "bb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
