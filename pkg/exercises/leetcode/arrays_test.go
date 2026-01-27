package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"No Solution", []int{1, 2, 3}, 7, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TwoSum() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestMinimumAbsDifference(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want [][]int
	}{
		{"Example 1", []int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{"Example 2", []int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{"Example 3", []int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
		{"Single Element", []int{5}, [][]int{}},
		{"No consecutive Pairs", []int{1, 5, 3, 19, 18, 25}, [][]int{{18, 19}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimumAbsDifference(tt.arr)
			if len(got) != len(tt.want) {
				t.Errorf("MinimumAbsDifference() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("MinimumAbsDifference() = %v, want %v", got, tt.want)
						return
					}
				}
			}
		})
	}
}
