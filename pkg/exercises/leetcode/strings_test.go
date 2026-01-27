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

func strPtr(s string) *string {
	return &s
}
