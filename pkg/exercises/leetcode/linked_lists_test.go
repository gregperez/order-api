package leetcode

import "testing"

func TestLinkedList_AddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want []int
	}{
		{
			name: "Test Case 1",
			l1:   createList([]int{2, 4, 3}),
			l2:   createList([]int{5, 6, 4}),
			want: []int{7, 0, 8},
		},
		{
			name: "Test Case 2",
			l1:   createList([]int{0}),
			l2:   createList([]int{0}),
			want: []int{0},
		},
		{
			name: "Test Case 3",
			l1:   createList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   createList([]int{9, 9, 9, 9}),
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTwoNumbers(tt.l1, tt.l2)
			var result []int
			for got != nil {
				result = append(result, got.Val)
				got = got.Next
			}
			if len(result) != len(tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", result, tt.want)
				return
			}
			for i := range result {
				if result[i] != tt.want[i] {
					t.Errorf("AddTwoNumbers() = %v, want %v", result, tt.want)
					return
				}
			}
		})
	}
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
